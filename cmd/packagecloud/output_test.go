package main

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// newTestCmd returns a command with stdout/stderr captured into the returned
// buffers.
func newTestCmd() (*cobra.Command, *bytes.Buffer, *bytes.Buffer) {
	cmd := &cobra.Command{}
	var out, errBuf bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&errBuf)
	return cmd, &out, &errBuf
}

func TestEmitRoutesDataToStdout(t *testing.T) {
	t.Cleanup(func() { flagJSON = false })

	// JSON mode: the document goes to stdout, nothing to stderr.
	flagJSON = true
	cmd, out, errBuf := newTestCmd()
	if err := emit(cmd, map[string]int{"n": 1}, func(w io.Writer) { io.WriteString(w, "HUMAN") }); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), `"n":1`) {
		t.Errorf("stdout = %q, want JSON", out.String())
	}
	if strings.Contains(out.String(), "HUMAN") {
		t.Error("human renderer should not run in JSON mode")
	}
	if errBuf.Len() != 0 {
		t.Errorf("stderr = %q, want empty", errBuf.String())
	}

	// Human mode: the rows go to stdout.
	flagJSON = false
	cmd, out, errBuf = newTestCmd()
	if err := emit(cmd, map[string]int{"n": 1}, func(w io.Writer) { io.WriteString(w, "HUMAN") }); err != nil {
		t.Fatal(err)
	}
	if out.String() != "HUMAN" {
		t.Errorf("stdout = %q, want HUMAN", out.String())
	}
	if errBuf.Len() != 0 {
		t.Errorf("stderr = %q, want empty", errBuf.String())
	}
}

func TestStatusGoesToStderrAndIsSuppressedInJSON(t *testing.T) {
	t.Cleanup(func() { flagJSON = false })

	flagJSON = false
	cmd, out, errBuf := newTestCmd()
	status(cmd, "working %d\n", 1)
	if out.Len() != 0 {
		t.Errorf("stdout = %q, want empty", out.String())
	}
	if errBuf.String() != "working 1\n" {
		t.Errorf("stderr = %q, want %q", errBuf.String(), "working 1\n")
	}

	// JSON mode: status is suppressed entirely.
	flagJSON = true
	cmd, out, errBuf = newTestCmd()
	status(cmd, "working\n")
	if out.Len() != 0 || errBuf.Len() != 0 {
		t.Errorf("status emitted output in JSON mode: stdout=%q stderr=%q", out.String(), errBuf.String())
	}
}

func TestConfirmRouting(t *testing.T) {
	t.Cleanup(func() { flagJSON = false })

	// Human mode: confirmation is status → stderr, stdout stays empty.
	flagJSON = false
	cmd, out, errBuf := newTestCmd()
	if err := confirm(cmd, map[string]bool{"success": true}, "done: %s", "x"); err != nil {
		t.Fatal(err)
	}
	if out.Len() != 0 {
		t.Errorf("stdout = %q, want empty", out.String())
	}
	if errBuf.String() != "done: x\n" {
		t.Errorf("stderr = %q, want %q", errBuf.String(), "done: x\n")
	}

	// JSON mode: the success object goes to stdout.
	flagJSON = true
	cmd, out, errBuf = newTestCmd()
	if err := confirm(cmd, map[string]bool{"success": true}, "done"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.String(), `"success":true`) {
		t.Errorf("stdout = %q, want JSON", out.String())
	}
	if errBuf.Len() != 0 {
		t.Errorf("stderr = %q, want empty", errBuf.String())
	}
}

func TestRenderError(t *testing.T) {
	var buf bytes.Buffer
	renderError(&buf, errors.New("boom"), true)
	if !strings.Contains(buf.String(), `"error":"boom"`) {
		t.Errorf("json error = %q", buf.String())
	}
	buf.Reset()
	renderError(&buf, errors.New("boom"), false)
	if buf.String() != "error: boom\n" {
		t.Errorf("text error = %q", buf.String())
	}
}

func TestWriteJSONDoesNotEscapeHTML(t *testing.T) {
	var buf bytes.Buffer
	if err := writeJSON(&buf, map[string]string{"u": "a?x=1&y=2"}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "a?x=1&y=2") {
		t.Errorf("URL was escaped: %q", buf.String())
	}
}
