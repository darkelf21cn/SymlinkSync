package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func createSymbolicLink(sourcePath, destPath string) error {
	// Check if the destination file already exists
	_, err := os.Lstat(destPath)
	if err == nil {
		// Remove the existing file
		fmt.Printf("Removing existing file %s\n", destPath)
		if err := os.Remove(destPath); err != nil {
			return err
		}
	} else if !os.IsNotExist(err) {
		return err
	}

	// Ensure the directory structure exists
	destDir := filepath.Dir(destPath)
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return err
	}

	// Create a symbolic link from destination to source
	return os.Symlink(sourcePath, destPath)
}

func visitFile(path string, info os.FileInfo, sourceDir, destDir string) error {
	if !info.Mode().IsRegular() {
		return nil // Skip non-regular files
	}

	relPath, err := filepath.Rel(sourceDir, path)
	if err != nil {
		return err
	}

	destPath := filepath.Join(destDir, relPath)

	fmt.Printf("Creating symbolic link for %s\n", path)
	return createSymbolicLink(path, destPath)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: symlink-tool <source_dir> <dest_dir>")
		return
	}

	sourceDir := os.Args[1]
	destDir := os.Args[2]

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return visitFile(path, info, sourceDir, destDir)
	})
	if err != nil {
		log.Fatal(err)
	}
}
