package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eludda/geekyland/unfold/app"
)

var config app.Config

func init() {
	config = app.Config{}
	config.Register(flag.CommandLine)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: unfold <folder>\n\n")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if flag.NArg() > 0 {
		config.SetRoot(flag.Arg(0))

		if !config.Root.IsDir() {
			fmt.Fprintf(os.Stderr, "%s is not a directory.\n", config.Root.Value)
			os.Exit(1)
		}
	}

	fold := app.Fold{Out: config.OutDir}
	fold.Unfold(config.Root)
}
