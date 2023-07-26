package git

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func HasUncommittedChanges(repoPath string) ([]string, error) {
	var changedFiles []string

	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = repoPath

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			fields := strings.Fields(line) // split by whitespace
			if len(fields) > 1 {
				changedFiles = append(changedFiles, fields[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return changedFiles, nil
}

func HasUnpushedCommits(repoPath string) ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "@{u}..")
	cmd.Dir = repoPath

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("git error: %s", out.String())
	}

	output := strings.TrimSpace(out.String())

	if output == "" {
		return []string{}, nil
	}

	return strings.Split(output, "\n"), nil
}
