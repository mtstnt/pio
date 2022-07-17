package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/mtstnt/pio/utils"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var addFlags = []cli.Flag{
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
		Flags:     addFlags,
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

	// Create and copy the template
	if err := createTemplate(
		templateName,
		fromPath,
	); err != nil {
		return err
	}

	return nil
}

func createTemplate(tmplName, fromPath string) error {

	// Create directory
	destPath := path.Join(utils.TEMPLATES_PATH, tmplName)
	err := os.MkdirAll(destPath, os.ModeDir)

	if err != nil {
		return err
	}

	fmt.Println("ADDING TEMPLATE " + tmplName + " TO PATH \"" + destPath + "\"")

	// Copy from `fromPath`
	if err = utils.Copy(
		fromPath,
		destPath,
		utils.LookupMap{
			".git":         true,
			"node_modules": true,
			"vendor":       true,
		},
		utils.Nothing,
		utils.Nothing,
		utils.Nothing,
	); err != nil {
		return err
	}

	// Setup pio file
	if err := setupPioConfig(destPath, tmplName); err != nil {
		return err
	}

	return nil
}

func setupPioConfig(destPath string, tmplName string) error {
	pioFilePath := path.Join(destPath, "pio.yml")

	tmplInfo := utils.TemplateInfo{
		Name: tmplName,
		Path: destPath,
	}

	yamlByte, err := yaml.Marshal(&tmplInfo)
	if err != nil {
		return err
	}

	if err = os.WriteFile(pioFilePath, yamlByte, 0777); err != nil {
		return err
	}

	return nil
}
