package pcloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	pc "github.com/cstrahan/packagecloud-go/sdk"
	"github.com/cstrahan/packagecloud-go/sdk/client"
	"github.com/cstrahan/packagecloud-go/sdk/core"
	"github.com/cstrahan/packagecloud-go/sdk/option"
)

// maxParallelPages caps the number of in-flight page requests during
// concurrent pagination, matching the reference client.
const maxParallelPages = 10

// defaultPerPage is the wire page size used when the caller doesn't specify
// one. It matches PackageCloud's Max-Per-Page so a full listing needs the
// fewest requests.
const defaultPerPage = 250

// ListOptions controls windowing for the paginated list endpoints. The zero
// value (Offset 0, Limit 0) fetches the entire listing.
type ListOptions struct {
	// PerPage is the requested wire page size; <=0 uses defaultPerPage.
	PerPage int
	// Offset is the number of items to skip from the start of the listing.
	Offset int
	// Limit caps the number of items returned; <=0 means no cap (all).
	Limit int
}

// App wraps the generated SDK client with the higher-level conveniences the
// CLI needs: a distributions cache + distro_version lookup, concurrent
// pagination, and a raw HTTP path for the polymorphic delete endpoint.
type App struct {
	sdk     *client.Client
	http    *http.Client
	token   string
	baseURL string // e.g. "https://packagecloud.io" (no trailing slash)

	distOnce  sync.Once
	distCache map[string][]*pc.Distribution
	distErr   error
	lookup    map[string]int // distro_version spelling -> distro_version_id
}

// New builds an App. token and configPath override the corresponding config
// sources; pass "" to fall back to the environment / default config file.
func New(token, configPath string) (*App, error) {
	cfg, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}
	tok := cfg.EffectiveToken(token)
	if tok == "" {
		return nil, fmt.Errorf(
			"API token required: set PACKAGECLOUD_TOKEN, create ~/.packagecloud, " +
				"or pass --token. Get your token at https://packagecloud.io/api_token")
	}
	baseURL := cfg.EffectiveURL()

	sdkClient := client.NewClient(
		option.WithBaseURL(baseURL),
		// PackageCloud uses HTTP basic auth with the token as the username
		// and an empty password.
		option.WithBasicAuth(tok, ""),
	)

	return &App{
		sdk:     sdkClient,
		http:    &http.Client{},
		token:   tok,
		baseURL: baseURL,
	}, nil
}

// SplitRepo parses a "user/repo" identifier into its two components.
func SplitRepo(repository string) (user, name string, err error) {
	parts := strings.SplitN(repository, "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("repository must be in the form \"user/repo\", got %q", repository)
	}
	return parts[0], parts[1], nil
}

// Distributions returns all supported distributions, caching the result and
// building the distro_version lookup on first call.
func (a *App) Distributions(ctx context.Context) (map[string][]*pc.Distribution, error) {
	a.distOnce.Do(func() {
		dists, err := a.sdk.Distributions.Index(ctx)
		if err != nil {
			a.distErr = err
			return
		}
		a.distCache = dists
		a.lookup = buildLookup(dists)
	})
	return a.distCache, a.distErr
}

// buildLookup keys every user-facing distro_version spelling to its numeric
// id, mirroring the reference client:
//   - "ubuntu/xenial"  (distro index / version index)
//   - "ubuntu/16.04"   (distro index / version number)
//   - "python", "py"   (single-version-type shortcuts: distro name + type name)
func buildLookup(dists map[string][]*pc.Distribution) map[string]int {
	lookup := make(map[string]int)
	for pkgType, distList := range dists {
		for _, dist := range distList {
			index := ""
			if dist.IndexName != nil {
				index = *dist.IndexName
			}
			for _, ver := range dist.Versions {
				if ver.VersionNumber != nil && index != "" {
					lookup[index+"/"+*ver.VersionNumber] = ver.ID
				}
				if ver.IndexName != nil && index != "" {
					lookup[index+"/"+*ver.IndexName] = ver.ID
				}
			}
			// Single-version package types (python, java, node, anyfile, …):
			// accept the bare distro name and the package_type name.
			if len(distList) == 1 && len(dist.Versions) == 1 {
				id := dist.Versions[0].ID
				if index != "" {
					lookup[index] = id
				}
				lookup[pkgType] = id
			}
		}
	}
	return lookup
}

