//go:build smoketest

package aarch64_linux_gnu

import "testing"

func TestSmoke(t *testing.T) {
	if !smoke() {
		t.Fatal("smoke failed")
	}
}
