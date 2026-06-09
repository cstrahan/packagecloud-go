// Package pcloud is a thin, ergonomic wrapper around the fern-generated
// PackageCloud SDK (the `sdk` module). It owns configuration/auth loading,
// concurrent pagination, and distro_version resolution — the bits the
// generated client deliberately leaves to the caller — so the CLI in
// cmd/packagecloud can stay focused on flag parsing and output.
package pcloud

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// defaultURL is the public packagecloud.io endpoint. Mirrors the SDK's
// Environments.Default (minus the requirement to carry a trailing slash).
const defaultURL = "https://packagecloud.io/"

// Config models the ~/.packagecloud JSON file: an API URL and an optional
// token. The same shape the Ruby/Rust clients read.
type Config struct {
	// URL is the base PackageCloud URL (defaults to https://packagecloud.io/).
	URL string `json:"url"`
	// Token is the API token used for HTTP basic auth (token as username).
	Token string `json:"token,omitempty"`
}

// DefaultConfigPath returns ~/.packagecloud.
func DefaultConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}
	return filepath.Join(home, ".packagecloud"), nil
}

// LoadConfig reads configuration from path. If path is empty, it tries the
// default location (~/.packagecloud); a missing default file is not an error
// and yields a zero-value Config.
func LoadConfig(path string) (Config, error) {
	cfg := Config{URL: defaultURL}

	if path == "" {
		def, err := DefaultConfigPath()
		if err != nil {
			return cfg, err
		}
		if _, statErr := os.Stat(def); statErr != nil {
			return cfg, nil // no config file; rely on env/flag for the token
		}
		path = def
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("reading config %s: %w", path, err)
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("parsing config %s: %w", path, err)
	}
	if cfg.URL == "" {
		cfg.URL = defaultURL
	}
	return cfg, nil
}

// EffectiveToken resolves the token to use, in priority order:
// explicit flag > PACKAGECLOUD_TOKEN env var > config file. Returns "" if
// none is set.
func (c Config) EffectiveToken(provided string) string {
	if provided != "" {
		return provided
	}
	if env := os.Getenv("PACKAGECLOUD_TOKEN"); env != "" {
		return env
	}
	return c.Token
}

// EffectiveURL returns the base URL without a trailing slash (the shape the
// SDK's WithBaseURL option expects; it appends /api/v1/... itself).
func (c Config) EffectiveURL() string {
	u := c.URL
	if u == "" {
		u = defaultURL
	}
	return strings.TrimRight(u, "/")
}
