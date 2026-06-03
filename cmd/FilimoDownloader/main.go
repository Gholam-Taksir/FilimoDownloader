package main

import (
	"os"

	"FilimoDownloader-GholamTaksir/internal/cli"
)

//go:generate rsrc -manifest filimo.manifest -ico icon.ico -o rsrc.syso

var isProduction string

func main() {
	app := cli.NewApp(isProduction == "true")
	args := cli.NewArgs(app.Name)

	if len(os.Args) > 1 && os.Args[1] == "--history" {
		app.History.Print()
		return
	}

	if args.Build != "" {
		cli.Build(args.Build, args.Output)
		return
	}

	cli.Download(app, args)
}
