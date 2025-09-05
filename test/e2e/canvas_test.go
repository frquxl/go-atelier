package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCanvasCommand(t *testing.T) {
	// Test Case 1: Basic canvas init and context awareness
	t.Run("basic_init_and_context", func(t *testing.T) {
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
	})

	// Test Case 2: Delete a canvas
	t.Run("delete_canvas", func(t *testing.T) {
		subTmpDir := t.TempDir()
		subAtelierName := "test-delete-canvas-atelier"
		// artistForCanvasName := "test-delete-canvas-artist" // Removed unused variable
		canvasToDeleteName := "canvas-to-delete"

		// Setup: Create an atelier, an artist, and the canvas to be deleted
		initCmd := exec.Command(cliPath, "init", subAtelierName)
		initCmd.Dir = subTmpDir
		if output, err := initCmd.CombinedOutput(); err != nil {
			t.Fatalf("Prerequisite `init` command failed: %v\nOutput:\n%s", err, string(output))
		}

		subAtelierPath := filepath.Join(subTmpDir, "atelier-"+subAtelierName)
		artistPath := filepath.Join(subAtelierPath, "artist-van-gogh") // Default artist created by init

		createCanvasCmd := exec.Command(cliPath, "canvas", "init", canvasToDeleteName)
		createCanvasCmd.Dir = artistPath
		if output, err := createCanvasCmd.CombinedOutput(); err != nil {
			t.Fatalf("Prerequisite `canvas init` command failed: %v\nOutput:\n%s", err, string(output))
		}

		// Verify canvas exists before deletion
		canvasToDeletePath := filepath.Join(artistPath, "canvas-"+canvasToDeleteName)
		assertDirExists(t, canvasToDeletePath)
		assertSubmodule(t, artistPath, "canvas-"+canvasToDeleteName)

		// Execute the canvas delete command
		deleteCmd := exec.Command(cliPath, "canvas", "delete", "canvas-"+canvasToDeleteName)
		deleteCmd.Dir = artistPath // Run from inside the artist
		deleteCmd.Stdin = strings.NewReader("yes\n") // Provide confirmation

		if output, err := deleteCmd.CombinedOutput(); err != nil {
			t.Fatalf("Command `canvas delete` failed: %v\nOutput:\n%s", err, string(output))
		}

		// Assertions after deletion
		if _, err := os.Stat(canvasToDeletePath); !os.IsNotExist(err) {
			t.Fatalf("Canvas directory %s was not deleted", canvasToDeletePath)
		}

		// Check .gitmodules content (should not contain the deleted canvas)
		gitmodulesContent, err := os.ReadFile(filepath.Join(artistPath, ".gitmodules"))
		if err != nil {
			t.Fatalf("Failed to read .gitmodules: %v", err)
		}
		if strings.Contains(string(gitmodulesContent), "canvas-"+canvasToDeleteName) {
			t.Fatalf(".gitmodules still contains reference to deleted canvas %s", canvasToDeleteName)
		}
	})
}