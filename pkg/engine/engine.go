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
	atelierDirName := "atelier-" + atelierBaseName
	atelierPath = filepath.Join(basePath, atelierDirName)

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
	if err = fs.WriteFile(filepath.Join(atelierPath, ".atelier"), []byte(atelierDirName)); err != nil {
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
		"GEMINI.md",
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
	atelierDirName := filepath.Base(atelierPath)
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
	artistContext := fmt.Sprintf("%s\n%s", atelierDirName, artistDirName)
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
		"GEMINI.md",
		"Makefile",
		".gitignore",
		".geminiignore",
	})...); err != nil {
		return err
	}
	if err = gitutil.Commit(artistPath, fmt.Sprintf("feat: initialize artist %s", artistName)); err != nil {
		return err
	}

	// 2. Create and initialize default Canvas for the artist
	if err = CreateCanvas(artistPath, canvasName); err != nil {
		return err
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
	atelierName := artistLines[0]
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
	canvasContext := fmt.Sprintf("%s\n%s\n%s", atelierName, artistDirName, canvasDirName)
	if err = fs.WriteFile(filepath.Join(canvasPath, ".canvas"), []byte(canvasContext)); err != nil {
		return err
	}
	// Create boilerplate files
	if err = templates.CreateBoilerplate(canvasPath, "canvas"); err != nil {
		return err
	}
	// Stage changes (marker + boilerplate)
	if err = gitutil.AddPaths(canvasPath, existingPaths(canvasPath, []string{
		".canvas",
		"README.md",
		"GEMINI.md",
		"Makefile",
		".gitignore",
		".geminiignore",
	})...); err != nil {
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
