package gitutil

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// RunGitCommand executes a git command in a specified directory.
func RunGitCommand(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir

	// Pass through stdout and stderr for visibility
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git command failed (%v) in %s: %w", args, dir, err)
	}
	return nil
}

// RunGitCommandOutput executes a git command and returns stdout as string (stderr included in errors).
func RunGitCommandOutput(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git command failed (%v) in %s: %s: %w", args, dir, strings.TrimSpace(stderr.String()), err)
	}
	return stdout.String(), nil
}

// Init initializes a new Git repository in the given directory.
func Init(dir string) error {
	return RunGitCommand(dir, "init")
}

// Add stages all changes in the given directory.
func Add(dir string) error {
	return RunGitCommand(dir, "add", ".")
}

// AddPaths stages specific paths in the given directory.
func AddPaths(dir string, paths ...string) error {
	if len(paths) == 0 {
		return nil
	}
	args := append([]string{"add", "--"}, paths...)
	return RunGitCommand(dir, args...)
}

// Commit creates a commit with the given message in the directory.
func Commit(dir, message string) error {
	return RunGitCommand(dir, "commit", "-m", message)
}

// AddSubmodule adds a submodule to the parent repository.
func AddSubmodule(parentDir, submodulePath string) error {
	// Submodule paths are relative to the parent directory
	return RunGitCommand(parentDir, "submodule", "add", "./"+submodulePath, submodulePath)
}

// SubmoduleDeinit deinitializes a submodule.
func SubmoduleDeinit(parentDir, submodulePath string) error {
	return RunGitCommand(parentDir, "submodule", "deinit", submodulePath)
}

// Remove removes a submodule from the Git index and .gitmodules.
// It performs a `git rm` which removes from index, .gitmodules, and work tree.
func Remove(parentDir, submodulePath string) error {
	return RunGitCommand(parentDir, "rm", submodulePath)
}

// IsPathDirty reports whether the given path has local modifications in repo at dir.
// It checks only the specified path (e.g., a submodule directory) via porcelain output.
func IsPathDirty(dir, path string) (bool, error) {
	out, err := RunGitCommandOutput(dir, "status", "--porcelain", "--", path)
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(out) != "", nil
}

// HasUnpushedCommits checks if the repository has commits that haven't been pushed to the remote.
func HasUnpushedCommits(dir string) (bool, error) {
	out, err := RunGitCommandOutput(dir, "log", "--oneline", "origin..HEAD")
	if err != nil {
		// If there's no remote or no origin branch, consider it as no unpushed commits
		if strings.Contains(err.Error(), "does not have a commit") ||
			strings.Contains(err.Error(), "unknown revision") ||
			strings.Contains(err.Error(), "No remote configured") {
			return false, nil
		}
		return false, err
	}
	return strings.TrimSpace(out) != "", nil
}
