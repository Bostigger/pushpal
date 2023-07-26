package main

import (
	"os"
	"strings"
	"testing"
)

func TestAddPathToGitignore(t *testing.T) {
	repoPath := "." // assuming the current repo for testing
	pathToIgnore := "test_ignore_path.txt"

	// Ensure cleanup after the test
	defer func() {
		os.Remove("./.gitignore")
	}()

	// Test the function
	err := addPathToGitignore(repoPath, pathToIgnore)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Now, let's check if the .gitignore file contains our path
	contents, err := os.ReadFile("./.gitignore")
	if err != nil {
		t.Fatalf("Error reading .gitignore: %v", err)
	}

	if !strings.Contains(string(contents), pathToIgnore) {
		t.Fatalf("Expected .gitignore to contain %s, but it didn't", pathToIgnore)
	}
}
