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

	// Test Case 1: Basic init command
	t.Run("basic_init", func(t *testing.T) {
		cmd := exec.Command(cliPath, "init", atelierName, artistName, canvasName)
		cmd.Dir = tmpDir // Run the command in the temp directory

		if output, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("Command `init` failed: %v\nOutput:\n%s", err, string(output))
		}

		// Define expected paths
		atelierPath := filepath.Join(tmpDir, "atelier-"+atelierName)
		primaryArtistPath := filepath.Join(atelierPath, "artist-"+artistName)
		primaryCanvasPath := filepath.Join(primaryArtistPath, "canvas-"+canvasName)

		// Assertions for Atelier
		assertDirExists(t, atelierPath)
		assertGitRepo(t, atelierPath)
		assertFileExists(t, filepath.Join(atelierPath, ".atelier"))
		assertFileContains(t, filepath.Join(atelierPath, ".atelier"), "atelier-"+atelierName)
		assertSubmodule(t, atelierPath, "artist-"+artistName)

		// Assertions for Primary Artist
		assertDirExists(t, primaryArtistPath)
		assertGitRepo(t, primaryArtistPath)
		assertFileExists(t, filepath.Join(primaryArtistPath, ".artist"))
		assertFileContains(t, filepath.Join(primaryArtistPath, ".artist"), "artist-"+artistName)
		assertSubmodule(t, primaryArtistPath, "canvas-"+canvasName)

		// Assertions for Primary Canvas
		assertDirExists(t, primaryCanvasPath)
		assertGitRepo(t, primaryCanvasPath)
		assertFileExists(t, filepath.Join(primaryCanvasPath, ".canvas"))
		assertFileContains(t, filepath.Join(primaryCanvasPath, ".canvas"), "canvas-"+canvasName)
	})

	// Test Case 2: init command with --sketch flag
	t.Run("init_with_sketch", func(t *testing.T) {
		// Use a fresh temp directory for each subtest
		subTmpDir := t.TempDir()
		subAtelierName := "test-sketch-atelier"

		cmd := exec.Command(cliPath, "init", subAtelierName, "--sketch")
		cmd.Dir = subTmpDir
		if output, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("Command `init --sketch` failed: %v\nOutput:\n%s", err, string(output))
		}

		subAtelierPath := filepath.Join(subTmpDir, "atelier-"+subAtelierName)
		vanGoghPath := filepath.Join(subAtelierPath, "artist-van-gogh")
		sketchPath := filepath.Join(subAtelierPath, "artist-sketch")

		// Assertions for van-gogh (default artist)
		assertDirExists(t, vanGoghPath)
		assertSubmodule(t, subAtelierPath, "artist-van-gogh")

		// Assertions for sketch artist
		assertDirExists(t, sketchPath)
		assertGitRepo(t, sketchPath)
		assertFileExists(t, filepath.Join(sketchPath, ".artist"))
		assertFileContains(t, filepath.Join(sketchPath, ".artist"), "artist-sketch")
		assertSubmodule(t, subAtelierPath, "artist-sketch")

		// Assert default canvas for sketch artist
		assertDirExists(t, filepath.Join(sketchPath, "canvas-example"))
		assertSubmodule(t, sketchPath, "canvas-example")
	})

	// Test Case 3: init command with --gallery flag
	t.Run("init_with_gallery", func(t *testing.T) {
		subTmpDir := t.TempDir()
		subAtelierName := "test-gallery-atelier"

		cmd := exec.Command(cliPath, "init", subAtelierName, "--gallery")
		cmd.Dir = subTmpDir
		if output, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("Command `init --gallery` failed: %v\nOutput:\n%s", err, string(output))
		}

		subAtelierPath := filepath.Join(subTmpDir, "atelier-"+subAtelierName)
		vanGoghPath := filepath.Join(subAtelierPath, "artist-van-gogh")
		galleryPath := filepath.Join(subAtelierPath, "artist-gallery")

		// Assertions for van-gogh (default artist)
		assertDirExists(t, vanGoghPath)
		assertSubmodule(t, subAtelierPath, "artist-van-gogh")

		// Assertions for gallery artist
		assertDirExists(t, galleryPath)
		assertGitRepo(t, galleryPath)
		assertFileExists(t, filepath.Join(galleryPath, ".artist"))
		assertFileContains(t, filepath.Join(galleryPath, ".artist"), "artist-gallery")
		assertSubmodule(t, subAtelierPath, "artist-gallery")

		// Assert default canvas for gallery artist
		assertDirExists(t, filepath.Join(galleryPath, "canvas-example"))
		assertSubmodule(t, galleryPath, "canvas-example")
	})

	// Test Case 4: init command with --sketch and --gallery flags
	t.Run("init_with_sketch_and_gallery", func(t *testing.T) {
		subTmpDir := t.TempDir()
		subAtelierName := "test-all-atelier"

		cmd := exec.Command(cliPath, "init", subAtelierName, "--sketch", "--gallery")
		cmd.Dir = subTmpDir
		if output, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("Command `init --sketch --gallery` failed: %v\nOutput:\n%s", err, string(output))
		}

		subAtelierPath := filepath.Join(subTmpDir, "atelier-"+subAtelierName)
		vanGoghPath := filepath.Join(subAtelierPath, "artist-van-gogh")
		sketchPath := filepath.Join(subAtelierPath, "artist-sketch")
		galleryPath := filepath.Join(subAtelierPath, "artist-gallery")

		// Assertions for van-gogh (default artist)
		assertDirExists(t, vanGoghPath)
		assertSubmodule(t, subAtelierPath, "artist-van-gogh")

		// Assertions for sketch artist
		assertDirExists(t, sketchPath)
		assertSubmodule(t, subAtelierPath, "artist-sketch")

		// Assertions for gallery artist
		assertDirExists(t, galleryPath)
		assertSubmodule(t, subAtelierPath, "artist-gallery")

		// Assert default canvases for sketch and gallery artists
		assertDirExists(t, filepath.Join(sketchPath, "canvas-example"))
		assertSubmodule(t, sketchPath, "canvas-example")
		assertDirExists(t, filepath.Join(galleryPath, "canvas-example"))
		assertSubmodule(t, galleryPath, "canvas-example")
	})
}