//go:build smoketest

package x86_64_linux_gnu

/*
char *curl_version(void);
*/
import "C"

func smoke() bool {
	return C.curl_version() != nil
}
