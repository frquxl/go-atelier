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
var artistTemplatesFS embed.FS

var artistCmd = &cobra.Command{
	Use:   "artist",
	Short: "Manage artists in the atelier",
	Long:  `Commands for managing artists within the atelier workspace.`,
}

var artistInitCmd = &cobra.Command{
	Use:   "init <artist-name>",
	Short: "Initialize a new artist studio",
	Long:  `Initialize a new artist studio within the existing atelier as a Git submodule.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// Check if we're in an atelier directory
		if _, statErr := os.Stat(".atelier"); os.IsNotExist(statErr) {
			listAvailableAteliers()
			return fmt.Errorf("not in an atelier directory. See available ateliers above")
		}

		artistName := args[0]
		canvasName := "example"

		artistDirName := "artist-" + artistName
		canvasDirName := "canvas-" + canvasName

		// Get current working directory to construct absolute paths
		atelierPath, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current working directory: %w", err)
		}

		artistPath := filepath.Join(atelierPath, artistDirName)
		canvasPath := filepath.Join(artistPath, canvasDirName)

		// Cleanup on failure
		defer func() {
			if err != nil {
				fmt.Printf("Initialization failed, cleaning up directory: %s\n", artistPath)
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
		if err = fs.WriteFile(filepath.Join(artistPath, ".artist"), []byte(artistName)); err != nil {
			return err
		}
		if err = createArtistBoilerplate(artistPath, "artist"); err != nil {
			return err
		}
		if err = gitutil.Add(artistPath); err != nil {
			return err
		}
		if err = gitutil.Commit(artistPath, fmt.Sprintf("feat: initialize artist %s", artistName)); err != nil {
			return err
		}

		// 2. Create and initialize default Canvas
		fmt.Println("Initializing default canvas...")
		if err = fs.CreateDir(canvasPath); err != nil {
			return err
		}
		if err = gitutil.Init(canvasPath); err != nil {
			return err
		}
		if err = fs.WriteFile(filepath.Join(canvasPath, ".canvas"), []byte(canvasName)); err != nil {
			return err
		}
		if err = createArtistBoilerplate(canvasPath, "canvas"); err != nil {
			return err
		}
		if err = gitutil.Add(canvasPath); err != nil {
			return err
		}
		if err = gitutil.Commit(canvasPath, fmt.Sprintf("feat: initialize canvas %s", canvasName)); err != nil {
			return err
		}

		// 3. Link canvas to artist
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

		// 4. Link artist to atelier
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

		fmt.Printf("Artist '%s' initialized successfully in atelier '%s'!\n", artistName, filepath.Base(atelierPath))
		return nil
	},
}

func createArtistBoilerplate(basePath, projectType string) error {
	// README
	readmePath := fmt.Sprintf("templates/%s/README.md", projectType)
	readmeContent, err := artistTemplatesFS.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", readmePath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, "README.md"), readmeContent); err != nil {
		return err
	}

	// GEMINI.md
	geminiPath := fmt.Sprintf("templates/%s/GEMINI.md", projectType)
	geminiContent, err := artistTemplatesFS.ReadFile(geminiPath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", geminiPath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, "GEMINI.md"), geminiContent); err != nil {
		return err
	}

	// .gitignore
	gitignorePath := fmt.Sprintf("templates/%s/gitignore", projectType)
	gitignoreContent, err := artistTemplatesFS.ReadFile(gitignorePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", gitignorePath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, ".gitignore"), gitignoreContent); err != nil {
		return err
	}

	// .geminiignore
	geminiignorePath := fmt.Sprintf("templates/%s/geminiignore", projectType)
	geminiignoreContent, err := artistTemplatesFS.ReadFile(geminiignorePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", geminiignorePath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, ".geminiignore"), geminiignoreContent); err != nil {
		return err
	}

	// Makefile
	makefilePath := fmt.Sprintf("templates/%s/Makefile", projectType)
	makefileContent, err := artistTemplatesFS.ReadFile(makefilePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", makefilePath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, "Makefile"), makefileContent); err != nil {
		return err
	}

	return nil
}

func listAvailableAteliers() {
	fmt.Println("Available ateliers in current directory:")

	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	found := false
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "atelier-") {
			atelierName := strings.TrimPrefix(entry.Name(), "atelier-")
			fmt.Printf("  - %s (cd %s)\n", atelierName, entry.Name())
			found = true
		}
	}

	if !found {
		fmt.Println("  No ateliers found in current directory.")
		fmt.Println("  Create one with: atelier init <name>")
	} else {
		fmt.Println("\nTo work with an atelier, run: cd <atelier-directory>")
	}
}

func init() {
	RootCmd.AddCommand(artistCmd)
	artistCmd.AddCommand(artistInitCmd)
}
