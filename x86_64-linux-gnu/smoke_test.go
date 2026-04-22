//go:build smoketest

package x86_64_linux_gnu

import "testing"

func TestSmoke(t *testing.T) {
	if !smoke() {
		t.Fatal("smoke failed")
	}
}
