package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
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
	fmt.Printf("Creating symbolic link for %s\n", sourcePath)
	return os.Symlink(sourcePath, destPath)
}

func visitFile(path string, info os.FileInfo, sourceDir, destDir string) error {
	if info.Mode().IsDir() {
		return nil // Skil on directories
	}

	relPath, err := filepath.Rel(sourceDir, path)
	if err != nil {
		return err
	}

	destPath := filepath.Join(destDir, relPath)

	return createSymbolicLink(path, destPath)
}

func main() {
	var sourceDir, destDir, excludePattern string

	var cmd = &cobra.Command{
		Use:   "symlinksync",
		Short: "Synchronize symbolic links from source to destination directory",
		Run: func(cmd *cobra.Command, args []string) {
			if sourceDir == "" || destDir == "" {
				fmt.Println("Both source and destination directories are required.")
				os.Exit(1)
			}

			err := symlinkSync(sourceDir, destDir, excludePattern)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
		}}

	cmd.Flags().StringVarP(&sourceDir, "source", "s", "", "Source directory")
	cmd.Flags().StringVarP(&destDir, "destination", "d", "", "Destination directory")
	cmd.Flags().StringVar(&excludePattern, "exclude", "", "Regular expression pattern to exclude files or directories in the source path")
	cmd.Execute()
}

func symlinkSync(sourceDir, destDir, excludePattern string) error {
	excludeRegex, err := regexp.Compile(excludePattern)
	if err != nil {
		return err
	}

	fmt.Printf("Synchronizing symbolic links from %s to %s\n", sourceDir, destDir)

	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if excludeRegex.MatchString(info.Name()) {
			if info.IsDir() {
				// Skip the entire directory by returning filepath.SkipDir
				fmt.Printf("Directory %s is excluded\n", path)
				return filepath.SkipDir
			} else {
				fmt.Printf("File %s is excluded\n", path)
				return nil
			}
		}
		return visitFile(path, info, sourceDir, destDir)
	})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
