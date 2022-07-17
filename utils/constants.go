package utils

import (
	"log"
	"os"
)

const (
	TEMPLATES_PATH = "./templates"
)

var (
	APP_PATH  string
	EXEC_PATH string
)

func init() {
	// Setup the app executable path
	appPath, err := os.Executable()
	if err != nil {
		log.Fatalln("Failed to get application path")
	}

	APP_PATH = appPath

	// Setup the path original exec path
	execPath, err := os.Getwd()
	if err != nil {
		log.Fatalln("Failed to get current executed path")
	}

	EXEC_PATH = execPath
}
