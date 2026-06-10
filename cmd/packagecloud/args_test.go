package main

import (
	"strings"
	"testing"
)

func TestExactArgs(t *testing.T) {
	v := exactArgs("repository")

	if err := v(nil, []string{"acme/widgets"}); err != nil {
		t.Errorf("exact count should pass, got %v", err)
	}
	err := v(nil, nil)
	if err == nil || !strings.Contains(err.Error(), "<repository>") {
		t.Errorf("missing arg error = %v, want it to name <repository>", err)
	}
	err = v(nil, []string{"a/b", "extra"})
	if err == nil || !strings.Contains(err.Error(), "unexpected") || !strings.Contains(err.Error(), "extra") {
		t.Errorf("extra arg error = %v, want it to flag the unexpected 'extra'", err)
	}
}

func TestExactArgsNamesAllMissing(t *testing.T) {
	v := exactArgs("user/repo", "master-token-name", "name")
	err := v(nil, []string{"a/b"})
	if err == nil {
		t.Fatal("expected error for missing args")
	}
	for _, want := range []string{"<master-token-name>", "<name>"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q should name %s", err, want)
		}
	}
	if strings.Contains(err.Error(), "<user/repo>") {
		t.Errorf("error %q should not list the satisfied <user/repo>", err)
	}
}

func TestMinArgsVariadic(t *testing.T) {
	v := minArgs("repository", "package_file")

	// One repo, no file → names the missing file.
	if err := v(nil, []string{"a/b"}); err == nil || !strings.Contains(err.Error(), "<package_file>") {
		t.Errorf("error = %v, want it to name <package_file>", err)
	}
	// Repo + several files → variadic, no error.
	if err := v(nil, []string{"a/b", "x.deb", "y.deb"}); err != nil {
		t.Errorf("variadic extra files should pass, got %v", err)
	}
}
