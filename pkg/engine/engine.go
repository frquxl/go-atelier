package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/frquxl/go-atelier/pkg/fs"
	"github.com/frquxl/go-atelier/pkg/gitutil"
	"github.com/frquxl/go-atelier/pkg/templates"
)

// CreateAtelier initializes the main atelier directory and repository.
func CreateAtelier(basePath, atelierBaseName string) (atelierPath string, err error) {
	ateliersDirName := "atelier-" + atelierBaseName
	atelierPath = filepath.Join(basePath, ateliersDirName)

	_ = templates.TemplatesFS // Dummy reference to prevent unused import error

	defer func() {
		if err != nil {
			fmt.Printf("Atelier initialization failed, cleaning up directory: %s\n", atelierPath)
			os.RemoveAll(atelierPath)
		}
	}()

	fmt.Println("Initializing atelier...")
	if err = fs.CreateDir(atelierPath); err != nil {
		return "", err
	}
	if err = gitutil.Init(atelierPath); err != nil {
		return "", err
	}
	// Write marker file
	if err = fs.WriteFile(filepath.Join(atelierPath, ".atelier"), []byte(ateliersDirName)); err != nil {
		return "", err
	}
	// Create boilerplate files
	if err = templates.CreateBoilerplate(atelierPath, "atelier"); err != nil {
		return "", err
	}
	// An initial commit is needed before adding submodules.
	if err = gitutil.AddPaths(atelierPath, existingPaths(atelierPath, []string{
		".atelier",
		"README.md",
		"AGENTS.md",
		"Makefile",
		".gitignore",
		".geminiignore",
	})...); err != nil {
		return "", err
	}
	if err = gitutil.Commit(atelierPath, fmt.Sprintf("feat: initialize atelier %s", atelierBaseName)); err != nil {
		return "", err
	}

	return atelierPath, nil
}

// CreateArtist initializes a new artist and a default canvas within an atelier.
func CreateArtist(atelierPath, artistName, canvasName string) (err error) {
	ateliersDirName := filepath.Base(atelierPath)
	artistDirName := "artist-" + artistName
	artistPath := filepath.Join(atelierPath, artistDirName)

	defer func() {
		if err != nil {
			fmt.Printf("Artist initialization failed, cleaning up directory: %s\n", artistPath)
			os.RemoveAll(artistPath)
		}
	}()

	fmt.Println("Initializing artist...")
	if err = fs.CreateDir(artistPath); err != nil {
		return err
	}
	if err = gitutil.Init(artistPath); err != nil {
		return err
	}
	// Write marker file
	artistContext := fmt.Sprintf("%s\n%s", ateliersDirName, artistDirName)
	if err = fs.WriteFile(filepath.Join(artistPath, ".artist"), []byte(artistContext)); err != nil {
		return err
	}

	// Determine which template to use based on artistName
	templateType := "artist-default"
	if artistName == "sketch" {
		templateType = "artist-sketch"
	} else if artistName == "gallery" {
		templateType = "artist-gallery"
	}

	// Create boilerplate files
	if err = templates.CreateBoilerplate(artistPath, templateType); err != nil {
		return err
	}
	// Stage changes (marker + boilerplate)
	if err = gitutil.AddPaths(artistPath, existingPaths(artistPath, []string{
		".artist",
		"README.md",
		"AGENTS.md",
		"Makefile",
		".gitignore",
		".geminiignore",
	})...); err != nil {
		return err
	}
	if err = gitutil.Commit(artistPath, fmt.Sprintf("feat: initialize artist %s", artistName)); err != nil {
		return err
	}

	// 2. Create and initialize default Canvas for the artist (if specified)
	if canvasName != "" {
		if err = CreateCanvas(artistPath, canvasName); err != nil {
			return err
		}
	}

	// 3. Link artist to atelier
	fmt.Println("Connecting artist to atelier...")
	if err = gitutil.AddSubmodule(atelierPath, artistDirName); err != nil {
		return err
	}
	// Stage .gitmodules and submodule path
	if err = gitutil.AddPaths(atelierPath, ".gitmodules", artistDirName); err != nil {
		return err
	}
	if err = gitutil.Commit(atelierPath, fmt.Sprintf("feat: add artist %s as submodule", artistName)); err != nil {
		return err
	}

	return nil
}

