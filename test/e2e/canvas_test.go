package e2e

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestCanvasCommand(t *testing.T) {
	tmpDir := t.TempDir()

	// First, create a default atelier and artist to run the command in.
	// The init command creates a default "van-gogh" artist.
	initCmd := exec.Command(cliPath, "init", "test-canvas-atelier")
	initCmd.Dir = tmpDir
	if output, err := initCmd.CombinedOutput(); err != nil {
		t.Fatalf("Prerequisite `init` command failed: %v\nOutput:\n%s", err, string(output))
	}

	// The actual path to the artist directory created by the init command
	atelierPath := filepath.Join(tmpDir, "atelier-test-canvas-atelier")
	artistPath := filepath.Join(atelierPath, "artist-van-gogh")
	canvasName := "my-new-canvas"

	// Execute the canvas init command from within the artist directory
	canvasCmd := exec.Command(cliPath, "canvas", "init", canvasName)
	canvasCmd.Dir = artistPath // Run from inside the artist dir

	if output, err := canvasCmd.CombinedOutput(); err != nil {
		t.Fatalf("Command `canvas init` failed: %v\nOutput:\n%s", err, string(output))
	}

	// Define expected paths
	canvasPath := filepath.Join(artistPath, "canvas-"+canvasName)

	// Assertions for Canvas
	assertDirExists(t, canvasPath)
	assertGitRepo(t, canvasPath)
	assertFileExists(t, filepath.Join(canvasPath, ".canvas"))
	assertFileContains(t, filepath.Join(canvasPath, ".canvas"), canvasName)

	// Assert that the canvas was added as a submodule to the artist
	assertSubmodule(t, artistPath, "canvas-"+canvasName)

	// Test context awareness: should fail outside an artist directory
	ctxCmd := exec.Command(cliPath, "canvas", "init", "should-fail")
	ctxCmd.Dir = atelierPath // Run from one level above the artist
	if err := ctxCmd.Run(); err == nil {
		t.Fatal("Command `canvas init` should have failed outside an artist directory, but it didn't")
	}
}