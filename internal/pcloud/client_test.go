package pcloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"testing"

	pc "github.com/cstrahan/packagecloud-go/sdk"
	"github.com/cstrahan/packagecloud-go/sdk/core"
)

func TestIsDuplicateError(t *testing.T) {
	dup := core.NewAPIError(http.StatusUnprocessableEntity, nil,
		errors.New(`{"filename":["has already been taken"]}`))
	if !IsDuplicateError(dup) {
		t.Error("expected duplicate 422 to be recognized")
	}
	// The wording the live API actually returns (caught in end-to-end testing).
	dupLive := core.NewAPIError(http.StatusUnprocessableEntity, nil,
		errors.New(`{"repository":["Validation failed: Full name already exists."]}`))
	if !IsDuplicateError(dupLive) {
		t.Error("expected live 'Full name already exists' 422 to be recognized")
	}
	// A 422 for a different reason is not a duplicate.
	other := core.NewAPIError(http.StatusUnprocessableEntity, nil,
		errors.New(`{"base":["something else"]}`))
	if IsDuplicateError(other) {
		t.Error("non-duplicate 422 should not be treated as a duplicate")
	}
	// A different status code is not a duplicate.
	notFound := core.NewAPIError(http.StatusNotFound, nil, errors.New("nope"))
	if IsDuplicateError(notFound) {
		t.Error("404 should not be treated as a duplicate")
	}
	if IsDuplicateError(errors.New("plain error")) {
		t.Error("plain error should not be treated as a duplicate")
	}
}

func TestSplitRepo(t *testing.T) {
	user, name, err := SplitRepo("acme/widgets")
	if err != nil || user != "acme" || name != "widgets" {
		t.Fatalf("got (%q, %q, %v), want (acme, widgets, nil)", user, name, err)
	}
	for _, bad := range []string{"", "noslash", "/widgets", "acme/"} {
		if _, _, err := SplitRepo(bad); err == nil {
			t.Errorf("SplitRepo(%q) = nil error, want error", bad)
		}
	}
}

func TestExtToPackageType(t *testing.T) {
	cases := map[string]string{
		"deb": "deb", "DEB": "deb", ".deb": "deb",
		"rpm": "rpm", "gem": "gem",
		"jar": "java", "aar": "java", "war": "java",
		"whl": "py", "egg": "py",
		"asc": "anyfile", "apk": "alpine",
		// genuinely ambiguous / unknown -> ""
		"tgz": "", "zip": "", "": "", "xyz": "",
	}
	for ext, want := range cases {
		if got := extToPackageType(ext); got != want {
			t.Errorf("extToPackageType(%q) = %q, want %q", ext, got, want)
		}
	}
}

func TestBuildDeleteURL(t *testing.T) {
	const base = "https://p.io"
	cases := []struct {
		name                    string
		distroVersion, filename string
		want                    string
	}{
		{"deb", "ubuntu/xenial", "foo.deb",
			"https://p.io/api/v1/repos/u/r/ubuntu/xenial/foo.deb"},
		{"gem explicit", "gems", "foo.gem",
			"https://p.io/api/v1/repos/u/r/gems/foo.gem"},
		{"gem auto-inferred", "", "foo.gem",
			"https://p.io/api/v1/repos/u/r/gems/foo.gem"},
		{"java maven group", "java/maven2/com.example", "widget-1.0.jar",
			"https://p.io/api/v1/repos/u/r/java/maven2/com.example/widget-1.0.jar"},
		{"strips path and slashes", "/ubuntu/xenial/", "/tmp/foo.deb",
			"https://p.io/api/v1/repos/u/r/ubuntu/xenial/foo.deb"},
		{"multi-extension filename kept whole", "python/1", "pkg-1.0.tar.gz",
			"https://p.io/api/v1/repos/u/r/python/1/pkg-1.0.tar.gz"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := buildDeleteURL(base, "u", "r", c.distroVersion, c.filename); got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestBuildLookup(t *testing.T) {
	str := func(s string) *string { return &s }
	dists := map[string][]*pc.Distribution{
		"deb": {{
			IndexName: str("ubuntu"),
			Versions: []*pc.DistroVersion{
				{ID: 166, IndexName: str("xenial"), VersionNumber: str("16.04")},
			},
		}},
		// single distro + single version => bare distro & package_type shortcuts
		"py": {{
			IndexName: str("python"),
			Versions:  []*pc.DistroVersion{{ID: 999, IndexName: str("1")}},
		}},
	}
	lookup := buildLookup(dists)
	want := map[string]int{
		"ubuntu/xenial": 166,
		"ubuntu/16.04":  166,
		"python/1":      999,
		"python":        999, // bare distro shortcut
		"py":            999, // package_type shortcut
	}
	for k, v := range want {
		if lookup[k] != v {
			t.Errorf("lookup[%q] = %d, want %d", k, lookup[k], v)
		}
	}
}

func TestHeaderInt(t *testing.T) {
	h := http.Header{}
	h.Set("Total", "  42 ")
	if got := headerInt(h, "Total"); got != 42 {
		t.Errorf("Total = %d, want 42", got)
	}
	if got := headerInt(h, "Missing"); got != 0 {
		t.Errorf("Missing = %d, want 0", got)
	}
	if got := headerInt(nil, "Total"); got != 0 {
		t.Errorf("nil header = %d, want 0", got)
	}
}

// fetchAllPages should fetch every page, preserve page order in the flattened
// result, and derive the page count from the Total / Per-Page headers.
func TestFetchAllPagesOrderAndCompleteness(t *testing.T) {
	const total, perPage = 25, 10 // => 3 pages: 10, 10, 5
	var mu sync.Mutex
	seen := map[int]bool{}

	fetch := func(_ context.Context, page int) ([]int, http.Header, error) {
		mu.Lock()
		seen[page] = true
		mu.Unlock()
		h := http.Header{}
		h.Set("Total", fmt.Sprint(total))
		h.Set("Per-Page", fmt.Sprint(perPage))
		start := (page - 1) * perPage
		var items []int
		for i := start; i < start+perPage && i < total; i++ {
			items = append(items, i)
		}
		return items, h, nil
	}

	got, err := fetchAllPages(context.Background(), fetch)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != total {
		t.Fatalf("len(got) = %d, want %d", len(got), total)
	}
	for i, v := range got {
		if v != i { // 0..24 in order
			t.Fatalf("got[%d] = %d, want %d (ordering broken)", i, v, i)
		}
	}
	for p := 1; p <= 3; p++ {
		if !seen[p] {
			t.Errorf("page %d was never fetched", p)
		}
	}
}

// Without pagination headers, fetchAllPages returns just the first page.
func TestFetchAllPagesNoHeaders(t *testing.T) {
	calls := 0
	fetch := func(_ context.Context, page int) ([]int, http.Header, error) {
		calls++
		return []int{1, 2, 3}, http.Header{}, nil
	}
	got, err := fetchAllPages(context.Background(), fetch)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 || calls != 1 {
		t.Fatalf("got %d items in %d calls, want 3 items in 1 call", len(got), calls)
	}
}
