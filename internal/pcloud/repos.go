package pcloud

import (
	"context"
	"net/http"

	pc "github.com/cstrahan/packagecloud-go/sdk"
)

// ListRepositories returns the repositories in the authenticated user's own
// namespace, fetching pages concurrently. If includeCollaborations is true,
// repositories the user collaborates on are included too.
//
// Note: this only ever lists repos under the authenticated user's namespace;
// PackageCloud has no API to enumerate an organization's repositories.
func (a *App) ListRepositories(ctx context.Context, includeCollaborations bool) ([]*pc.Repository, error) {
	return fetchAllPages(ctx, func(ctx context.Context, page int) ([]*pc.Repository, http.Header, error) {
		req := &pc.RepositoriesIndexRequest{Page: pc.Int(page)}
		if includeCollaborations {
			req.IncludeCollaborations = pc.String("true")
		}
		resp, err := a.sdk.Repositories.WithRawResponse.Index(ctx, req)
		if err != nil {
			return nil, nil, err
		}
		return resp.Body, resp.Header, nil
	})
}

// CreateRepository creates a repository. name may be a bare "repo" (created in
// the authenticated user's namespace) or a fully-qualified "user/repo".
func (a *App) CreateRepository(ctx context.Context, name string, private bool) (*pc.RepositoriesCreateResponse, error) {
	req := &pc.RepositoriesCreateRequest{
		Repository: &pc.RepositoriesCreateRequestRepository{Name: pc.String(name)},
	}
	if private {
		// The API takes repository[private] as a string ("true"/"false").
		req.Repository.Private = pc.String("true")
	}
	return a.sdk.Repositories.Create(ctx, req)
}

// ShowRepository fetches a single repository by "user/repo".
func (a *App) ShowRepository(ctx context.Context, repository string) (*pc.Repository, error) {
	user, name, err := SplitRepo(repository)
	if err != nil {
		return nil, err
	}
	return a.sdk.Repositories.Show(ctx, &pc.RepositoriesShowRequest{UserID: user, Name: name})
}
