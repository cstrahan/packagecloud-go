package pcloud

import (
	"context"
	"fmt"
	"os"

	pc "github.com/cstrahan/packagecloud-go/sdk"
)

// ListGpgKeys returns the GPG keys for a repository.
func (a *App) ListGpgKeys(ctx context.Context, repository string) ([]*pc.GpgKey, error) {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return nil, err
	}
	resp, err := a.sdk.GpgKeys.Index(ctx, &pc.GpgKeysIndexRequest{UserID: user, Repo: repo})
	if err != nil {
		return nil, err
	}
	return resp.GpgKeys, nil
}

// CreateGpgKey uploads a package-signing GPG key from a file on disk.
func (a *App) CreateGpgKey(ctx context.Context, repository, keyPath string) (*pc.GpgKey, error) {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(keyPath)
	if err != nil {
		return nil, fmt.Errorf("opening GPG key file: %w", err)
	}
	defer f.Close()

	return a.sdk.GpgKeys.Create(ctx, &pc.GpgKeysCreateRequest{
		UserID:        user,
		Repo:          repo,
		GpgKeyKeydata: f,
	})
}

// DestroyGpgKey deletes a package-signing GPG key by name. Repository signing
// keys cannot be deleted (the API returns 404).
func (a *App) DestroyGpgKey(ctx context.Context, repository, keyname string) error {
	user, repo, err := SplitRepo(repository)
	if err != nil {
		return err
	}
	_, err = a.sdk.GpgKeys.Destroy(ctx, &pc.GpgKeysDestroyRequest{
		UserID:  user,
		Repo:    repo,
		Keyname: keyname,
	})
	return err
}
