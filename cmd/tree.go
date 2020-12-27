package cmd

import (
	"github.com/urfave/cli/v2"
)

const numberOfLines = "number_of_lines"

//CmdTree is test command.
var CmdTree = cli.Command{
	Name:    "tree",
	Aliases: []string{"tr"},
	Usage:   "Print out a christmas tree on console.",
	Flags: []cli.Flag{
		&cli.IntFlag{Name: numberOfLines, Aliases: []string{"nof"}, Value: 5},
	},
	Description: "Print a christmas tree on console with the given high (numberOfLines).",
	Action:      runEmpty,
}

func runEmpty(ctx *cli.Context) error {
	return nil
}
