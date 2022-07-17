package utils

import (
	"fmt"
	"io"
	"os"
	"path"
)

type LookupMap map[string]bool

// Ngapain ngetik lookupmap kosong kayak orang gilaaA?
var Nothing = LookupMap{}

// Copies all files from fromPath to destPath, with exceptions and inclusions in CopyInfo
func Copy(fromPath, destPath string, excludeDirs, includeDirs, excludeFiles, includeFiles LookupMap) error {
	if err := os.MkdirAll(destPath, os.ModeDir); err != nil {
		return err
	}

	if err := copyRecursive(fromPath, destPath, excludeDirs, includeDirs, excludeFiles, includeFiles); err != nil {
		return err
	}

	return nil
}

// TODO: Include, Exclude Files not yet implemented.
func copyRecursive(fromPath, destPath string, excludeDirs, includeDirs, excludeFiles, includeFiles LookupMap) error {
	entries, err := os.ReadDir(fromPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		completeFromPath := path.Join(fromPath, entry.Name())
		completeDestPath := path.Join(destPath, entry.Name())

		// TODO: The excludeDir is still problematic for generate as it has the complete full path
		if (entry.IsDir() && isInList(completeFromPath, excludeDirs)) ||
			(!entry.IsDir() && isInList(entry.Name(), excludeFiles)) ||
			entry.Name() == "." || entry.Name() == ".." {
			continue
		}

		if entry.IsDir() {
			// Must not pass ptr because we append to its from and dest path
			cFromPath := completeFromPath
			cDestPath := completeDestPath

			if err := copyRecursive(cFromPath, cDestPath, excludeDirs, includeDirs, excludeFiles, includeFiles); err != nil {
				return err
			}

		} else {

			fptrSource, err := os.Open(completeFromPath)
			if err != nil {
				return err
			}

			err = os.MkdirAll(destPath, os.ModeDir)
			if err != nil {
				return err
			}

			fptrDest, err := os.Create(completeDestPath)
			if err != nil {
				return err
			}

			_, err = io.Copy(fptrDest, fptrSource)
			if err != nil {
				return err
			}

			fmt.Println("COPIED " + completeFromPath + " TO " + completeDestPath)
		}
	}

	return nil
}

func isInList(name string, list LookupMap) bool {
	_, ok := list[name]
	return ok
}
