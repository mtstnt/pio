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
		Name:  "Pio",
		Usage: "Enter -h or --help to show help",
		Action: func(*cli.Context) error {
			fmt.Println("Mantap")
			return nil
		},
		Commands: registerCommands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
