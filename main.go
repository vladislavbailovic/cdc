package main

import (
	"fmt"
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
}
