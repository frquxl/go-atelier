package e2e

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestInitCommand(t *testing.T) {
	tmpDir := t.TempDir() // Create a temporary dir that cleans itself up

	atelierName := "test-init-atelier"
	artistName := "dali"
	canvasName := "clocks"

	// Execute the init command
	cmd := exec.Command(cliPath, "init", atelierName, artistName, canvasName)
	cmd.Dir = tmpDir // Run the command in the temp directory

	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Command `init` failed: %v\nOutput:\n%s", err, string(output))
	}

	// Define expected paths
	atelierPath := filepath.Join(tmpDir, "atelier-"+atelierName)
	artistPath := filepath.Join(atelierPath, "artist-"+artistName)
	canvasPath := filepath.Join(artistPath, "canvas-"+canvasName)

	// Assertions for Atelier
	assertDirExists(t, atelierPath)
	assertGitRepo(t, atelierPath)
	assertFileExists(t, filepath.Join(atelierPath, ".atelier"))
	assertFileContains(t, filepath.Join(atelierPath, ".atelier"), "atelier-"+atelierName)
	assertSubmodule(t, atelierPath, "artist-"+artistName)

	// Assertions for Artist
	assertDirExists(t, artistPath)
	assertGitRepo(t, artistPath)
	assertFileExists(t, filepath.Join(artistPath, ".artist"))
	assertFileContains(t, filepath.Join(artistPath, ".artist"), "artist-"+artistName)
	assertSubmodule(t, artistPath, "canvas-"+canvasName)

	// Assertions for Canvas
	assertDirExists(t, canvasPath)
	assertGitRepo(t, canvasPath)
	assertFileExists(t, filepath.Join(canvasPath, ".canvas"))
	assertFileContains(t, filepath.Join(canvasPath, ".canvas"), "canvas-"+canvasName)
}
