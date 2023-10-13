package main

import (
	"fmt"
	"os"
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
