package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/mtstnt/pio/utils"
	"github.com/urfave/cli/v2"
)

var flags = []cli.Flag{
	&cli.StringSliceFlag{
		Name:    "exclude-dirs",
		Aliases: []string{"e"},
		Usage:   "Exclude directories to add to the templates list",
	},
	&cli.StringSliceFlag{
		Name:    "exclude-files",
		Aliases: []string{"x"},
		Usage:   "Exclude files to add to the templates list",
	},
	&cli.StringSliceFlag{
		Name:    "include-dirs",
		Aliases: []string{"i"},
		Usage:   "Include directories to add to the templates list",
	},
	&cli.StringSliceFlag{
		Name:    "include-files",
		Aliases: []string{"n"},
		Usage:   "Include files to add to the templates list",
	},
}

func Add() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Usage:     "Add a project to your project templates list",
		ArgsUsage: "The name of the template",
		Action:    addActionFunc,
		Flags:     flags,
	}
}

func addActionFunc(ctx *cli.Context) error {
	args := ctx.Args()

	templateName := ""
	fromPath := "."

	// Get template name
	if args.Len() >= 1 {
		templateName = args.Get(0)
	} else {
		return fmt.Errorf("must supply template name")
	}

	// Get path
	if args.Len() >= 2 {
		fromPath = args.Get(1)
	}

	// Check if duplicate
	tmpls, err := utils.GetTemplatesList()
	if err != nil {
		return err
	}

	for _, t := range tmpls {
		if t.Name == templateName {
			return fmt.Errorf("template name %s already exists", templateName)
		}
	}

	createTemplate(
		templateName,
		fromPath,
	)

	return nil
}

func createTemplate(tmplName, fromPath string) error {

	// Create directory
	err := os.MkdirAll(path.Join(utils.APP_PATH, "templates", tmplName), os.ModeDir)

	if err != nil {
		return err
	}

	// Copy from `fromPath`

	return nil
}
