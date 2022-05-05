package opts

import "strings"

func GetHelp() string {
	help := []string{
		"OPTIONS:",
		"\t-u, --url: URL format (default)",
		"\t-x, --xml: XML format",
		"\t-s, --slug: Slug format",
		"",
		"\t-e, --encode: Encode (default)",
		"\t-d, --decode: Decode",
		"",
		"\t-a, --all: All chars",
		"",
		"\t-v, --version: Show version",
		"\t-h, --help: Show help",
	}
	return strings.Join(help, "\n")
}
