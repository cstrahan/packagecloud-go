package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// Output discipline for this CLI:
//
//   - stdout carries the *result* a downstream tool consumes: the JSON
//     document in --json mode, or the human-readable data rows otherwise.
//   - stderr carries everything else: progress, section headers, status
//     confirmations, warnings, and errors.
//
// All command output routes through cmd.OutOrStdout()/cmd.ErrOrStderr() so it
// can be redirected in tests.

// writeJSON encodes v as a single JSON line to w. HTML escaping is disabled so
// the URLs in responses (which contain &, <, >) stay readable.
func writeJSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

// emit renders a command's result payload to stdout. In --json mode it writes
// jsonValue; otherwise it invokes human to write the human-readable rows.
func emit(cmd *cobra.Command, jsonValue any, human func(w io.Writer)) error {
	if flagJSON {
		return writeJSON(cmd.OutOrStdout(), jsonValue)
	}
	human(cmd.OutOrStdout())
	return nil
}

// confirm reports the outcome of an action that has no data payload (delete,
// destroy, …). In --json mode it writes jsonValue to stdout; otherwise it
// writes a human confirmation to stderr (status, not data).
func confirm(cmd *cobra.Command, jsonValue any, format string, args ...any) error {
	if flagJSON {
		return writeJSON(cmd.OutOrStdout(), jsonValue)
	}
	fmt.Fprintf(cmd.ErrOrStderr(), format+"\n", args...)
	return nil
}

// status writes a human-mode progress/header line to stderr. It is suppressed
// in --json mode so stderr stays reserved for errors. No trailing newline is
// added — include it in format when you want one.
func status(cmd *cobra.Command, format string, args ...any) {
	if !flagJSON {
		fmt.Fprintf(cmd.ErrOrStderr(), format, args...)
	}
}

// renderError writes err to w: a JSON object when asJSON, plain text otherwise.
func renderError(w io.Writer, err error, asJSON bool) {
	if asJSON {
		_ = writeJSON(w, map[string]string{"error": err.Error()})
		return
	}
	fmt.Fprintln(w, "error:", err)
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func derefInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
