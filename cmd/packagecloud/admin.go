package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/cstrahan/packagecloud-go/internal/pcloud"
)

// ---------------------------------------------------------------------------
// repo

func newRepoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo",
		Short: "Manage repositories",
	}

	var perPage, limit, offset int
	list := &cobra.Command{
		Use:   "list",
		Short: "List repositories in your namespace",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			collaborations, _ := cmd.Flags().GetBool("include-collaborations")
			repos, err := app.ListRepositories(cmd.Context(), collaborations,
				pcloud.ListOptions{PerPage: perPage, Limit: limit, Offset: offset})
			if err != nil {
				return err
			}
			if len(repos) == 0 {
				status(cmd, "You have no repositories in your own namespace.\n")
			} else {
				status(cmd, "Found %d repositories:\n", len(repos))
			}
			return emit(cmd, repos, func(w io.Writer) {
				for _, r := range repos {
					vis := "public"
					if r.Private {
						vis = "private"
					}
					fmt.Fprintf(w, "%s (%s) — last push: %s | packages: %s\n",
						r.Fqname, vis, r.LastPushHuman, r.PackageCountHuman)
				}
			})
		},
	}
	list.Flags().Bool("include-collaborations", false, "include repos you collaborate on")
	addListFlags(list, &perPage, &limit, &offset)

	create := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a repository (bare name, or user/repo)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			private, _ := cmd.Flags().GetBool("private")
			repo, err := app.CreateRepository(cmd.Context(), args[0], private)
			if err != nil {
				return err
			}
			return confirm(cmd, repo, "Created repository: %s", deref(repo.URL))
		},
	}
	create.Flags().Bool("private", false, "make the repository private")

	show := &cobra.Command{
		Use:   "show <user/repo>",
		Short: "Show a repository",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			repo, err := app.ShowRepository(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			return emit(cmd, repo, func(w io.Writer) {
				vis := "public"
				if repo.Private {
					vis = "private"
				}
				fmt.Fprintf(w, "%s (%s)\n", repo.Fqname, vis)
				fmt.Fprintf(w, "  URL:       %s\n", repo.URL)
				fmt.Fprintf(w, "  Created:   %s\n", repo.CreatedAt)
				fmt.Fprintf(w, "  Last push: %s\n", repo.LastPushHuman)
				fmt.Fprintf(w, "  Packages:  %s\n", repo.PackageCountHuman)
			})
		},
	}

	cmd.AddCommand(list, create, show)
	return cmd
}

// ---------------------------------------------------------------------------
// gpg-key

func newGpgKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gpg-key",
		Aliases: []string{"gpg_key"},
		Short:   "Manage package-signing GPG keys",
	}

	list := &cobra.Command{
		Use:   "list <user/repo>",
		Short: "List GPG keys for a repository",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			keys, err := app.ListGpgKeys(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			status(cmd, "Found %d GPG key(s):\n", len(keys))
			return emit(cmd, keys, func(w io.Writer) {
				for _, k := range keys {
					fmt.Fprintf(w, "%s\n  type: %s\n  fingerprint: %s\n  url: %s\n",
						deref(k.Name), deref(k.Keytype), deref(k.Fingerprint), deref(k.DownloadURL))
				}
			})
		},
	}

	create := &cobra.Command{
		Use:   "create <user/repo> <path/to/key>",
		Short: "Upload a package-signing GPG key",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			key, err := app.CreateGpgKey(cmd.Context(), args[0], args[1])
			if err != nil {
				return err
			}
			return confirm(cmd, key, "Created GPG key: %s (%s)", deref(key.Name), deref(key.Fingerprint))
		},
	}

	destroy := &cobra.Command{
		Use:   "destroy <user/repo> <keyname>",
		Short: "Delete a package-signing GPG key",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			if err := app.DestroyGpgKey(cmd.Context(), args[0], args[1]); err != nil {
				return err
			}
			return confirm(cmd,
				map[string]any{"success": true, "keyname": args[1]},
				"Deleted GPG key: %s", args[1])
		},
	}

	cmd.AddCommand(list, create, destroy)
	return cmd
}

// ---------------------------------------------------------------------------
// master-token

func newMasterTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "master-token",
		Aliases: []string{"master_token"},
		Short:   "Manage master tokens",
	}

	list := &cobra.Command{
		Use:   "list <user/repo>",
		Short: "List master tokens (with their read tokens)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			tokens, err := app.ListMasterTokens(cmd.Context(), args[0])
			if err != nil {
				return err
			}
			status(cmd, "Found %d master token(s):\n", len(tokens))
			return emit(cmd, tokens, func(w io.Writer) {
				for _, t := range tokens {
					fmt.Fprintf(w, "%s (%s)\n", deref(t.Name), deref(t.Value))
					for _, rt := range t.ReadTokens {
						fmt.Fprintf(w, "    read token: %s (%s)\n", deref(rt.Name), deref(rt.Value))
					}
				}
			})
		},
	}

	create := &cobra.Command{
		Use:   "create <user/repo> <name>",
		Short: "Create a master token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			tok, err := app.CreateMasterToken(cmd.Context(), args[0], args[1])
			if err != nil {
				return err
			}
			return confirm(cmd, tok, "Created master token %s with value %s", deref(tok.Name), deref(tok.Value))
		},
	}

	destroy := &cobra.Command{
		Use:   "destroy <user/repo> <name>",
		Short: "Destroy a master token by name",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			if err := app.DestroyMasterToken(cmd.Context(), args[0], args[1]); err != nil {
				return err
			}
			return confirm(cmd,
				map[string]any{"success": true, "name": args[1]},
				"Destroyed master token: %s", args[1])
		},
	}

	cmd.AddCommand(list, create, destroy)
	return cmd
}

// ---------------------------------------------------------------------------
// read-token

func newReadTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "read-token",
		Aliases: []string{"read_token"},
		Short:   "Manage read tokens (nested under a master token)",
	}

	list := &cobra.Command{
		Use:   "list <user/repo> <master-token-name>",
		Short: "List read tokens under a master token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			tokens, err := app.ListReadTokens(cmd.Context(), args[0], args[1])
			if err != nil {
				return err
			}
			status(cmd, "Found %d read token(s) under %q:\n", len(tokens), args[1])
			return emit(cmd, tokens, func(w io.Writer) {
				for _, t := range tokens {
					fmt.Fprintf(w, "%s (%s)\n", deref(t.Name), deref(t.Value))
				}
			})
		},
	}

	create := &cobra.Command{
		Use:   "create <user/repo> <master-token-name> <name>",
		Short: "Create a read token under a master token",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			tok, err := app.CreateReadToken(cmd.Context(), args[0], args[1], args[2])
			if err != nil {
				return err
			}
			return confirm(cmd, tok, "Created read token %s with value %s", deref(tok.Name), deref(tok.Value))
		},
	}

	destroy := &cobra.Command{
		Use:   "destroy <user/repo> <master-token-name> <read-token-name>",
		Short: "Destroy a read token under a master token",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			if err := app.DestroyReadToken(cmd.Context(), args[0], args[1], args[2]); err != nil {
				return err
			}
			return confirm(cmd,
				map[string]any{"success": true, "name": args[2]},
				"Destroyed read token: %s", args[2])
		},
	}

	cmd.AddCommand(list, create, destroy)
	return cmd
}
