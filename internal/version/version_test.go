package version

import "testing"

func TestVersion(t *testing.T) {
	version := Version()
	if version != VersionText {
		t.Errorf("Version() is %q, expected %q", version, VersionText)
	}
}
