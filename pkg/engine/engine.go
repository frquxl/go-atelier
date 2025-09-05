package engine

import (
	"atelier-cli/pkg/fs"
	"atelier-cli/pkg/gitutil"
	"atelier-cli/pkg/templates"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	if err = fs.WriteFile(filepath.Join(atelierPath, ".atelier"), []byte(atelierDirName)); err != nil {
		return "", err
	}
	if err = templates.CreateBoilerplate(atelierPath, "atelier"); err != nil {
		return "", err
	}
	// An initial commit is needed before adding submodules.
	if err = gitutil.Add(atelierPath); err != nil {
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

	// 1. Create and initialize Artist (as a standalone repo first)
	fmt.Println("Initializing artist...")
	if err = fs.CreateDir(artistPath); err != nil {
		return err
	}
	if err = gitutil.Init(artistPath); err != nil {
		return err
	}
	artistContext := fmt.Sprintf("%s\n%s", atelierDirName, artistDirName)
	if err = fs.WriteFile(filepath.Join(artistPath, ".artist"), []byte(artistContext)); err != nil {
		return err
	}
	if err = templates.CreateBoilerplate(artistPath, "artist"); err != nil {
		return err
	}
	if err = gitutil.Add(artistPath); err != nil {
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
	if err = gitutil.Add(atelierPath); err != nil {
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

	// 1. Create and initialize Canvas (as a standalone repo first)
	fmt.Println("Initializing canvas...")
	if err = fs.CreateDir(canvasPath); err != nil {
		return err
	}
	if err = gitutil.Init(canvasPath); err != nil {
		return err
	}
	canvasContext := fmt.Sprintf("%s\n%s\n%s", atelierName, artistDirName, canvasDirName)
	if err = fs.WriteFile(filepath.Join(canvasPath, ".canvas"), []byte(canvasContext)); err != nil {
		return err
	}
	if err = templates.CreateBoilerplate(canvasPath, "canvas"); err != nil {
		return err
	}
	if err = gitutil.Add(canvasPath); err != nil {
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
	if err = gitutil.Add(artistPath); err != nil {
		return err
	}
	if err = gitutil.Commit(artistPath, fmt.Sprintf("feat: add canvas %s as submodule", canvasName)); err != nil {
		return err
	}

	return nil
}