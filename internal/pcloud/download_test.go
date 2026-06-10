package pcloud

import (
	"strings"
	"testing"

	pc "github.com/cstrahan/packagecloud-go/sdk"
)

func frag(filename, distroVersion, downloadURL string) *pc.PackageFragment {
	return &pc.PackageFragment{Filename: filename, DistroVersion: distroVersion, DownloadURL: downloadURL}
}

func TestSelectDownload(t *testing.T) {
	matches := []*pc.PackageFragment{
		frag("foo_1.0_amd64.deb", "ubuntu/xenial", "https://x/xenial"),
		frag("foo_1.0_amd64.deb", "ubuntu/bionic", "https://x/bionic"),
		frag("foo_1.0_amd64.deb.sig", "ubuntu/xenial", "https://x/sig"), // substring, must be excluded
		frag("bar_2.0_amd64.deb", "ubuntu/xenial", "https://x/bar"),
	}

	// Unambiguous when narrowed by distro.
	url, name, err := selectDownload(matches, "u/r", "foo_1.0_amd64.deb", "ubuntu/bionic")
	if err != nil || url != "https://x/bionic" || name != "foo_1.0_amd64.deb" {
		t.Fatalf("got (%q, %q, %v), want the bionic match", url, name, err)
	}

	// Ambiguous across distros without a distro filter.
	_, _, err = selectDownload(matches, "u/r", "foo_1.0_amd64.deb", "")
	if err == nil || !strings.Contains(err.Error(), "multiple distros") {
		t.Errorf("ambiguous match error = %v, want it to mention multiple distros", err)
	}

	// Not found (filename absent).
	_, _, err = selectDownload(matches, "u/r", "nope.deb", "")
	if err == nil || !strings.Contains(err.Error(), "no package") {
		t.Errorf("missing match error = %v", err)
	}

	// Not found for the given distro specifically.
	_, _, err = selectDownload(matches, "u/r", "bar_2.0_amd64.deb", "ubuntu/focal")
	if err == nil || !strings.Contains(err.Error(), "focal") {
		t.Errorf("wrong-distro error = %v, want it to mention the distro", err)
	}

	// Single exact match (one distro) needs no disambiguation.
	url, _, err = selectDownload(matches, "u/r", "bar_2.0_amd64.deb", "")
	if err != nil || url != "https://x/bar" {
		t.Fatalf("got (%q, %v), want the single bar match", url, err)
	}
}
