package main

import (
	"os"
	"path/filepath"
)

func createUpdateExe(destinationPath string) error {
	// Ensure the destination directory exists
	if err := os.MkdirAll(filepath.Dir(destinationPath), 0755); err != nil {
		return err
	}

	// Create the update.exe file at the destination
	file, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the desired content to the file
	content := "This would be the actual executable content"
	if _, err := file.WriteString(content); err != nil {
		return err
	}

	// Flush written content to disk
	if err := file.Sync(); err != nil {
		return err
	}

	// Set file permissions to read-only
	return os.Chmod(destinationPath, 0444)
}

func main() {
	// Get home directory
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Destination paths
	dst1 := filepath.Join(home, "AppData", "Local", "CapCut", "Apps", "update.exe")
	dst2 := filepath.Join(home, "AppData", "Local", "CapCut", "User Data", "Download", "update.exe")

	// Create the first update.exe file
	if err := createUpdateExe(dst1); err != nil {
		panic(err)
	}

	// Create the second update.exe file
	if err := createUpdateExe(dst2); err != nil {
		panic(err)
	}
}