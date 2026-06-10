# packagecloud-go

A Go client SDK and CLI for the [PackageCloud](https://packagecloud.io) API.

- **`sdk/`** — a Go client SDK generated from `openapi.yaml` with
  [Fern](https://buildwithfern.com). Code generated; do not edit by hand.
- **`cmd/packagecloud/`** — a CLI built on top of the generated SDK.
- **`internal/pcloud/`** — a thin wrapper around the SDK that adds the
  conveniences the generated client deliberately leaves out: config/auth
  loading, concurrent pagination, and `distro_version` resolution.

## CLI

### Install

```sh
go build -o packagecloud ./cmd/packagecloud
```

### Authentication

The CLI resolves the API token in priority order:

1. `--token` flag
2. `PACKAGECLOUD_TOKEN` environment variable
3. `~/.packagecloud` config file (`{"url": "...", "token": "..."}`)

Get your token at <https://packagecloud.io/api_token>.

### Commands

```sh
# Packages
packagecloud distributions                                    # list supported distributions
packagecloud list <user/repo>                                 # list all packages (concurrent pagination)
packagecloud list <user/repo> --limit 20 --offset 40          # window results (fetches only needed pages)
packagecloud list <user/repo> --concurrency 4                 # cap simultaneous page requests (default 10)
packagecloud search <user/repo> -q QUERY -f deb               # search packages
packagecloud push <user/repo> ./dist/*.deb -d ubuntu/xenial   # multiple files (shell-expanded)
packagecloud push <user/repo> ./*.deb --skip-duplicates --skip-errors
packagecloud download <user/repo> pkg_1.0_amd64.deb -d ubuntu/xenial   # → ./pkg_1.0_amd64.deb
packagecloud download <user/repo> pkg.gem --output-dir /tmp            # into a directory, keep name
packagecloud download <user/repo> pkg.gem -o /tmp/x.gem               # exact path ("-" = stdout)
# download shows a progress bar on stderr when stderr is a terminal (suppressed when piped or --json)
packagecloud delete <user/repo> pkg.gem                       # yank a package (.gem distro inferred)
packagecloud promote <src> <dst> ubuntu/xenial pkg_1.0_amd64.deb
packagecloud contents <user/repo> ./pkg.dsc -d ubuntu/xenial

# Repositories
packagecloud repo list [--include-collaborations]
packagecloud repo create <name> [--private]
packagecloud repo show <user/repo>

# GPG keys
packagecloud gpg-key list <user/repo>
packagecloud gpg-key create <user/repo> ./signing-key.gpg
packagecloud gpg-key destroy <user/repo> <keyname>

# Master tokens (and their nested read tokens)
packagecloud master-token list <user/repo>
packagecloud master-token create <user/repo> <name>
packagecloud master-token destroy <user/repo> <name>

# Read tokens (scoped under a master token)
packagecloud read-token list <user/repo> <master-token-name>
packagecloud read-token create <user/repo> <master-token-name> <name>
packagecloud read-token destroy <user/repo> <master-token-name> <read-token-name>
```

Add `--json` to any command for machine-readable output. Run
`packagecloud <command> --help` for per-command flags (notably the several
accepted spellings of `--distro-version`).

### Output streams

The CLI keeps a strict stdout/stderr split so it composes in pipelines:

- **stdout** carries the *result* — the JSON document under `--json`, or the
  human-readable data rows otherwise. Nothing else is written there.
- **stderr** carries progress, section headers, status confirmations, and
  errors. Under `--json`, errors are emitted there as `{"error": "..."}`.

So `packagecloud --json list user/repo | jq` always sees either a valid
document (exit 0) or empty stdout (exit ≠ 0). For multi-file `push`, the
result is a per-file array (`{"file","status":"pushed|skipped|failed",...}`),
and the exit code is non-zero if any file failed — including under
`--skip-errors`.

#### A note on `list`/`search` and repository access

`list` and `search` take a fully-qualified `user/repo`. PackageCloud has no
API to enumerate an organization's repositories — `GET /api/v1/repos` only
returns repos owned by the **authenticated user's own** namespace (it returns
an empty list for a service account that merely has collaborator access to an
org's repos). Access org repos directly by their `user/repo` name.

## Regenerating the SDK

The SDK is generated from [`openapi.yaml`](openapi.yaml). Fern config lives in
[`fern/`](fern/). Local generation runs the Go generator in Docker, so Docker
must be running.

```sh
fern generate --group go --local
```

The OpenAPI spec itself is produced by a separate project
(`packagecloud-openapi`) that scrapes the public API docs; structural
corrections to the spec belong there (in its `schema-overrides.yaml`), not in
this repo's copy. After regenerating the spec upstream, copy it here and
re-run the command above.
