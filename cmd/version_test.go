package cmd

import (
	"strings"
	"testing"
)

func Test_GetVersion(t *testing.T) {
	v := GetVersion()
	if !strings.Contains(v, "Version") {
		t.Fatalf("version info missing")
	}
	if !strings.Contains(v, "Build") {
		t.Fatalf("build info missing")
	}
}
