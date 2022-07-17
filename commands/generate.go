package commands

import (
	"fmt"
	"path"

	"github.com/mtstnt/pio/utils"
	"github.com/urfave/cli/v2"
)

var generateFlags = []cli.Flag{
	&cli.StringSliceFlag{
		Name:    "exclude-dirs",
		Aliases: []string{"e"},
		Usage:   "Exclude directories in generated template",
	},
	&cli.StringSliceFlag{
		Name:    "exclude-files",
		Aliases: []string{"x"},
		Usage:   "Exclude files in generated template",
	},
	&cli.StringSliceFlag{
		Name:    "include-dirs",
		Aliases: []string{"i"},
		Usage:   "Include directories in generated template",
	},
	&cli.StringSliceFlag{
		Name:    "include-files",
		Aliases: []string{"n"},
		Usage:   "Include files in generated template",
	},
}

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
	var (
		tmplName     string = ""
		generatePath string = "."

		args = ctx.Args()
	)

	if args.Len() >= 1 {
		tmplName = args.Get(0)
	} else {
		return fmt.Errorf("no template name provided")
	}

	if args.Len() >= 2 {
		generatePath = args.Get(1)
	}

	// Check if template name exists in list
	tmplList, err := utils.GetTemplatesList()
	if err != nil {
		return err
	}

	var found bool = false
	for _, t := range tmplList {
		if t.Name == tmplName {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("template " + tmplName + " does not exist")
	}

	tmplPath := path.Join(utils.TEMPLATES_PATH, tmplName)
	targetPath := path.Join(generatePath)

	if err := utils.Copy(
		tmplPath,
		targetPath,
		utils.Nothing,
		utils.Nothing,

		utils.LookupMap{
			"pio.yml": true,
		},
		utils.Nothing,
	); err != nil {
		return err
	}

	return nil
}
