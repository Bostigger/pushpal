package main

import (
	"fmt"
	"github.com/bostigger/pushpal/pkg/git"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func addPathToGitignore(repoPath, pathToIgnore string) error {
	gitignorePath := filepath.Join(repoPath, ".gitignore")

	// Check if the repoPath directory exists
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		return fmt.Errorf("repository directory %s does not exist", repoPath)
	}

	// Check if .gitignore already contains the path
	contents, err := os.ReadFile(gitignorePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error reading .gitignore: %v", err)
	}

	if strings.Contains(string(contents), pathToIgnore) {
		// Already ignored
		return nil
	}

	// Append the path to the .gitignore
	f, err := os.OpenFile(gitignorePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening .gitignore for appending: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(pathToIgnore + "\n")
	if err != nil {
		return fmt.Errorf("failed to write to .gitignore: %v", err)
	}

	return nil
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.Lshortfile)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	changedFiles, err := git.HasUncommittedChanges(cwd)
	if err != nil {
		log.Fatalf("Error checking uncommitted changes for %v: %v", cwd, err)
	}

	unpushedFiles, err := git.HasUnpushedCommits(cwd)
	if err != nil {
		log.Fatalf("Error checking unpushed commits for %v: %v", cwd, err)
	}

	ignorablePathsSet := map[string]struct{}{
		".DS_Store":     {},
		"Thumbs.db":     {},
		"node_modules/": {},
		"__pycache__/":  {},
		"*.pyc":         {},
		"venv/":         {},
		"target/":       {},
		"bin/":          {},
		"obj/":          {},
		"*.log":         {},
		".idea/":        {},
		".vscode/":      {},
		".env":          {},
		"dist/":         {},
		"build/":        {},
	}

	var pathsToOffer []string
	for _, file := range changedFiles {
		for ignorable := range ignorablePathsSet {
			if strings.HasPrefix(file, ignorable) {
				pathsToOffer = append(pathsToOffer, ignorable)
				delete(ignorablePathsSet, ignorable) // Prevents adding the same path multiple times
			}
		}
	}

	handlePathsOffer(cwd, pathsToOffer, changedFiles)
	handleUnpushedFiles(cwd, unpushedFiles)
}

func handlePathsOffer(cwd string, pathsToOffer, changedFiles []string) {
	if len(pathsToOffer) > 0 {
		fmt.Printf("Detected uncommitted changes in following directories: %s\n", strings.Join(pathsToOffer, ", "))
		fmt.Println("Do you want to add these directories to your .gitignore? (yes/no)")
		var input string
		fmt.Scanln(&input)
		if strings.ToLower(input) == "yes" {
			for _, path := range pathsToOffer {
				if err := addPathToGitignore(cwd, path); err != nil {
					log.Fatalf("Error adding %s to .gitignore: %v", path, err)
				} else {
					fmt.Printf("%s added to .gitignore.\n", path)
				}
			}
		}
	} else if len(changedFiles) > 0 {
		fmt.Printf("%v has uncommitted changes in the following files: %s\n", cwd, strings.Join(changedFiles, ", "))
	}
}

func handleUnpushedFiles(cwd string, unpushedFiles []string) {
	if len(unpushedFiles) > 0 && unpushedFiles[0] != "" {
		fmt.Printf("%v has unpushed changes in the following files: %s\n", cwd, strings.Join(unpushedFiles, ", "))
	}
}
