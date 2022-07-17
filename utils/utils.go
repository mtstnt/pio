package utils

import (
	"io"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

const (
	TEMPLATES_DIR_NAME = "templates"
)

// Format of the pio.yml
type TemplateInfo struct {
	Namespace string
	Name      string
	Path      string
}

func GetTemplatesList() ([]TemplateInfo, error) {

	// Read the templates directory and format each of them into templates list
	baseTmplPath := path.Join(APP_PATH, TEMPLATES_DIR_NAME)
	entries, err := os.ReadDir(baseTmplPath)
	if err != nil {
		return nil, err
	}

	templates := []TemplateInfo{}

	for _, entry := range entries {
		if entry.IsDir() {
			tmplPath := path.Join(baseTmplPath, entry.Name())

			// Read the templates' internals in slurp.yml
			tmplInfo, err := readPioConfig(path.Join(tmplPath, "slurp.yml"))
			if err != nil {
				return nil, err
			}

			templates = append(templates, tmplInfo)
		}
	}

	return templates, nil
}

func readPioConfig(pioConfigPath string) (TemplateInfo, error) {
	fptr, err := os.Open(pioConfigPath)
	if err != nil {
		return TemplateInfo{}, err
	}

	pioConfigContentsByte, err := io.ReadAll(fptr)
	if err != nil {
		return TemplateInfo{}, err
	}

	var tmplInfo TemplateInfo

	err = yaml.Unmarshal(pioConfigContentsByte, &tmplInfo)
	if err != nil {
		return TemplateInfo{}, err
	}

	return TemplateInfo{}, nil
}
