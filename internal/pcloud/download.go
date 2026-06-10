package pcloud

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	pc "github.com/cstrahan/packagecloud-go/sdk"
)

// FindDownloadURL locates a package by filename and returns its download URL
// (which carries an embedded read token) and the server-side filename.
//
// The same filename can exist under several distro/versions; pass
// distroVersion (matching the value shown by `list`/`search`, e.g.
// "ubuntu/xenial" or "python/1") to disambiguate. A zero match is an error;
// an ambiguous match lists the candidates so the caller can narrow it.
func (a *App) FindDownloadURL(ctx context.Context, repository, filename, distroVersion string) (url, name string, err error) {
	base := path.Base(filename)
	matches, err := a.SearchPackages(ctx, repository, SearchParams{Query: base}, ListOptions{})
	if err != nil {
		return "", "", err
	}
	return selectDownload(matches, repository, base, distroVersion)
}

// selectDownload picks the single package whose filename exactly equals base
// (optionally also matching distroVersion) from a search result set, and
// returns its download URL and filename. It is the pure core of
// FindDownloadURL, split out so the not-found / exact / ambiguous branches are
// unit-testable without the network.
func selectDownload(matches []*pc.PackageFragment, repository, base, distroVersion string) (url, name string, err error) {
	var hits []*pc.PackageFragment
	for _, p := range matches {
		if p.Filename != base {
			continue // search is a substring match; require the exact filename
		}
		if distroVersion != "" && p.DistroVersion != distroVersion {
			continue
		}
		hits = append(hits, p)
	}

	switch len(hits) {
	case 0:
		if distroVersion != "" {
			return "", "", fmt.Errorf("no package %q in %s for distro %q", base, repository, distroVersion)
		}
		return "", "", fmt.Errorf("no package %q found in %s", base, repository)
	case 1:
		return hits[0].DownloadURL, hits[0].Filename, nil
	default:
		dvs := make([]string, len(hits))
		for i, h := range hits {
			dvs[i] = h.DistroVersion
		}
		return "", "", fmt.Errorf(
			"%q matches packages in multiple distros (%s); disambiguate with --distro-version",
			base, strings.Join(dvs, ", "))
	}
}

// OpenDownload issues the GET (following redirects) and returns the response
// body and its content length (-1 if unknown). The caller must close the body.
// Returning the stream rather than copying it lets the CLI layer drive an
// optional progress bar.
func (a *App) OpenDownload(ctx context.Context, url string) (io.ReadCloser, int64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}
	// The download URL already authorizes via its read_token; basic auth is a
	// harmless belt-and-suspenders for the initial packagecloud.io request.
	// (Go strips it automatically on a cross-host redirect to e.g. S3.)
	req.SetBasicAuth(a.token, "")

	resp, err := a.http.Do(req)
	if err != nil {
		return nil, 0, err
	}
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		resp.Body.Close()
		return nil, 0, fmt.Errorf("download failed (status %d): %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	return resp.Body, resp.ContentLength, nil
}
