package git

import (
	"testing"
)

func TestHasUncommittedChanges(t *testing.T) {
	repoPath := "." // assuming the current repo for testing

	// Since we're testing on a real repo, the results might vary
	_, err := HasUncommittedChanges(repoPath)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Here, you might want to add more checks depending on your setup or expectations
}

func TestHasUnpushedCommits(t *testing.T) {
	repoPath := "." // again, assuming the current repo for testing

	_, err := HasUnpushedCommits(repoPath)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Here, you can add more checks too
}
