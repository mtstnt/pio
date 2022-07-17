package main

import (
	"fmt"
	"log"
	"os"

	cmd "github.com/mtstnt/pio/commands"
	"github.com/mtstnt/pio/utils"
	"github.com/urfave/cli/v2"
)

func registerCommands() []*cli.Command {
	return []*cli.Command{
		cmd.Add(),
		cmd.Generate(),
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	utils.SetupConstants()

	// Create the templates folder
	err := os.MkdirAll(utils.TEMPLATES_PATH, os.ModeDir)
	if err != nil {
		log.Fatalln(err)
	}

	app := &cli.App{
		Name:  "pio",
		Usage: "Enter -h or --help to show help",
		Action: func(ctx *cli.Context) error {
			if ctx.Bool("version") {
				fmt.Println("Pio version 0.0.1a")
				fmt.Println("A simplistic project and template generator by mtstnt (https://github.com/mtstnt)")
				return nil
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "version",
				Value: true,
			},
		},
		Commands: registerCommands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
