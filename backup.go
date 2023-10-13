package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// backupFiles copies files from the source directory to the destination directory.
func backupFiles(sourcePath, destPath string) error {

	// Read the contents of the source directory
	sourceInfo, err := os.Stat(sourcePath)
	if err != nil {
		return err
	}

	if !sourceInfo.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
		return err
	}

	// Read the source directory contents
	sourceFiles, err := os.ReadDir(sourcePath)
	if err != nil {
		return err
	}

	// Copy files from the source to the destination
	for _, fileInfo := range sourceFiles {
		sourceFile := filepath.Join(sourcePath, fileInfo.Name())
		destFile := filepath.Join(destPath, fileInfo.Name())

		if fileInfo.IsDir() {
			// Recursively copy subdirectories
			if err := backupFiles(sourceFile, destFile); err != nil {
				return err
			}
		} else {
			// If it's a file, copy it to the destination directory
			err := copyFile(sourceFile, destFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// copyFile copies a file from source to destination.
func copyFile(sourceFile, destFile string) error {
	source, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}
