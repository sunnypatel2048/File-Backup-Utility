package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("---- This is a File Backup Utility -----")

	// Get source and destination paths from user input
	sourcePath := getInput("Enter the source path: ")
	destPath := getInput("Enter the destination path: ")

	// Create the destination directory if it doesn't exist
	// "os.ModePerm" ensures that the newly created directory will have full read, write, and execute permissions for the owner, group, and others.
	if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
		fmt.Printf("Error creating destination directory: %v\n", err)
		return
	}

	// Perform Backup
	if err := backupFiles(sourcePath, destPath); err != nil {
		fmt.Printf("Error occured during backup: %v\n", err)
		return
	}

	fmt.Println("\nBackup succeeded!")
	fmt.Scanln()
}

// getInput prompts the user for input and returns the entered text.
func getInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

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