// extToPackageType maps a file extension to a package_type, used to infer a
// default distro when the caller didn't pass one. Returns "" for extensions
// that legitimately span several types (the caller must then disambiguate).
func extToPackageType(ext string) string {
	switch strings.ToLower(strings.TrimPrefix(ext, ".")) {
	case "deb":
		return "deb"
	case "dsc":
		return "dsc"
	case "rpm":
		return "rpm"
	case "gem":
		return "gem"
	case "jar", "aar", "war":
		return "java"
	case "whl", "egg", "egg-info", "tar", "gz", "bz2", "z":
		return "py"
	case "asc":
		return "anyfile"
	case "apk":
		return "alpine"
	default:
		return ""
	}
}

// resolveDistroVersionID resolves the numeric distro_version_id the upload
// form requires. distroVersion may be a numeric id, a slug ("ubuntu/xenial",
// "python/1", "python", "py"), or "" to infer from the file extension. A nil
// result means "omit the id" — correct for rubygems, which doesn't take one.
func (a *App) resolveDistroVersionID(ctx context.Context, ext, distroVersion string) (*int, error) {
	if distroVersion != "" {
		if id, err := strconv.Atoi(distroVersion); err == nil {
			return &id, nil
		}
	} else if extToPackageType(ext) == "gem" {
		return nil, nil // rubygems: no distro_version_id
	}

	if _, err := a.Distributions(ctx); err != nil {
		return nil, err
	}

	key := distroVersion
	if key == "" {
		key = extToPackageType(ext)
		if key == "" {
			return nil, fmt.Errorf(
				"cannot infer package type from %q extension; pass --distro-version "+
					"(e.g. ubuntu/xenial, python/1, java/maven2)", ext)
		}
	}
	id, ok := a.lookup[key]
	if !ok {
		return nil, fmt.Errorf(
			"unknown distro_version %q — run `packagecloud distributions` to see valid values", key)
	}
	return &id, nil
}

// ListPackages returns the requested window of packages in a repository,
// fetching only the overlapping pages (concurrently).
func (a *App) ListPackages(ctx context.Context, repository string, opts ListOptions) ([]*pc.PackageFragment, error) {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return nil, err
	}
	return fetchWindow(ctx, opts, func(ctx context.Context, page, perPage int) ([]*pc.PackageFragment, http.Header, error) {
		req := &pc.PackagesAllRequest{UserID: user, Repo: repo, Page: pc.Int(page), PerPage: pc.Int(perPage)}
		resp, err := a.sdk.Packages.WithRawResponse.All(ctx, req)
		if err != nil {
			return nil, nil, err
		}
		return resp.Body, resp.Header, nil
	})
}

// SearchParams holds the optional filters for SearchPackages. Empty fields
// are omitted from the request.
type SearchParams struct {
	Query  string
	Dist   string
	Filter string
	Arch   string
}

// SearchPackages searches a repository, returning the requested window and
// fetching only the overlapping pages (concurrently).
func (a *App) SearchPackages(ctx context.Context, repository string, params SearchParams, opts ListOptions) ([]*pc.PackageFragment, error) {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return nil, err
	}
	return fetchWindow(ctx, opts, func(ctx context.Context, page, perPage int) ([]*pc.PackageFragment, http.Header, error) {
		req := &pc.PackagesSearchRequest{
			UserID:  user,
			Repo:    repo,
			Page:    pc.Int(page),
			PerPage: pc.String(strconv.Itoa(perPage)),
		}
		if params.Query != "" {
			req.Q = pc.String(params.Query)
		}
		if params.Dist != "" {
			req.Dist = pc.String(params.Dist)
		}
		if params.Filter != "" {
			req.Filter = pc.String(params.Filter)
		}
		if params.Arch != "" {
			req.Arch = pc.String(params.Arch)
		}
		resp, err := a.sdk.Packages.WithRawResponse.Search(ctx, req)
		if err != nil {
			return nil, nil, err
		}
		return resp.Body, resp.Header, nil
	})
}

