package main

import (
	"fmt"
	"komunalno/cdc/cmd"
	"komunalno/cdc/pkg/codec"
	"komunalno/cdc/pkg/opts"
	"os"
)

func main() {
	options := opts.New(os.Args[1:])
	if options.HasFlag(opts.Help) {
		fmt.Println(opts.GetHelp())
		os.Exit(0)
	}

	if options.HasFlag(opts.Version) {
		fmt.Println(opts.GetVersion())
		os.Exit(0)
	}

	subject, ok := cmd.GetStdin()
	if !ok {
		subject = options.Subject()
	}

	var cdc codec.Codec
	if options.HasFlag(opts.Xml) {
		cdc = codec.Html(subject)
	}
	if options.HasFlag(opts.Url) {
		cdc = codec.Url(subject)
	}
	if options.HasFlag(opts.Slug) {
		cdc = codec.Slug(subject)
	}

	var result string
	if options.HasFlag(opts.Encode) {
		if options.HasFlag(opts.AllChars) {
			result = cdc.EncodeAll()
		} else {
			result = cdc.Encode()
		}
	}
	if options.HasFlag(opts.Decode) {
		result = cdc.Decode()
	}
	fmt.Println(result)
}
