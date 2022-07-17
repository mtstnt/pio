package utils

import (
	"io"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

// Format of the pio.yml
type TemplateInfo struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

func GetTemplatesList() ([]TemplateInfo, error) {

	// Read the templates directory and format each of them into templates list
	entries, err := os.ReadDir(TEMPLATES_PATH)
	if err != nil {
		return nil, err
	}

	templates := []TemplateInfo{}

	for _, entry := range entries {
		if entry.IsDir() {
			tmplPath := path.Join(TEMPLATES_PATH, entry.Name())

			// Read the templates' internals in slurp.yml
			tmplInfo, err := readPioConfig(path.Join(tmplPath, "pio.yml"))
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
