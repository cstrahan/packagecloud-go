package pcloud

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// DeletePackage yanks a package from a repository.
//
// PackageCloud's delete URL is polymorphic in the package type — the path
// segment between the repo and the filename varies (deb/rpm use
// "distro/version", gems use "gems", Maven uses "java/maven2/<group>", etc.).
// The generated SDK exposes three separate typed endpoints that each split
// the filename into "<package>.<ext>", which can't represent multi-extension
// names like "foo-1.0.tar.gz". So we issue the request directly, mirroring
// the reference client's URL shaping, while reusing the same auth.
//
// distroVersion is the path segment(s); pass "" for a .gem file to default to
// "gems". scope carries the npm scope for scoped Node.js packages.
func (a *App) DeletePackage(ctx context.Context, repository, distroVersion, filename, scope string) error {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return err
	}
	endpoint := buildDeleteURL(a.baseURL, user, repo, distroVersion, filename)

	if scope != "" {
		q := url.Values{}
		q.Set("scope", scope)
		endpoint += "?" + q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(a.token, "")
	req.Header.Set("Accept", "application/json")

	resp, err := a.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("delete failed (status %d): %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	return nil
}

// buildDeleteURL assembles the DELETE URL from the caller-facing
// (distroVersion, filename) shape. Only the basename of filename is used;
// a .gem with an empty distroVersion defaults to the "gems" segment, matching
// the Ruby/Rust CLI's yank convenience.
func buildDeleteURL(baseURL, user, repo, distroVersion, filename string) string {
	name := path.Base(filename)
	if distroVersion == "" && strings.HasSuffix(name, ".gem") {
		distroVersion = "gems"
	} else {
		distroVersion = strings.Trim(distroVersion, "/")
	}
	return fmt.Sprintf("%s/api/v1/repos/%s/%s/%s/%s",
		strings.TrimRight(baseURL, "/"), user, repo, distroVersion, name)
}