// CreateCanvas initializes a new canvas within an artist's workspace.
func CreateCanvas(artistPath string, canvasName string) (err error) {
	artistContent, err := os.ReadFile(filepath.Join(artistPath, ".artist"))
	if err != nil {
		return fmt.Errorf("could not read .artist file in %s: %w", artistPath, err)
	}
	artistLines := strings.Split(strings.TrimSpace(string(artistContent)), "\n")
	if len(artistLines) < 2 {
		return fmt.Errorf("invalid .artist file format")
	}
	ateliersName := artistLines[0]
	artistDirName := artistLines[1]

	canvasDirName := "canvas-" + canvasName
	canvasPath := filepath.Join(artistPath, canvasDirName)

	defer func() {
		if err != nil {
			fmt.Printf("Canvas initialization failed, cleaning up directory: %s\n", canvasPath)
			os.RemoveAll(canvasPath)
		}
	}()

	fmt.Println("Initializing canvas...")
	if err = fs.CreateDir(canvasPath); err != nil {
		return err
	}
	if err = gitutil.Init(canvasPath); err != nil {
		return err
	}
	// Write marker file
	canvasContext := fmt.Sprintf("%s\n%s\n%s", ateliersName, artistDirName, canvasDirName)
	if err = fs.WriteFile(filepath.Join(canvasPath, ".canvas"), []byte(canvasContext)); err != nil {
		return err
	}
	// Create boilerplate files
	if err = templates.CreateBoilerplate(canvasPath, "canvas"); err != nil {
		return err
	}

	// --- Special handling for 'sunflowers' canvas ---
	if canvasName == "sunflowers" {
		fmt.Println("Configuring 'sunflowers' canvas with Van Gogh CLI...")
		if err = copySunflowersAssets(canvasPath); err != nil {
			return fmt.Errorf("failed to copy 'sunflowers' assets: %w", err)
		}
	}

	// Stage changes (marker + boilerplate + special assets)
	pathsToCommit := []string{
		".canvas",
		"README.md",
		"AGENTS.md",
		"Makefile",
		".gitignore",
		".geminiignore",
	}
	if canvasName == "sunflowers" {
		pathsToCommit = append(pathsToCommit, "vincent")
	}

	if err = gitutil.AddPaths(canvasPath, existingPaths(canvasPath, pathsToCommit)...); err != nil {
		return err
	}
	if err = gitutil.Commit(canvasPath, fmt.Sprintf("feat: initialize canvas %s", canvasName)); err != nil {
		return err
	}

	// 2. Link canvas to artist
	fmt.Println("Connecting canvas to artist...")
	if err = gitutil.AddSubmodule(artistPath, canvasDirName); err != nil {
		return err
	}
	// Stage .gitmodules and the submodule path
	if err = gitutil.AddPaths(artistPath, ".gitmodules", canvasDirName); err != nil {
		return err
	}
	if err = gitutil.Commit(artistPath, fmt.Sprintf("feat: add canvas %s as submodule", canvasName)); err != nil {
		return err
	}

	return nil
}

// copySunflowersAssets copies the specific assets for the sunflowers canvas.
func copySunflowersAssets(destPath string) error {
	// Define the files to copy from the embedded assets
	filesToCopy := map[string]os.FileMode{
		"assets/canvas-sunflowers/README.md": 0644,
		"assets/canvas-sunflowers/vincent":   0755, // Executable
	}

	for src, perm := range filesToCopy {
		content, err := templates.TemplatesFS.ReadFile(src)
		if err != nil {
			return fmt.Errorf("failed to read embedded asset %s: %w", src, err)
		}

		dest := filepath.Join(destPath, filepath.Base(src))

		// Overwrite the default README with the themed one, and write the binary.
		if err := os.WriteFile(dest, content, perm); err != nil {
			return fmt.Errorf("failed to write asset %s: %w", dest, err)
		}
	}
	return nil
}

// DeleteArtist deletes an artist studio and removes it from Git tracking.
func DeleteArtist(atelierPath, artistFullName string) (err error) {
	artistPath := filepath.Join(atelierPath, artistFullName)

	defer func() {
		if err != nil {
			fmt.Printf("Artist deletion failed, directory %s might need manual cleanup.\n", artistPath)
		}
	}()

	fmt.Printf("Deleting artist %s...\n", artistFullName)

	// 1. Deinitialize the submodule
	if err = gitutil.SubmoduleDeinit(atelierPath, artistFullName); err != nil {
		return fmt.Errorf("failed to deinitialize artist submodule: %w", err)
	}

	// 2. Remove the submodule entry from .gitmodules and index
	if err = gitutil.Remove(atelierPath, artistFullName); err != nil {
		return fmt.Errorf("failed to remove artist from git tracking: %w", err)
	}

	// 3. Remove the actual directory
	if err = os.RemoveAll(artistPath); err != nil {
		return fmt.Errorf("failed to remove artist directory: %w", err)
	}

	// No automatic commit here. User is responsible for committing the changes.
	fmt.Printf("Artist '%s' deleted. Remember to 'git add %s' and 'git commit' in the parent repository.\n", artistFullName, artistFullName)
	return nil
}

