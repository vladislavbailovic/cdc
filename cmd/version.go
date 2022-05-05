package cmd

import (
	"fmt"
	"strings"
)

var GitCommitHash string = "development"
var BuildDate string = "right now"

func GetVersion() string {
	return strings.Join([]string{
		fmt.Sprintf("Version:    %s", GitCommitHash),
		fmt.Sprintf("Build date: %s", BuildDate),
	}, "\n")
}
