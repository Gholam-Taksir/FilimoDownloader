package cli

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	appName string
	Id      string
	Token   string
	Build   string
	Output  string
	Author  bool
	Version bool
	Help    bool
}

func (args Args) PrintHelp() {
	fmt.Println("Usage:", args.appName, "[OPTIONS]")
	fmt.Println()
	for _, f := range GetFlags(nil, nil) {
		fmt.Printf("  -%s, --%s\n      %s\n\n", f.ShortName, f.FullName, f.Usage)
	}
}

func NewArgs(appName string) Args {
	args := Args{appName: appName}
	flagSet := flag.NewFlagSet(appName, flag.ExitOnError)
	for _, f := range GetFlags(flagSet, &args) {
		f.Register(&f)
	}
	flagSet.Usage = func() {
		args.PrintHelp()
	}
	flagSet.Parse(os.Args[1:])
	return args
}
