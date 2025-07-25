package main

import (
	"io"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	// Open source
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	// Create destination dir
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	// Create destination
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy data
	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	// Sync to disk
	if err := out.Sync(); err != nil {
		return err
	}

	// Set file to read-only mode (0444)
	return os.Chmod(dst, 0444)
}

func main() {
	// Get home dir
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// src is the source file that we want to copy
	src := "update.exe"

	// Check if the source file exists
	if _, err := os.Stat(src); os.IsNotExist(err) {
		panic("source file does not exist")
	}

	// dst1 is the destination path that we want to copy to
	dst1 := filepath.Join(home, "AppData", "Local", "CapCut", "Apps", "update.exe")

	// Copy the file to the first location
	if err := copyFile(src, dst1); err != nil {
		panic(err)
	}

	// dst2 is the destination path that we want to copy to
	dst2 := filepath.Join(home, "AppData", "Local", "CapCut", "User Data", "Download", "update.exe")

	// Copy the file to the second location
	if err := copyFile(src, dst2); err != nil {
		panic(err)
	}
}