package commands

import "github.com/urfave/cli/v2"

var generateFlags = []cli.Flag{}

func Generate() *cli.Command {
	return &cli.Command{
		Name:      "generate",
		Usage:     "Generate from template",
		ArgsUsage: "The name of the template",
		Action:    generateActionFunc,
		Flags:     generateFlags,
	}
}

func generateActionFunc(ctx *cli.Context) error {
	return nil
}
