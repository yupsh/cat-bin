package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/cat"
)

const (
	flagNumberLines  = "number"
	flagShowEnds     = "show-ends"
	flagShowTabs     = "show-tabs"
	flagSqueezeBlank = "squeeze-blank"
	flagTrimSpaces   = "trim-spaces"
)

func main() {
	app := &cli.App{
		Name:  "cat",
		Usage: "concatenate files and print on the standard output",
		UsageText: `cat [OPTIONS] [FILE...]

   Concatenate FILE(s) to standard output.
   With no FILE, or when FILE is -, read standard input.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagNumberLines,
				Aliases: []string{"n"},
				Usage:   "number all output lines",
			},
			&cli.BoolFlag{
				Name:    flagShowEnds,
				Aliases: []string{"E"},
				Usage:   "display $ at end of each line",
			},
			&cli.BoolFlag{
				Name:    flagShowTabs,
				Aliases: []string{"T"},
				Usage:   "display TAB characters as ^I",
			},
			&cli.BoolFlag{
				Name:    flagSqueezeBlank,
				Aliases: []string{"s"},
				Usage:   "suppress repeated empty output lines",
			},
			&cli.BoolFlag{
				Name:  flagTrimSpaces,
				Usage: "trim trailing spaces from lines",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "cat: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (or none for stdin)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, yup.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.Bool(flagNumberLines) {
		params = append(params, NumberLines)
	}
	if c.Bool(flagShowEnds) {
		params = append(params, ShowEnds)
	}
	if c.Bool(flagShowTabs) {
		params = append(params, ShowTabs)
	}
	if c.Bool(flagSqueezeBlank) {
		params = append(params, SqueezeBlank)
	}
	if c.Bool(flagTrimSpaces) {
		params = append(params, TrimSpaces)
	}

	// Create and execute the cat command
	cmd := Cat(params...)
	return yup.Run(cmd)
}