// PushOptions carries the optional knobs for PushPackage.
type PushOptions struct {
	// DistroVersion is the routing target; see resolveDistroVersionID for the
	// accepted forms. Empty means "infer from the file extension".
	DistroVersion string
	// Coordinates overrides the Maven coordinates for a JAR/WAR/AAR when they
	// can't be derived from the artifact's metadata.
	Coordinates string
}

// PushPackage uploads a package file to a repository.
func (a *App) PushPackage(ctx context.Context, repository, packageFile string, opts PushOptions) (*pc.PackageDetails, error) {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(packageFile)
	if err != nil {
		return nil, fmt.Errorf("opening package file: %w", err)
	}
	defer f.Close()

	id, err := a.resolveDistroVersionID(ctx, filepath.Ext(packageFile), opts.DistroVersion)
	if err != nil {
		return nil, err
	}
	req := &pc.PackagesCreateRequest{
		UserID:                 user,
		Repo:                   repo,
		PackagePackageFile:     f,
		PackageDistroVersionID: id,
	}
	if opts.Coordinates != "" {
		req.PackageCoordinates = pc.String(opts.Coordinates)
	}
	return a.sdk.Packages.Create(ctx, req)
}

// IsDuplicateError reports whether err is the 422 the API returns when a
// package already exists — used to implement `push --skip-duplicates`.
//
// The message wording varies by package type / endpoint: the live API has
// returned both "Validation failed: Full name already exists." (under a
// "repository" key) and the Ruby client's "has already been taken" (under
// "filename"), so we match either, case-insensitively.
func IsDuplicateError(err error) bool {
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusUnprocessableEntity {
		return false
	}
	msg := strings.ToLower(apiErr.Error())
	return strings.Contains(msg, "already been taken") || strings.Contains(msg, "already exists")
}

// PackageContents inspects a .dsc source package, returning the files it
// references. distroVersion is required (a .dsc has no default distro).
func (a *App) PackageContents(ctx context.Context, repository, packageFile, distroVersion string) ([]*pc.PackageContents, error) {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(packageFile)
	if err != nil {
		return nil, fmt.Errorf("opening package file: %w", err)
	}
	defer f.Close()

	id, err := a.resolveDistroVersionID(ctx, filepath.Ext(packageFile), distroVersion)
	if err != nil {
		return nil, err
	}
	if id == nil {
		return nil, fmt.Errorf("packages/contents requires --distro-version (e.g. ubuntu/xenial)")
	}
	return a.sdk.Packages.Contents(ctx, &pc.PackagesContentsRequest{
		UserID:                 user,
		Repo:                   repo,
		PackagePackageFile:     f,
		PackageDistroVersionID: id,
	})
}

// PromotePackage moves a package from one repository to another.
// distroVersion is "distro/version" (e.g. "ubuntu/xenial").
func (a *App) PromotePackage(ctx context.Context, source, destination, distroVersion, filename, scope string) (*pc.PackageDetails, error) {
	user, repo, err := SplitRepo(source)
	if err != nil {
		return nil, err
	}
	distro, version, ok := strings.Cut(distroVersion, "/")
	if !ok {
		return nil, fmt.Errorf("distro_version must be in the form \"distro/version\", got %q", distroVersion)
	}
	req := &pc.PackagesPromoteRequest{
		UserID:      user,
		Repo:        repo,
		Distro:      distro,
		Version:     version,
		Package:     filepath.Base(filename),
		Destination: pc.String(destination),
	}
	if scope != "" {
		req.Scope = pc.String(scope)
	}
	return a.sdk.Packages.Promote(ctx, req)
}

