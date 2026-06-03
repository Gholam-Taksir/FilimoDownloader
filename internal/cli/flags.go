package cli

import "flag"

type Flag struct {
	ShortName string
	FullName  string
	Usage     string
	Register  func(flag *Flag)
}

func GetFlags(flagSet *flag.FlagSet, args *Args) []Flag {
	return []Flag{
		{
			ShortName: "i",
			FullName:  "id",
			Usage:     "<string> Movie or Series ID (or full URL)",
			Register: func(f *Flag) {
				flagSet.StringVar(&args.Id, f.ShortName, args.Id, f.Usage)
				flagSet.StringVar(&args.Id, f.FullName, args.Id, f.Usage)
			},
		},
		{
			ShortName: "t",
			FullName:  "token",
			Usage:     "<string> Auth token value",
			Register: func(f *Flag) {
				flagSet.StringVar(&args.Token, f.ShortName, args.Token, f.Usage)
				flagSet.StringVar(&args.Token, f.FullName, args.Token, f.Usage)
			},
		},
		{
			ShortName: "b",
			FullName:  "build",
			Usage:     "<string> Build a previously downloaded directory",
			Register: func(f *Flag) {
				flagSet.StringVar(&args.Build, f.ShortName, args.Build, f.Usage)
				flagSet.StringVar(&args.Build, f.FullName, args.Build, f.Usage)
			},
		},
		{
			ShortName: "o",
			FullName:  "output",
			Usage:     "<string> Output directory or filename",
			Register: func(f *Flag) {
				flagSet.StringVar(&args.Output, f.ShortName, args.Output, f.Usage)
				flagSet.StringVar(&args.Output, f.FullName, args.Output, f.Usage)
			},
		},
		{
			ShortName: "a",
			FullName:  "author",
			Usage:     "Print author name",
			Register: func(f *Flag) {
				flagSet.BoolVar(&args.Author, f.ShortName, args.Author, f.Usage)
				flagSet.BoolVar(&args.Author, f.FullName, args.Author, f.Usage)
			},
		},
		{
			ShortName: "v",
			FullName:  "version",
			Usage:     "Print app version",
			Register: func(f *Flag) {
				flagSet.BoolVar(&args.Version, f.ShortName, args.Version, f.Usage)
				flagSet.BoolVar(&args.Version, f.FullName, args.Version, f.Usage)
			},
		},
		{
			ShortName: "h",
			FullName:  "help",
			Usage:     "Show this help message",
			Register: func(f *Flag) {
				flagSet.BoolVar(&args.Help, f.ShortName, args.Help, f.Usage)
				flagSet.BoolVar(&args.Help, f.FullName, args.Help, f.Usage)
			},
		},
	}
}
