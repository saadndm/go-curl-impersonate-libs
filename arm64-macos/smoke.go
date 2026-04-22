//go:build smoketest

package arm64_macos

/*
char *curl_version(void);
*/
import "C"

func smoke() bool {
	return C.curl_version() != nil
}
