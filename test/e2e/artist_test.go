package e2e

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestArtistCommand(t *testing.T) {
	tmpDir := t.TempDir()

	// First, create a bare atelier to run the command in
	initCmd := exec.Command(cliPath, "init", "test-artist-atelier")
	initCmd.Dir = tmpDir
	if output, err := initCmd.CombinedOutput(); err != nil {
		t.Fatalf("Prerequisite `init` command failed: %v\nOutput:\n%s", err, string(output))
	}

	atelierPath := filepath.Join(tmpDir, "atelier-test-artist-atelier")
	artistName := "test-artist"

	// Execute the artist init command from within the atelier directory
	artistCmd := exec.Command(cliPath, "artist", "init", artistName)
	artistCmd.Dir = atelierPath // Run from inside the atelier

	if output, err := artistCmd.CombinedOutput(); err != nil {
		t.Fatalf("Command `artist init` failed: %v\nOutput:\n%s", err, string(output))
	}

	// Define expected paths
	artistPath := filepath.Join(atelierPath, "artist-"+artistName)
	canvasPath := filepath.Join(artistPath, "canvas-example") // artist init creates a default 'example' canvas

	// Assertions for Artist
	assertDirExists(t, artistPath)
	assertGitRepo(t, artistPath)
	assertFileExists(t, filepath.Join(artistPath, ".artist"))
	assertFileContains(t, filepath.Join(artistPath, ".artist"), artistName)
	assertSubmodule(t, artistPath, "canvas-example")

	// Assertions for default Canvas
	assertDirExists(t, canvasPath)
	assertGitRepo(t, canvasPath)
	assertFileExists(t, filepath.Join(canvasPath, ".canvas"))
	assertFileContains(t, filepath.Join(canvasPath, ".canvas"), "canvas-example")

	// Assert that the artist was added as a submodule to the atelier
	assertSubmodule(t, atelierPath, "artist-"+artistName)

	// Test context awareness: should fail outside an atelier
	ctxCmd := exec.Command(cliPath, "artist", "init", "should-fail")
	ctxCmd.Dir = tmpDir // Run from one level above the atelier
	if err := ctxCmd.Run(); err == nil {
		t.Fatal("Command `artist init` should have failed outside an atelier directory, but it didn't")
	}
}