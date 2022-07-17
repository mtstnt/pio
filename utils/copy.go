package utils

import (
	"fmt"
	"os"
	"path"
)

type LookupMap map[string]bool

// Ngapain ngetik lookupmap kosong kayak orang gilaaA?
var Nothing = LookupMap{}

// Copies all files from fromPath to destPath, with exceptions and inclusions in CopyInfo
func Copy(fromPath, destPath string, excludeDirs, includeDirs, excludeFiles, includeFiles LookupMap) error {
	return copyRecursive(fromPath, destPath, excludeDirs, includeDirs, excludeFiles, includeFiles)
}

func copyRecursive(fromPath, destPath string, excludeDirs, includeDirs, excludeFiles, includeFiles LookupMap) error {
	entries, err := os.ReadDir(fromPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		completePath := path.Join(fromPath, entry.Name())

		if isInList(completePath, excludeDirs) ||
			isInList(completePath, excludeFiles) ||
			entry.Name() == "." ||
			entry.Name() == ".." {
			continue
		}

		if entry.IsDir() {
			// Must not pass ptr because we append to its from and dest path
			cFromPath := completePath
			cDestPath := path.Join(destPath, entry.Name())

			if err := copyRecursive(cFromPath, cDestPath, excludeDirs, includeDirs, excludeFiles, includeFiles); err != nil {
				return err
			}
		} else {
			// TODO: Actually copy the files lmao
			fmt.Println("COPIED " + path.Join(fromPath, entry.Name()))
		}
	}

	return nil
}

func isInList(name string, list LookupMap) bool {
	_, ok := list[name]
	return ok
}