// DeleteCanvas deletes a canvas and removes it from Git tracking.
func DeleteCanvas(artistPath, canvasFullName string) (err error) {
	canvasPath := filepath.Join(artistPath, canvasFullName)

	defer func() {
		if err != nil {
			fmt.Printf("Canvas deletion failed, directory %s might need manual cleanup.\n", canvasPath)
		}
	}()

	fmt.Printf("Deleting canvas %s...\n", canvasFullName)

	// 1. Deinitialize the submodule
	if err = gitutil.SubmoduleDeinit(artistPath, canvasFullName); err != nil {
		return fmt.Errorf("failed to deinitialize canvas submodule: %w", err)
	}

	// 2. Remove the submodule entry from .gitmodules and index
	if err = gitutil.Remove(artistPath, canvasFullName); err != nil {
		return fmt.Errorf("failed to remove canvas from git tracking: %w", err)
	}

	// 3. Remove the actual directory
	if err = os.RemoveAll(canvasPath); err != nil {
		return fmt.Errorf("failed to remove canvas directory: %w", err)
	}

	// No automatic commit here. User is responsible for committing the changes.
	fmt.Printf("Canvas '%s' deleted. Remember to 'git add %s' and 'git commit' in the parent repository.\n", canvasFullName, canvasFullName)
	return nil
}

// MoveCanvas moves a canvas from one artist to another.
func MoveCanvas(canvasFullName, newArtistFullName string) error {
	// Find the atelier root by walking up from current directory
	atelierPath, err := findAtelierRoot()
	if err != nil {
		return fmt.Errorf("could not find atelier root: %w", err)
	}

	// Find which artist currently contains the canvas
	currentArtistPath, err := findCanvasArtist(atelierPath, canvasFullName)
	if err != nil {
		return fmt.Errorf("could not find artist containing canvas %s: %w", canvasFullName, err)
	}

	// Validate that the new artist exists
	newArtistPath := filepath.Join(atelierPath, newArtistFullName)
	if _, err := os.Stat(newArtistPath); os.IsNotExist(err) {
		return fmt.Errorf("new artist %s does not exist", newArtistFullName)
	}

	// Check if new artist already has a canvas with this name
	newCanvasPath := filepath.Join(newArtistPath, canvasFullName)
	if _, err := os.Stat(newCanvasPath); err == nil {
		return fmt.Errorf("canvas %s already exists in artist %s", canvasFullName, newArtistFullName)
	}

	// Get current artist name for context
	currentArtistName := filepath.Base(currentArtistPath)

	fmt.Printf("Moving canvas %s from artist %s to artist %s...\n", canvasFullName, currentArtistName, newArtistFullName)

	// 1. Remove canvas from current artist's git tracking (but keep the directory)
	// First, remove from index but keep working directory
	if err = gitutil.RunGitCommand(currentArtistPath, "rm", "--cached", canvasFullName); err != nil {
		return fmt.Errorf("failed to remove canvas from current artist's git index: %w", err)
	}

	// Remove from .gitmodules manually
	if err = removeFromGitmodules(currentArtistPath, canvasFullName); err != nil {
		return fmt.Errorf("failed to remove canvas from .gitmodules: %w", err)
	}

	// 2. Move the canvas directory to the new artist
	canvasPath := filepath.Join(currentArtistPath, canvasFullName)
	if err = os.Rename(canvasPath, newCanvasPath); err != nil {
		return fmt.Errorf("failed to move canvas directory: %w", err)
	}

	// 3. Update the .canvas file with new artist context
	if err = updateCanvasContext(newCanvasPath, newArtistFullName); err != nil {
		return fmt.Errorf("failed to update canvas context: %w", err)
	}

	// 4. Add canvas as submodule to new artist
	if err = gitutil.AddSubmodule(newArtistPath, canvasFullName); err != nil {
		return fmt.Errorf("failed to add canvas as submodule to new artist: %w", err)
	}

	// 5. Stage changes in both artists
	if err = gitutil.AddPaths(currentArtistPath, ".gitmodules"); err != nil {
		return fmt.Errorf("failed to stage .gitmodules changes in current artist: %w", err)
	}
	if err = gitutil.AddPaths(newArtistPath, ".gitmodules", canvasFullName); err != nil {
		return fmt.Errorf("failed to stage changes in new artist: %w", err)
	}

	// 6. Commit changes in both artists
	if err = gitutil.Commit(currentArtistPath, fmt.Sprintf("feat: remove canvas %s (moved to %s)", canvasFullName, newArtistFullName)); err != nil {
		return fmt.Errorf("failed to commit changes in current artist: %w", err)
	}
	if err = gitutil.Commit(newArtistPath, fmt.Sprintf("feat: add canvas %s (moved from %s)", canvasFullName, currentArtistName)); err != nil {
		return fmt.Errorf("failed to commit changes in new artist: %w", err)
	}

	fmt.Printf("Canvas %s successfully moved from %s to %s!\n", canvasFullName, currentArtistName, newArtistFullName)
	return nil
}

