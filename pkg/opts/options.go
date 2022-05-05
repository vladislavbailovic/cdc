package opts

import "strings"

type Options struct {
	options Flag
	rest    []string
}

func New(flags []string) Options {
	options := addFlag(Encode, Url)
	rest := []string{}
	for _, f := range flags {
		switch f {
		case "-u", "--url":
			options = clearFlag(options, Xml)
			options = clearFlag(options, Slug)
			options = addFlag(options, Url)
		case "-x", "--xml":
			options = clearFlag(options, Url)
			options = clearFlag(options, Slug)
			options = addFlag(options, Xml)
		case "-s", "--slug":
			options = clearFlag(options, Url)
			options = clearFlag(options, Xml)
			options = addFlag(options, Slug)
		case "-e", "--encode":
			options = clearFlag(options, Decode)
			options = addFlag(options, Encode)
		case "-d", "--decode":
			options = clearFlag(options, Encode)
			options = addFlag(options, Decode)
		case "-a", "--all":
			options = addFlag(options, AllChars)
		case "-h", "--help":
			options = addFlag(options, Help)
		case "-v", "--version":
			options = addFlag(options, Version)
		default:
			rest = append(rest, f)
		}
	}
	return Options{
		options: options,
		rest:    rest,
	}
}

func (x Options) HasFlag(which Flag) bool {
	return x.options.hasFlag(which)
}

func (x Options) Subject() string {
	return strings.Join(x.rest, " ")
}
