//go:build smoketest

package arm64_macos

import "testing"

func TestSmoke(t *testing.T) {
	if !smoke() {
		t.Fatal("smoke failed")
	}
}
