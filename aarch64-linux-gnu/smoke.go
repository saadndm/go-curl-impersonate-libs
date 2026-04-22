//go:build smoketest

package aarch64_linux_gnu

/*
char *curl_version(void);
*/
import "C"

func smoke() bool {
	return C.curl_version() != nil
}
