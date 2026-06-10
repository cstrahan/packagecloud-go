package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
	"golang.org/x/term"
)

func newDownloadCmd() *cobra.Command {
	var distroVersion, output, outputDir string
	cmd := &cobra.Command{
		Use:   "download <repository> <filename>",
		Short: "Download a package from a repository",
		Long: `Download a package from a repository.

By default the package is written to the current directory under its own
filename. Use --output-dir to choose the directory (keeping the filename), or
--output to set an exact path ("-" streams to stdout).

Use --distro-version to disambiguate when the same filename exists under more
than one distribution (e.g. ubuntu/xenial, python/1).`,
		Args: exactArgs("repository", "filename"),
		RunE: func(cmd *cobra.Command, args []string) error {
			if output != "" && outputDir != "" {
				return fmt.Errorf("--output and --output-dir are mutually exclusive")
			}
			app, err := newApp()
			if err != nil {
				return err
			}
			url, name, err := app.FindDownloadURL(cmd.Context(), args[0], args[1], distroVersion)
			if err != nil {
				return err
			}
			body, size, err := app.OpenDownload(cmd.Context(), url)
			if err != nil {
				return err
			}
			defer body.Close()

			// A progress bar only makes sense on an interactive terminal; it
			// goes to stderr so it never mingles with `-o -` stdout output, and
			// is suppressed under --json.
			progress := !flagJSON && term.IsTerminal(int(os.Stderr.Fd()))

			// Stream straight to stdout when asked; the file bytes are the result.
			if output == "-" {
				_, err := streamDownload(cmd.OutOrStdout(), body, size, name, progress)
				return err
			}

			dest := output
			if dest == "" {
				dir := outputDir
				if dir == "" {
					dir = "."
				} else if err := os.MkdirAll(dir, 0o755); err != nil {
					return err
				}
				dest = filepath.Join(dir, name)
			}

			f, err := os.Create(dest)
			if err != nil {
				return err
			}
			n, err := streamDownload(f, body, size, name, progress)
			if closeErr := f.Close(); err == nil {
				err = closeErr
			}
			if err != nil {
				os.Remove(dest) // don't leave a truncated/partial file behind
				return err
			}
			return confirm(cmd,
				map[string]any{"file": dest, "bytes": n},
				"Downloaded %s (%d bytes)", dest, n)
		},
	}
	cmd.Flags().StringVarP(&distroVersion, "distro-version", "d", "",
		"distribution/version to disambiguate (e.g. ubuntu/xenial, python/1)")
	cmd.Flags().StringVarP(&output, "output", "o", "",
		"write to this path; \"-\" streams to stdout")
	cmd.Flags().StringVar(&outputDir, "output-dir", "",
		"directory to download into (default current dir), keeping the filename")
	return cmd
}

// streamDownload copies r into w, returning the byte count. When progress is
// true it renders a byte progress bar to stderr, removed on completion. mpb
// measures the terminal and truncates the line to its width rather than
// wrapping, so a long filename or narrow window can't make it spew newlines.
// A spinner is shown when the size is unknown (-1).
func streamDownload(w io.Writer, r io.Reader, size int64, desc string, progress bool) (int64, error) {
	if !progress {
		return io.Copy(w, r)
	}

	p := mpb.New(mpb.WithOutput(os.Stderr), mpb.WithWidth(48))
	name := decor.Name(truncateName(desc, 30)+" ", decor.WC{C: decor.DindentRight})
	counters := decor.CountersKibiByte("% .1f / % .1f ")
	speed := decor.AverageSpeed(decor.SizeB1024(0), " % .1f")

	var bar *mpb.Bar
	if size > 0 {
		bar = p.New(size, mpb.BarStyle(),
			mpb.BarRemoveOnComplete(),
			mpb.PrependDecorators(name, counters),
			mpb.AppendDecorators(decor.Percentage(decor.WC{W: 5}), speed),
		)
	} else {
		bar = p.New(0, mpb.SpinnerStyle(),
			mpb.BarRemoveOnComplete(),
			mpb.PrependDecorators(name, counters),
			mpb.AppendDecorators(speed),
		)
	}

	proxy := bar.ProxyReader(r)
	n, err := io.Copy(w, proxy)
	proxy.Close()
	bar.SetTotal(n, true) // mark complete (covers unknown/short content-length)
	p.Wait()
	return n, err
}

// truncateName shortens s to max runes with an ellipsis, so a long filename
// can't crowd out the progress stats.
func truncateName(s string, max int) string {
	rs := []rune(s)
	if len(rs) <= max {
		return s
	}
	return string(rs[:max-1]) + "…"
}
