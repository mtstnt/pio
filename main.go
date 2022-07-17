package main

import (
	"log"
	"os"
	"path"

	cmd "github.com/mtstnt/pio/commands"
	"github.com/mtstnt/pio/utils"
	"github.com/urfave/cli/v2"
)

func registerCommands() []*cli.Command {
	return []*cli.Command{
		cmd.Add(),
	}
}

func main() {
	utils.SetupConstants()

	app := &cli.App{
		Name:  "Pio",
		Usage: "Enter -h or --help to show help",
		Action: func(*cli.Context) error {
			return utils.Copy(
				".",
				path.Join(utils.APP_PATH, "templates", "testingwoe"),
				utils.LookupMap{
					".git":            true,
					"build/templates": true,
				},
				utils.Nothing,
				utils.Nothing,
				utils.Nothing,
			)
		},
		Commands: []*cli.Command{},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