// findAtelierRoot finds the atelier root directory by walking up from current directory
func findAtelierRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".atelier")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root directory
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("could not find atelier root (.atelier file not found)")
}

// findCanvasArtist finds which artist contains the specified canvas
func findCanvasArtist(atelierPath, canvasFullName string) (string, error) {
	entries, err := os.ReadDir(atelierPath)
	if err != nil {
		return "", fmt.Errorf("could not read atelier directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "artist-") {
			artistPath := filepath.Join(atelierPath, entry.Name())
			canvasPath := filepath.Join(artistPath, canvasFullName)
			if _, err := os.Stat(canvasPath); err == nil {
				return artistPath, nil
			}
		}
	}

	return "", fmt.Errorf("canvas %s not found in any artist", canvasFullName)
}

// updateCanvasContext updates the .canvas file with new artist context
func updateCanvasContext(canvasPath, newArtistFullName string) error {
	canvasFile := filepath.Join(canvasPath, ".canvas")
	content, err := os.ReadFile(canvasFile)
	if err != nil {
		return fmt.Errorf("could not read .canvas file: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	if len(lines) < 3 {
		return fmt.Errorf("invalid .canvas file format")
	}

	// Update the artist line (second line)
	lines[1] = newArtistFullName
	newContent := strings.Join(lines, "\n")

	if err = os.WriteFile(canvasFile, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("could not write updated .canvas file: %w", err)
	}

	return nil
}

// removeFromGitmodules removes a submodule entry from .gitmodules file
func removeFromGitmodules(repoPath, submodulePath string) error {
	gitmodulesPath := filepath.Join(repoPath, ".gitmodules")

	// Read the .gitmodules file
	content, err := os.ReadFile(gitmodulesPath)
	if err != nil {
		if os.IsNotExist(err) {
			// No .gitmodules file, nothing to remove
			return nil
		}
		return fmt.Errorf("could not read .gitmodules file: %w", err)
	}

	lines := strings.Split(string(content), "\n")
	var newLines []string
	inSubmoduleBlock := false

	for _, line := range lines {
		// Check if we're entering a submodule block for our target
		if strings.HasPrefix(line, "[submodule \"") && strings.Contains(line, submodulePath+"\"]") {
			inSubmoduleBlock = true
			continue // Skip this line
		}

		// If we're in the target submodule block, skip all indented lines
		if inSubmoduleBlock {
			if strings.HasPrefix(line, "\t") || strings.HasPrefix(line, " ") {
				continue // Skip indented lines in the submodule block
			} else {
				// We've exited the submodule block
				inSubmoduleBlock = false
			}
		}

		// Keep non-submodule lines
		if !inSubmoduleBlock {
			newLines = append(newLines, line)
		}
	}

	// Write the updated content back
	newContent := strings.Join(newLines, "\n")
	if err = os.WriteFile(gitmodulesPath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("could not write updated .gitmodules file: %w", err)
	}

	return nil
}

// existingPaths returns only those names that currently exist under base.
// It prevents staging non-existent files when generating boilerplate.
func existingPaths(base string, names []string) []string {
	out := []string{}
	for _, n := range names {
		if _, err := os.Stat(filepath.Join(base, n)); err == nil {
			out = append(out, n)
		}
	}
	return out
}
