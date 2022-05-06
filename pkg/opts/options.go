package opts

import (
	"strings"
)

type Options struct {
	options Flag
	rest    []string
}

func New(flags []string) Options {
	options := addFlag(Encode, Url)
	rest := []string{}
	for _, f := range flags {
		if string(f[0]) == "-" {
			options = parseOption(f, options)
		} else {
			rest = append(rest, f)
		}
		continue
	}
	return Options{
		options: options,
		rest:    rest,
	}
}

func parseOption(opt string, options Flag) Flag {
	if string(opt[1]) == "-" {
		return parseLongOption(opt, options)
	}
	return parseShortOption(opt, options)
}

func parseLongOption(opt string, options Flag) Flag {
	switch opt {
	case "--url":
		options = setFlag(options, Url, Slug, Xml)
	case "--xml":
		options = setFlag(options, Xml, Url, Slug)
	case "--slug":
		options = setFlag(options, Slug, Xml, Url)
	case "--encode":
		options = setFlag(options, Encode, Decode)
	case "--decode":
		options = setFlag(options, Decode, Encode)
	case "--all":
		options = addFlag(options, AllChars)
	case "--help":
		options = addFlag(options, Help)
	case "--version":
		options = addFlag(options, Version)
	}
	return options
}

func parseShortOption(opt string, options Flag) Flag {
	for _, flag := range []rune(opt) {
		switch string(flag) {
		case "u":
			options = setFlag(options, Url, Slug, Xml)
		case "x":
			options = setFlag(options, Xml, Url, Slug)
		case "s":
			options = setFlag(options, Slug, Xml, Url)
		case "e":
			options = setFlag(options, Encode, Decode)
		case "d":
			options = setFlag(options, Decode, Encode)
		case "a":
			options = addFlag(options, AllChars)
		case "h":
			options = addFlag(options, Help)
		case "v":
			options = addFlag(options, Version)
		}
	}
	return options
}

func setFlag(options Flag, which Flag, clears ...Flag) Flag {
	for _, c := range clears {
		options = clearFlag(options, c)
	}
	options = addFlag(options, which)
	return options
}

func (x Options) HasFlag(which Flag) bool {
	return x.options.hasFlag(which)
}

func (x Options) Subject() string {
	return strings.Join(x.rest, " ")
}
