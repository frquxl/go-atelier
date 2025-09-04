package cmd

import (
	"atelier-cli/pkg/fs"
	"atelier-cli/pkg/gitutil"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var canvasTemplatesFS embed.FS

var canvasCmd = &cobra.Command{
	Use:   "canvas",
	Short: "Manage canvases in the artist workspace",
	Long:  `Commands for managing canvases within the artist workspace.`,
}

var canvasInitCmd = &cobra.Command{
	Use:   "init <canvas-name>",
	Short: "Initialize a new canvas",
	Long:  `Initialize a new canvas within the current artist workspace as a Git submodule. Must be run from an artist directory.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// Check if we're in an artist directory
		if _, statErr := os.Stat(".artist"); os.IsNotExist(statErr) {
			listAvailableArtists()
			return fmt.Errorf("not in an artist directory. See available artists above")
		}

		canvasName := args[0]
		canvasDirName := "canvas-" + canvasName

		// Get current working directory to construct absolute paths
		artistPath, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current working directory: %w", err)
		}

		canvasPath := filepath.Join(artistPath, canvasDirName)

		// Cleanup on failure
		defer func() {
			if err != nil {
				fmt.Printf("Initialization failed, cleaning up directory: %s\n", canvasPath)
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
		if err = fs.WriteFile(filepath.Join(canvasPath, ".canvas"), []byte(canvasName)); err != nil {
			return err
		}
		if err = createCanvasBoilerplate(canvasPath, "canvas"); err != nil {
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

		fmt.Printf("Canvas '%s' initialized successfully in artist '%s'!\n", canvasName, filepath.Base(artistPath))
		return nil
	},
}

func createCanvasBoilerplate(basePath, projectType string) error {
	// README
	readmePath := fmt.Sprintf("templates/%s/README.md", projectType)
	readmeContent, err := canvasTemplatesFS.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", readmePath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, "README.md"), readmeContent); err != nil {
		return err
	}

	// GEMINI.md
	geminiPath := fmt.Sprintf("templates/%s/GEMINI.md", projectType)
	geminiContent, err := canvasTemplatesFS.ReadFile(geminiPath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", geminiPath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, "GEMINI.md"), geminiContent); err != nil {
		return err
	}

	// .gitignore
	gitignorePath := fmt.Sprintf("templates/%s/gitignore", projectType)
	gitignoreContent, err := canvasTemplatesFS.ReadFile(gitignorePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", gitignorePath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, ".gitignore"), gitignoreContent); err != nil {
		return err
	}

	// .geminiignore
	geminiignorePath := fmt.Sprintf("templates/%s/geminiignore", projectType)
	geminiignoreContent, err := canvasTemplatesFS.ReadFile(geminiignorePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", geminiignorePath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, ".geminiignore"), geminiignoreContent); err != nil {
		return err
	}

	return nil
}

func listAvailableArtists() {
	fmt.Println("Available artists in current atelier:")

	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	found := false
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "artist-") {
			artistName := strings.TrimPrefix(entry.Name(), "artist-")
			fmt.Printf("  - %s (cd %s)\n", artistName, entry.Name())
			found = true
		}
	}

	if !found {
		fmt.Println("  No artists found in current atelier.")
		fmt.Println("  Create one with: artist init <name>")
	} else {
		fmt.Println("\nTo work with an artist, run: cd <artist-directory>")
	}
}

func init() {
	RootCmd.AddCommand(canvasCmd)
	canvasCmd.AddCommand(canvasInitCmd)
}
