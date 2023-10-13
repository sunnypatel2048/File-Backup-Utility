package main

import (
	"os"
	"testing"
)

func Test_BackupFiles(t *testing.T) {
	// Create a temporary source directory
	sourceDir := "test_source"
	err := os.Mkdir(sourceDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create the source directory: %v", err)
	}
	defer os.Remove(sourceDir)

	// Create a temporary destination directory
	destDir := "test_dest"
	err = os.Mkdir(destDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create the destination directory: %v", err)
	}
	defer os.Remove(destDir)

	// Create a temporary file in the source directory
	sourceFile := "test_source/testfile.txt"
	_, err = os.Create(sourceFile)
	if err != nil {
		t.Fatalf("Failed to create a test file: %v", err)
	}

	// Run the backupFiles function
	err = backupFiles(sourceDir, destDir)
	if err != nil {
		t.Errorf("BackupFiles returned an error: %v", err)
	}

	// Check if the file was copied to the destination directory
	destFile := "test_dest/testfile.txt"
	_, err = os.Stat(destFile)
	if err != nil {
		t.Errorf("File was not copied to the destination directory: %v", err)
	}
}

func Test_BackupFiles_NonExistentSource(t *testing.T) {
	nonExistentSourceDir := "non_existent_source"
	err := backupFiles(nonExistentSourceDir, "destination")
	if err == nil {
		t.Error("Expected an error when backing up a non-existent source directory.")
	}
}

func Test_BackupFiles_FileAsSource(t *testing.T) {
	sourceFile := "test_source/testfile.txt"
	err := backupFiles(sourceFile, "destination")
	if err == nil {
		t.Error("Expected an error when backing up a file as the source.")
	}
}

func Test_BackupFiles_UnwritableDestination(t *testing.T) {
	unwritableDestDir := "unwritable_dest"
	err := os.Mkdir(unwritableDestDir, 0) // Create an unwritable directory
	if err != nil {
		t.Fatalf("Failed to create unwritable destination directory: %v", err)
	}
	defer os.RemoveAll(unwritableDestDir)
	err = backupFiles("source", unwritableDestDir)
	if err == nil {
		t.Error("Expected an error when the destination directory is not writable.")
	}
}
