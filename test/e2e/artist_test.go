package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestArtistCommand(t *testing.T) {
	// Test Case 1: Basic artist init and context awareness
	t.Run("basic_init_and_context", func(t *testing.T) {
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
	})

	// Test Case 2: Delete an artist
	t.Run("delete_artist", func(t *testing.T) {
		subTmpDir := t.TempDir()
		subAtelierName := "test-delete-atelier"
		artistToDeleteName := "artist-to-delete"

		// Setup: Create an atelier and the artist to be deleted
		initCmd := exec.Command(cliPath, "init", subAtelierName)
		initCmd.Dir = subTmpDir
		if output, err := initCmd.CombinedOutput(); err != nil {
			t.Fatalf("Prerequisite `init` command failed: %v\nOutput:\n%s", err, string(output))
		}

		subAtelierPath := filepath.Join(subTmpDir, "atelier-"+subAtelierName)
		createArtistCmd := exec.Command(cliPath, "artist", "init", artistToDeleteName)
		createArtistCmd.Dir = subAtelierPath
		if output, err := createArtistCmd.CombinedOutput(); err != nil {
			t.Fatalf("Prerequisite `artist init` command failed: %v\nOutput:\n%s", err, string(output))
		}

		// Verify artist exists before deletion
		artistToDeletePath := filepath.Join(subAtelierPath, "artist-"+artistToDeleteName)
		assertDirExists(t, artistToDeletePath)
		assertSubmodule(t, subAtelierPath, "artist-"+artistToDeleteName)

		// Execute the artist delete command
		deleteCmd := exec.Command(cliPath, "artist", "delete", "artist-"+artistToDeleteName)
		deleteCmd.Dir = subAtelierPath // Run from inside the atelier
		deleteCmd.Stdin = strings.NewReader("yes\n") // Provide confirmation

		if output, err := deleteCmd.CombinedOutput(); err != nil {
			t.Fatalf("Command `artist delete` failed: %v\nOutput:\n%s", err, string(output))
		}

		// Assertions after deletion
		if _, err := os.Stat(artistToDeletePath); !os.IsNotExist(err) {
			t.Fatalf("Artist directory %s was not deleted", artistToDeletePath)
		}

		// Check .gitmodules content (should not contain the deleted artist)
		gitmodulesContent, err := os.ReadFile(filepath.Join(subAtelierPath, ".gitmodules"))
		if err != nil {
			t.Fatalf("Failed to read .gitmodules: %v", err)
		}
		if strings.Contains(string(gitmodulesContent), "artist-"+artistToDeleteName) {
			t.Fatalf(".gitmodules still contains reference to deleted artist %s", artistToDeleteName)
		}
	})
}