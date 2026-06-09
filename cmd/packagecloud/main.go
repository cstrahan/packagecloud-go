// Command packagecloud is a CLI for the PackageCloud API, built on the
// fern-generated SDK in the sibling `sdk` module.
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/cstrahan/packagecloud-go/internal/pcloud"
	pc "github.com/cstrahan/packagecloud-go/sdk"
)

// global flags shared by every subcommand.
var (
	flagToken  string
	flagConfig string
	flagJSON   bool
)

// Build metadata, overridden at release time via -ldflags (see .goreleaser.yaml).
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if err := newRootCmd().Execute(); err != nil {
		renderError(os.Stderr, err, flagJSON)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "packagecloud",
		Short:         "Command-line interface for the PackageCloud API",
		Version:       version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.SetVersionTemplate("packagecloud {{.Version}} (commit " + commit + ", built " + date + ")\n")

	root.PersistentFlags().StringVar(&flagToken, "token", "",
		"API token (overrides PACKAGECLOUD_TOKEN and the config file)")
	root.PersistentFlags().StringVar(&flagConfig, "config", "",
		"config file path (defaults to ~/.packagecloud)")
	root.PersistentFlags().BoolVar(&flagJSON, "json", false,
		"output results as JSON instead of human-readable text")

	root.AddCommand(
		newDistributionsCmd(),
		newPushCmd(),
		newListCmd(),
		newSearchCmd(),
		newDeleteCmd(),
		newPromoteCmd(),
		newContentsCmd(),
		newRepoCmd(),
		newGpgKeyCmd(),
		newMasterTokenCmd(),
		newReadTokenCmd(),
	)
	return root
}

// newApp constructs the client wrapper from the global flags.
func newApp() (*pcloud.App, error) {
	return pcloud.New(flagToken, flagConfig)
}

func newDistributionsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "distributions",
		Short: "List all supported distributions",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			dists, err := app.Distributions(cmd.Context())
			if err != nil {
				return err
			}
			status(cmd, "Supported distributions:\n")
			return emit(cmd, dists, func(w io.Writer) {
				first := true
				for pkgType, distList := range dists {
					if !first {
						fmt.Fprintln(w)
					}
					first = false
					fmt.Fprintf(w, "%s:\n", pkgType)
					for _, dist := range distList {
						fmt.Fprintf(w, "  %s (%s)\n", deref(dist.DisplayName), deref(dist.IndexName))
						for _, ver := range dist.Versions {
							label := ver.DisplayName
							if ver.VersionNumber != nil {
								label = *ver.VersionNumber
							} else if ver.IndexName != nil {
								label = *ver.IndexName
							}
							fmt.Fprintf(w, "    - %s (ID: %d)\n", label, ver.ID)
						}
					}
				}
			})
		},
	}
}

// pushResult is the per-file outcome of a push, both the JSON document element
// and the human progress source.
type pushResult struct {
	File    string             `json:"file"`
	Status  string             `json:"status"` // pushed | skipped | failed
	Error   string             `json:"error,omitempty"`
	Package *pc.PackageDetails `json:"package,omitempty"`
}

func newPushCmd() *cobra.Command {
	var (
		distroVersion  string
		coordinates    string
		skipErrors     bool
		skipDuplicates bool
	)
	cmd := &cobra.Command{
		Use:   "push <repository> <package_file>...",
		Short: "Push one or more packages to a repository",
		Long: `Push one or more packages to a repository. Multiple files are accepted;
use your shell's glob expansion to push many at once (e.g. "dist/*.deb").

--distro-version accepts any of:
  - a numeric distro_version_id (e.g. 166)
  - distro/version (e.g. ubuntu/xenial, el/9)
  - distro/version_number (e.g. ubuntu/16.04)
  - a single-version-type shortcut (e.g. python, java) or package_type (e.g. py, jar)
Omit it for rubygems and single-version types to let the file extension pick a default.`,
		Args: cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			files := args[1:]
			opts := pcloud.PushOptions{DistroVersion: distroVersion, Coordinates: coordinates}

			results := make([]pushResult, 0, len(files))
			var pushedCount, failed int
			for _, file := range files {
				status(cmd, "Pushing %s... ", file)
				pkg, err := app.PushPackage(cmd.Context(), args[0], file, opts)
				switch {
				case err == nil:
					status(cmd, "done\n")
					pushedCount++
					results = append(results, pushResult{File: file, Status: "pushed", Package: pkg})
				case skipDuplicates && pcloud.IsDuplicateError(err):
					status(cmd, "skipped (already exists)\n")
					results = append(results, pushResult{File: file, Status: "skipped"})
				case skipErrors:
					status(cmd, "error (skipping)\n")
					failed++
					results = append(results, pushResult{File: file, Status: "failed", Error: err.Error()})
				default:
					// No --skip-errors: abort on the first hard failure with no
					// partial document (stdout stays empty on failure).
					status(cmd, "error\n")
					return fmt.Errorf("%s: %w", file, err)
				}
			}

			// Emit the per-file result document before signalling failure, so a
			// --skip-errors run still reports what happened on stdout.
			if err := emit(cmd, results, func(w io.Writer) {
				for _, r := range results {
					if r.Status == "pushed" && r.Package != nil {
						fmt.Fprintln(w, r.Package.PackageHtmlUrl)
					}
				}
			}); err != nil {
				return err
			}
			status(cmd, "\nPushed %d/%d package(s).\n", pushedCount, len(files))
			if failed > 0 {
				return fmt.Errorf("%d package(s) failed to push", failed)
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&distroVersion, "distro-version", "d", "",
		"where to route the package (see --help for accepted forms)")
	cmd.Flags().StringVar(&coordinates, "coordinates", "",
		"explicit Maven coordinates for a JAR/WAR/AAR")
	cmd.Flags().BoolVar(&skipErrors, "skip-errors", false,
		"continue pushing remaining files when one fails")
	cmd.Flags().BoolVar(&skipDuplicates, "skip-duplicates", false,
		"skip packages that already exist instead of failing")
	return cmd
}

func newListCmd() *cobra.Command {
	var perPage int
	cmd := &cobra.Command{
		Use:   "list <repository>",
		Short: "List packages in a repository",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			pkgs, err := app.ListPackages(cmd.Context(), args[0], perPage)
			if err != nil {
				return err
			}
			return printPackageFragments(cmd, pkgs)
		},
	}
	cmd.Flags().IntVar(&perPage, "per-page", 250, "results per page")
	return cmd
}