// fetchWindow returns the [Offset, Offset+Limit) item window of a paginated
// endpoint, fetching only the pages that overlap the window (concurrently, up
// to maxParallelPages). It exploits PackageCloud's page/per_page query params
// plus the Total/Per-Page response headers, so a small --limit doesn't pull
// the entire listing. With Offset 0 and Limit <=0 it returns every page.
//
// The fetch callback receives the 1-based page index and the per_page to
// request, and returns that page's items plus the response headers.
func fetchWindow[T any](
	ctx context.Context,
	opts ListOptions,
	fetch func(ctx context.Context, page, perPage int) ([]T, http.Header, error),
) ([]T, error) {
	perPage := opts.PerPage
	if perPage <= 0 {
		perPage = defaultPerPage
	}
	offset := max(opts.Offset, 0)
	// When the whole window fits in less than a page, shrink per_page so we
	// transfer only what we need (e.g. --limit 5 fetches 5 items, not 250).
	if opts.Limit > 0 && offset+opts.Limit < perPage {
		perPage = offset + opts.Limit
	}

	// Fetch the first page that can contain the window's start (1-based).
	startPage := offset/perPage + 1
	first, hdr, err := fetch(ctx, startPage, perPage)
	if err != nil {
		return nil, err
	}

	total := headerInt(hdr, "Total")
	effPerPage := headerInt(hdr, "Per-Page")

	// Header-less / unpaginated response: we already hold the full result.
	if total <= 0 || effPerPage <= 0 {
		return windowSlice(first, offset, opts.Limit), nil
	}
	// If the server clamped per_page below what we assumed, our page math was
	// against the wrong size — recompute against the effective size and refetch.
	if effPerPage != perPage {
		perPage = effPerPage
		startPage = offset/perPage + 1
		first, hdr, err = fetch(ctx, startPage, perPage)
		if err != nil {
			return nil, err
		}
		total = headerInt(hdr, "Total")
	}
	if offset >= total {
		return []T{}, nil
	}

	end := total // exclusive upper item bound we need
	if opts.Limit > 0 && offset+opts.Limit < end {
		end = offset + opts.Limit
	}
	totalPages := (total + perPage - 1) / perPage
	endPage := min((end+perPage-1)/perPage, totalPages)

	// Collect pages [startPage, endPage]; startPage is already fetched.
	pages := make([][]T, endPage-startPage+1)
	pages[0] = first

	var (
		wg       sync.WaitGroup
		mu       sync.Mutex
		firstErr error
	)
	sem := make(chan struct{}, maxParallelPages)
	for p := startPage + 1; p <= endPage; p++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(page int) {
			defer wg.Done()
			defer func() { <-sem }()
			items, _, err := fetch(ctx, page, perPage)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				if firstErr == nil {
					firstErr = err
				}
				return
			}
			pages[page-startPage] = items
		}(p)
	}
	wg.Wait()
	if firstErr != nil {
		return nil, firstErr
	}

	var collected []T
	for _, pg := range pages {
		collected = append(collected, pg...)
	}
	// The first collected item sits at global index (startPage-1)*perPage.
	return windowSlice(collected, offset-(startPage-1)*perPage, opts.Limit), nil
}

// windowSlice drops offset items from the front of items and caps the result
// to limit (limit <=0 means no cap). Safe for offsets/limits past the end.
func windowSlice[T any](items []T, offset, limit int) []T {
	offset = max(offset, 0)
	if offset >= len(items) {
		return []T{}
	}
	items = items[offset:]
	if limit > 0 && limit < len(items) {
		items = items[:limit]
	}
	return items
}

// headerInt parses an integer-valued response header, returning 0 when the
// header is absent or unparseable.
func headerInt(h http.Header, key string) int {
	if h == nil {
		return 0
	}
	n, err := strconv.Atoi(strings.TrimSpace(h.Get(key)))
	if err != nil {
		return 0
	}
	return n
}
