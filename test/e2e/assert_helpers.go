package e2e

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func assertDirExists(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("Directory should exist, but doesn't: %s", path)
	}
}

func assertFileExists(t *testing.T, path string) {
	t.Helper()
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		t.Fatalf("File should exist, but doesn't: %s", path)
	}
	if info.IsDir() {
		t.Fatalf("Path should be a file, but is a directory: %s", path)
	}
}

func assertFileContains(t *testing.T, path string, content string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", path, err)
	}
	if !strings.Contains(string(data), content) {
		t.Fatalf("File %s was expected to contain %q, but it didn't", path, content)
	}
}

func assertGitRepo(t *testing.T, path string) {
	t.Helper()
	gitPath := filepath.Join(path, ".git")
	if _, err := os.Stat(gitPath); os.IsNotExist(err) {
		t.Fatalf("Path should be a git repository, but .git directory is missing: %s", path)
	}
}

func assertSubmodule(t *testing.T, parentDir, submoduleName string) {
	t.Helper()
	gitmodulesPath := filepath.Join(parentDir, ".gitmodules")
	assertFileExists(t, gitmodulesPath)
	assertFileContains(t, gitmodulesPath, submoduleName)
}
