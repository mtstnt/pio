package utils

import (
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

const (
	TEMPLATES_DIR_NAME = "templates"
)

var (
	APP_PATH       string
	EXEC_PATH      string
	TEMPLATES_PATH string
)

func SetupConstants() {
	// Setup the app executable path
	appPath, err := os.Executable()
	if err != nil {
		log.Fatalln("Failed to get application path")
	}

	// cd ..
	a := strings.Split(appPath, string(os.PathSeparator))
	b := strings.Join(a[:len(a)-1], string(os.PathSeparator))

	APP_PATH = b

	// Setup the path original exec path
	execPath, err := os.Getwd()
	if err != nil {
		log.Fatalln("Failed to get current executed path")
	}

	EXEC_PATH = execPath

	if runtime.GOOS == "windows" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln("Failed to get home directory (Windows)")
		}

		TEMPLATES_PATH = path.Join(homeDir, ".pio", "templates")
	} else {
		TEMPLATES_PATH = path.Join("/", "etc", "pio", "templates")
	}
}
