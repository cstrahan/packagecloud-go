package pcloud

import (
	"context"
	"errors"
	"net/http"
	"slices"
	"strconv"
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

// fakeBackend models a paginated server holding items [0, total) with a
// per_page ceiling of serverMax. It records every (page, perPage) it's asked
// for so tests can assert which pages were actually fetched.
type fakeBackend struct {
	total, serverMax int
	mu               sync.Mutex
	calls            [][2]int // (page, perPage)
}

func (b *fakeBackend) fetch(_ context.Context, page, perPage int) ([]int, http.Header, error) {
	b.mu.Lock()
	b.calls = append(b.calls, [2]int{page, perPage})
	b.mu.Unlock()
	eff := perPage
	if b.serverMax > 0 && eff > b.serverMax {
		eff = b.serverMax
	}
	h := http.Header{}
	h.Set("Total", strconv.Itoa(b.total))
	h.Set("Per-Page", strconv.Itoa(eff))
	var items []int
	for i := (page - 1) * eff; i < page*eff && i < b.total; i++ {
		items = append(items, i)
	}
	return items, h, nil
}

func (b *fakeBackend) pagesFetched() map[int]bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	m := map[int]bool{}
	for _, c := range b.calls {
		m[c[0]] = true
	}
	return m
}

// seq returns the slice [start, start+n).
func seq(start, n int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = start + i
	}
	return out
}

func TestFetchWindow(t *testing.T) {
	cases := []struct {
		name      string
		total     int
		opts      ListOptions
		want      []int
		wantPages []int // pages that MUST have been fetched (exact set)
		serverMax int
	}{
		{
			name:      "full listing in order",
			total:     25,
			opts:      ListOptions{PerPage: 10},
			want:      seq(0, 25),
			wantPages: []int{1, 2, 3},
			serverMax: 250,
		},
		{
			name:      "limit within first page fetches one page",
			total:     100,
			opts:      ListOptions{PerPage: 10, Limit: 3},
			want:      seq(0, 3),
			wantPages: []int{1},
			serverMax: 250,
		},
		{
			name:      "offset+limit spanning two pages",
			total:     50,
			opts:      ListOptions{PerPage: 10, Offset: 8, Limit: 5},
			want:      seq(8, 5), // 8..12, on pages 1 and 2
			wantPages: []int{1, 2},
			serverMax: 250,
		},
		{
			name:      "offset deep in listing skips earlier pages",
			total:     50,
			opts:      ListOptions{PerPage: 10, Offset: 22, Limit: 3},
			want:      seq(22, 3), // 22..24, page 3 only
			wantPages: []int{3},
			serverMax: 250,
		},
		{
			name:      "offset past end yields empty",
			total:     10,
			opts:      ListOptions{PerPage: 10, Offset: 99, Limit: 5},
			want:      nil,
			wantPages: []int{10}, // startPage computed from offset; returns empty
			serverMax: 250,
		},
		{
			name:      "per_page shrinks to the window",
			total:     100,
			opts:      ListOptions{PerPage: 250, Limit: 4},
			want:      seq(0, 4),
			wantPages: []int{1},
			serverMax: 250,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b := &fakeBackend{total: c.total, serverMax: c.serverMax}
			got, err := fetchWindow(context.Background(), c.opts, b.fetch)
			if err != nil {
				t.Fatal(err)
			}
			if !slices.Equal(got, c.want) {
				t.Errorf("window = %v, want %v", got, c.want)
			}
			pages := b.pagesFetched()
			if len(pages) != len(c.wantPages) {
				t.Errorf("fetched pages %v, want exactly %v", keysOf(pages), c.wantPages)
			}
			for _, p := range c.wantPages {
				if !pages[p] {
					t.Errorf("expected page %d to be fetched; fetched %v", p, keysOf(pages))
				}
			}
		})
	}
}

// When the server clamps per_page below what we requested, fetchWindow must
// recompute pages against the effective size and still return the right items.
func TestFetchWindowServerClamp(t *testing.T) {
	b := &fakeBackend{total: 25, serverMax: 10} // we ask for 250, server caps at 10
	got, err := fetchWindow(context.Background(), ListOptions{PerPage: 250}, b.fetch)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(got, seq(0, 25)) {
		t.Fatalf("window = %v, want 0..24", got)
	}
}

// Header-less (unpaginated) responses return the single page, windowed.
func TestFetchWindowNoHeaders(t *testing.T) {
	calls := 0
	fetch := func(_ context.Context, page, perPage int) ([]int, http.Header, error) {
		calls++
		return []int{0, 1, 2, 3, 4}, http.Header{}, nil
	}
	got, err := fetchWindow(context.Background(), ListOptions{Offset: 1, Limit: 2}, fetch)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(got, []int{1, 2}) || calls != 1 {
		t.Fatalf("got %v in %d calls, want [1 2] in 1 call", got, calls)
	}
}

func keysOf(m map[int]bool) []int {
	var out []int
	for k := range m {
		out = append(out, k)
	}
	slices.Sort(out)
	return out
}
