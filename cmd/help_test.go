package cmd

import (
	"strings"
	"testing"
)

func Test_GetHelp(t *testing.T) {
	h := GetHelp()
	if !strings.Contains(h, "--encode") {
		t.Fatalf("help text should contain encode flag")
	}
}