func newSearchCmd() *cobra.Command {
	var (
		query, dist, filter, arch string
		perPage                   int
	)
	cmd := &cobra.Command{
		Use:   "search <repository>",
		Short: "Search for packages in a repository",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			pkgs, err := app.SearchPackages(cmd.Context(), args[0], pcloud.SearchParams{
				Query:   query,
				Dist:    dist,
				Filter:  filter,
				Arch:    arch,
				PerPage: perPage,
			})
			if err != nil {
				return err
			}
			return printPackageFragments(cmd, pkgs)
		},
	}
	cmd.Flags().StringVarP(&query, "query", "q", "", "search query (filename substring)")
	cmd.Flags().StringVarP(&dist, "dist", "d", "", "filter by distribution (e.g. ubuntu, el/6)")
	cmd.Flags().StringVarP(&filter, "filter", "f", "", "filter by package type (rpm, deb, gem, python, node)")
	cmd.Flags().StringVarP(&arch, "arch", "a", "", "filter by architecture")
	cmd.Flags().IntVar(&perPage, "per-page", 250, "results per page")
	return cmd
}

func newDeleteCmd() *cobra.Command {
	var distroVersion, scope string
	cmd := &cobra.Command{
		Use:   "delete <repository> <filename>",
		Short: "Delete (yank) a package from a repository",
		Long: `Delete (yank) a package from a repository.

--distro-version is the path segment(s) identifying the package's location.
Omit it for .gem files (defaults to "gems"). Otherwise:
  - deb/rpm/dsc: distro/version (e.g. ubuntu/xenial, el/9)
  - python:      python/1
  - node:        node/1
  - anyfile:     anyfile/1
  - java:        java/maven2/<groupId>`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			if err := app.DeletePackage(cmd.Context(), args[0], distroVersion, args[1], scope); err != nil {
				return err
			}
			return confirm(cmd,
				map[string]any{"success": true, "filename": args[1]},
				"Deleted package: %s", args[1])
		},
	}
	cmd.Flags().StringVarP(&distroVersion, "distro-version", "d", "",
		"package location path segment(s); omit for .gem (see --help)")
	cmd.Flags().StringVar(&scope, "scope", "", "npm scope for scoped Node.js packages (e.g. @example)")
	return cmd
}

func newPromoteCmd() *cobra.Command {
	var scope string
	cmd := &cobra.Command{
		Use:   "promote <source> <destination> <distro/version> <filename>",
		Short: "Promote (move) a package between repositories",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			pkg, err := app.PromotePackage(cmd.Context(), args[0], args[1], args[2], args[3], scope)
			if err != nil {
				return err
			}
			status(cmd, "Successfully promoted package:\n")
			return emit(cmd, pkg, func(w io.Writer) {
				fmt.Fprintf(w, "Name:       %s\n", pkg.Name)
				fmt.Fprintf(w, "Version:    %s\n", pkg.Version)
				fmt.Fprintf(w, "Repository: %s\n", pkg.Repository)
			})
		},
	}
	cmd.Flags().StringVar(&scope, "scope", "", "npm scope for scoped Node.js packages (e.g. @example)")
	return cmd
}

func newContentsCmd() *cobra.Command {
	var distroVersion string
	cmd := &cobra.Command{
		Use:   "contents <repository> <package_file>",
		Short: "Inspect a Debian source package (.dsc) — list referenced files",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			contents, err := app.PackageContents(cmd.Context(), args[0], args[1], distroVersion)
			if err != nil {
				return err
			}
			var total int
			for _, c := range contents {
				total += len(c.Files)
			}
			status(cmd, "Package contents (%d files):\n", total)
			return emit(cmd, contents, func(w io.Writer) {
				for _, c := range contents {
					for _, file := range c.Files {
						fmt.Fprintf(w, "%s (%d bytes, MD5: %s)\n",
							deref(file.Filename), derefInt(file.Size), deref(file.Md5Sum))
					}
				}
			})
		},
	}
	cmd.Flags().StringVarP(&distroVersion, "distro-version", "d", "",
		"distribution and version (e.g. ubuntu/xenial); required")
	return cmd
}

// printPackageFragments renders a package list: a count header to stderr and
// the rows (or JSON) to stdout. Shared by `list` and `search`.
func printPackageFragments(cmd *cobra.Command, pkgs []*pc.PackageFragment) error {
	status(cmd, "Found %d packages:\n", len(pkgs))
	return emit(cmd, pkgs, func(w io.Writer) {
		for _, p := range pkgs {
			fmt.Fprintf(w, "%s %s (%s) - %s\n", p.Name, p.Version, p.DistroVersion, p.Filename)
		}
	})
}
