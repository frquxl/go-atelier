package gitutil

import (
	"fmt"
	"os"
	"os/exec"
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
