package pcloud

import (
	"context"
	"fmt"
	"path"
	"strconv"

	pc "github.com/cstrahan/packagecloud-go/sdk"
)

// ListMasterTokens returns the master tokens for a repository (each with its
// nested read tokens).
func (a *App) ListMasterTokens(ctx context.Context, repository string) ([]*pc.MasterToken, error) {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return nil, err
	}
	return a.sdk.MasterTokens.Index(ctx, &pc.MasterTokensIndexRequest{UserID: user, Repo: repo})
}

// CreateMasterToken creates a master token with the given name.
func (a *App) CreateMasterToken(ctx context.Context, repository, name string) (*pc.MasterToken, error) {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return nil, err
	}
	return a.sdk.MasterTokens.Create(ctx, &pc.MasterTokensCreateRequest{
		UserID:          user,
		Repo:            repo,
		MasterTokenName: pc.String(name),
	})
}

// DestroyMasterToken deletes a master token by name. The "default" master
// token cannot be deleted.
func (a *App) DestroyMasterToken(ctx context.Context, repository, name string) error {
	if name == "default" {
		return fmt.Errorf("the default master token cannot be deleted")
	}
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return err
	}
	_, id, err := a.resolveMasterToken(ctx, repository, name)
	if err != nil {
		return err
	}
	return a.sdk.MasterTokens.Destroy(ctx, &pc.MasterTokensDestroyRequest{
		UserID: user,
		Repo:   repo,
		ID:     id,
	})
}

// ListReadTokens returns the read tokens nested under a named master token.
func (a *App) ListReadTokens(ctx context.Context, repository, masterTokenName string) ([]*pc.ReadToken, error) {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return nil, err
	}
	_, id, err := a.resolveMasterToken(ctx, repository, masterTokenName)
	if err != nil {
		return nil, err
	}
	resp, err := a.sdk.ReadTokens.Index(ctx, &pc.ReadTokensIndexRequest{
		UserID:      user,
		Repo:        repo,
		MasterToken: strconv.Itoa(id),
	})
	if err != nil {
		return nil, err
	}
	return resp.ReadTokens, nil
}

// CreateReadToken creates a read token with the given name under a named
// master token.
func (a *App) CreateReadToken(ctx context.Context, repository, masterTokenName, name string) (*pc.ReadToken, error) {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return nil, err
	}
	_, id, err := a.resolveMasterToken(ctx, repository, masterTokenName)
	if err != nil {
		return nil, err
	}
	return a.sdk.ReadTokens.Create(ctx, &pc.ReadTokensCreateRequest{
		UserID:        user,
		Repo:          repo,
		MasterToken:   strconv.Itoa(id),
		ReadTokenName: pc.String(name),
	})
}

// DestroyReadToken deletes a read token (by name) under a named master token.
func (a *App) DestroyReadToken(ctx context.Context, repository, masterTokenName, readTokenName string) error {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return err
	}
	_, masterID, err := a.resolveMasterToken(ctx, repository, masterTokenName)
	if err != nil {
		return err
	}
	readTokens, err := a.ListReadTokens(ctx, repository, masterTokenName)
	if err != nil {
		return err
	}
	var readID *int
	for _, rt := range readTokens {
		if rt.Name != nil && *rt.Name == readTokenName {
			readID = rt.ID
			break
		}
	}
	if readID == nil {
		return fmt.Errorf("no read token named %q under master token %q", readTokenName, masterTokenName)
	}
	return a.sdk.ReadTokens.Destroy(ctx, &pc.ReadTokensDestroyRequest{
		UserID:      user,
		Repo:        repo,
		MasterToken: strconv.Itoa(masterID),
		ID:          *readID,
	})
}

// resolveMasterToken finds a master token by name and returns it along with
// its numeric id, which the API exposes only as the trailing segment of
// paths.self (e.g. ".../master_tokens/54184").
func (a *App) resolveMasterToken(ctx context.Context, repository, name string) (*pc.MasterToken, int, error) {
	tokens, err := a.ListMasterTokens(ctx, repository)
	if err != nil {
		return nil, 0, err
	}
	for _, tok := range tokens {
		if tok.Name == nil || *tok.Name != name {
			continue
		}
		if tok.Paths == nil || tok.Paths.Self == nil {
			return nil, 0, fmt.Errorf("master token %q has no self path; cannot determine its id", name)
		}
		id, err := strconv.Atoi(path.Base(*tok.Paths.Self))
		if err != nil {
			return nil, 0, fmt.Errorf("could not parse master token id from %q: %w", *tok.Paths.Self, err)
		}
		return tok, id, nil
	}
	return nil, 0, fmt.Errorf("no master token named %q in %s", name, repository)
}
