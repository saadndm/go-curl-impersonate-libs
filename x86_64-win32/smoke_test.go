//go:build smoketest

package x86_64_win32

import "testing"

func TestSmoke(t *testing.T) {
	if !smoke() {
		t.Fatal("Smoke test failed")
	}
}
