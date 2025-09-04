package cmd

import (
	"atelier-cli/pkg/fs"
	"atelier-cli/pkg/gitutil"
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var templatesFS embed.FS

var initCmd = &cobra.Command{
	Use:   "init <atelier-name> [<artist-name> <canvas-name>]",
	Short: "Initialize a new atelier workspace",
	Long: `Initialize a new atelier workspace with 3-level Git submodule structure.
Creates atelier-<atelier-name> as main repo, artist as submodule, canvas as submodule of artist.
If no artist/canvas provided, defaults to 'van-gogh' and 'sunflowers'.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) == 0 {
			return fmt.Errorf("atelier name is required")
		}

		atelierBaseName := args[0]
		artistName := "van-gogh"
		canvasName := "sunflowers"

		if len(args) >= 3 {
			artistName = args[1]
			canvasName = args[2]
		}

		// Define directory names
		atelierDirName := "atelier-" + atelierBaseName
		artistDirName := "artist-" + artistName
		canvasDirName := "canvas-" + canvasName

		// Get current working directory to construct absolute paths
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current working directory: %w", err)
		}

		// Define full paths
		atelierPath := filepath.Join(wd, atelierDirName)
		artistPath := filepath.Join(atelierPath, artistDirName)
		canvasPath := filepath.Join(artistPath, canvasDirName)

		// Cleanup on failure
		defer func() {
			if err != nil {
				fmt.Printf("Initialization failed, cleaning up directory: %s\n", atelierPath)
				os.RemoveAll(atelierPath)
			}
		}()

		// 1. Create and initialize Atelier
		fmt.Println("Initializing atelier...")
		if err = fs.CreateDir(atelierPath); err != nil {
			return err
		}
		if err = gitutil.Init(atelierPath); err != nil {
			return err
		}
		if err = fs.WriteFile(filepath.Join(atelierPath, ".atelier"), []byte(atelierBaseName)); err != nil {
			return err
		}
		if err = createBoilerplate(atelierPath, "atelier"); err != nil {
			return err
		}

		// 2. Create and initialize Artist (as a standalone repo first)
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
		if err = createBoilerplate(artistPath, "artist"); err != nil {
			return err
		}
		if err = gitutil.Add(artistPath); err != nil {
			return err
		}
		if err = gitutil.Commit(artistPath, fmt.Sprintf("feat: initialize artist %s", artistName)); err != nil {
			return err
		}

		// 3. Create and initialize Canvas (as a standalone repo first)
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
		if err = createBoilerplate(canvasPath, "canvas"); err != nil {
			return err
		}
		if err = gitutil.Add(canvasPath); err != nil {
			return err
		}
		if err = gitutil.Commit(canvasPath, fmt.Sprintf("feat: initialize canvas %s", canvasName)); err != nil {
			return err
		}

		// 4. Link everything together with submodules
		fmt.Println("Connecting repositories as submodules...")
		// Add canvas to artist
		if err = gitutil.AddSubmodule(artistPath, canvasDirName); err != nil {
			return err
		}
		if err = gitutil.Add(artistPath); err != nil {
			return err
		}
		if err = gitutil.Commit(artistPath, fmt.Sprintf("feat: add canvas %s as submodule", canvasName)); err != nil {
			return err
		}

		// Add artist to atelier
		if err = gitutil.AddSubmodule(atelierPath, artistDirName); err != nil {
			return err
		}
		if err = gitutil.Add(atelierPath); err != nil {
			return err
		}
		if err = gitutil.Commit(atelierPath, fmt.Sprintf("feat: add artist %s as submodule", artistName)); err != nil {
			return err
		}

		fmt.Printf("Atelier '%s' initialized successfully!\n", atelierBaseName)
		fmt.Printf("  - Path: %s\n", atelierPath)
		fmt.Printf("  - Contains artist: %s\n", artistName)
		fmt.Printf("  - Which contains canvas: %s\n", canvasName)

		return nil
	},
}

func createBoilerplate(basePath, projectType string) error {
	// README
	readmePath := fmt.Sprintf("templates/%s/README.md", projectType)
	readmeContent, err := templatesFS.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", readmePath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, "README.md"), readmeContent); err != nil {
		return err
	}

	// GEMINI.md
	geminiPath := fmt.Sprintf("templates/%s/GEMINI.md", projectType)
	geminiContent, err := templatesFS.ReadFile(geminiPath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", geminiPath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, "GEMINI.md"), geminiContent); err != nil {
		return err
	}

	// .gitignore
	gitignorePath := fmt.Sprintf("templates/%s/gitignore", projectType)
	gitignoreContent, err := templatesFS.ReadFile(gitignorePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", gitignorePath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, ".gitignore"), gitignoreContent); err != nil {
		return err
	}

	// .geminiignore
	geminiignorePath := fmt.Sprintf("templates/%s/geminiignore", projectType)
	geminiignoreContent, err := templatesFS.ReadFile(geminiignorePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded template %s: %w", geminiignorePath, err)
	}
	if err := fs.WriteFile(filepath.Join(basePath, ".geminiignore"), geminiignoreContent); err != nil {
		return err
	}

	return nil
}

func init() {
	RootCmd.AddCommand(initCmd)
}
