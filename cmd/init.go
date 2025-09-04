package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <atelier-name> [<artist-name> <canvas-name>]",
	Short: "Initialize a new atelier workspace",
	Long: `Initialize a new atelier workspace with the basic skeleton structure.
Creates atelier-<atelier-name> directory. If no artist/canvas provided, defaults to 'van-gogh' and 'sunflowers'.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: atelier name is required")
			return
		}

		atelierBaseName := args[0]
		artist := "van-gogh"
		canvas := "sunflowers"

		if len(args) >= 3 {
			artist = args[1]
			canvas = args[2]
		}

		// Always create atelier-XXXXX format
		atelierDir := "atelier-" + atelierBaseName
		artistDir := "artist-" + artist
		canvasDir := "canvas-" + canvas

		fullArtistDir := filepath.Join(atelierDir, artistDir)
		fullCanvasDir := filepath.Join(fullArtistDir, canvasDir)

		dirs := []string{atelierDir, fullArtistDir, fullCanvasDir}

		for _, dir := range dirs {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", dir, err)
				return
			}
		}

		// Create marker files
		markerFiles := map[string]string{
			filepath.Join(atelierDir, ".atelier"):   atelierBaseName,
			filepath.Join(fullArtistDir, ".artist"): artist,
			filepath.Join(fullCanvasDir, ".canvas"): canvas,
		}

		for markerPath, content := range markerFiles {
			if err := os.WriteFile(markerPath, []byte(content), 0644); err != nil {
				fmt.Printf("Error creating marker file %s: %v\n", markerPath, err)
				return
			}
		}

		// Initialize Git repository
		if err := exec.Command("git", "init", atelierDir).Run(); err != nil {
			fmt.Printf("Error initializing git repository: %v\n", err)
			return
		}

		// Create boilerplate files
		createBoilerplateFiles(atelierDir, fullArtistDir, fullCanvasDir)

		fmt.Printf("Atelier '%s' initialized with artist '%s' and canvas '%s'\n", atelierBaseName, artist, canvas)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}

func createBoilerplateFiles(dirs ...string) {
	for _, dir := range dirs {
		readmePath := filepath.Join(dir, "README.md")
		geminiPath := filepath.Join(dir, "GEMINI.md")

		// Determine template type based on directory prefix
		var templateType string
		baseName := filepath.Base(dir)
		if strings.HasPrefix(baseName, "atelier-") {
			templateType = "atelier"
		} else if strings.HasPrefix(baseName, "artist-") {
			templateType = "artist"
		} else if strings.HasPrefix(baseName, "canvas-") {
			templateType = "canvas"
		} else {
			templateType = "canvas" // fallback
		}

		// Create README file with default content
		readmeContent := getDefaultContent(templateType, "README.md", dir)
		if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
			fmt.Printf("Error creating %s: %v\n", readmePath, err)
		}

		// Create GEMINI file with default content
		geminiContent := getDefaultContent(templateType, "GEMINI.md", dir)
		if err := os.WriteFile(geminiPath, []byte(geminiContent), 0644); err != nil {
			fmt.Printf("Error creating %s: %v\n", geminiPath, err)
		}
	}
}

func findTemplatePath(templateType, filename string) string {
	// For global installation, we need to find templates relative to the source
	// Since go install doesn't copy templates, we'll look in common locations

	// First, try relative to current working directory (for local development)
	if _, err := os.Stat(filepath.Join("templates", templateType, filename)); err == nil {
		return filepath.Join("templates", templateType, filename)
	}

	// For global installation, we need to embed templates or use a different approach
	// For now, create simple default content when templates aren't found
	return ""
}

func getDefaultContent(templateType, filename, dir string) string {
	dirName := filepath.Base(dir)

	switch filename {
	case "README.md":
		switch templateType {
		case "atelier":
			return fmt.Sprintf("# %s\n\nWelcome to your atelier workspace!\n\nThis is the root directory for your software projects.\n\n## Getting Started\n\n1. Add artists: `atelier artist init <artist-name>`\n2. Navigate to artists and add canvases\n3. Start developing!\n", dirName)
		case "artist":
			return fmt.Sprintf("# %s\n\nArtist workspace for creative development.\n\nThis directory contains your personal workspace and canvases.\n\n## Canvases\n\nAdd canvases with: `atelier canvas init <canvas-name>`\n", dirName)
		case "canvas":
			return fmt.Sprintf("# %s\n\nProject canvas for development.\n\nThis is where your actual project code and files go.\n\n## Getting Started\n\nStart developing your project here!\n", dirName)
		}
	case "GEMINI.md":
		switch templateType {
		case "atelier":
			return fmt.Sprintf("# AI Context for %s\n\nThis atelier contains multiple artists and their canvases.\n\n## Structure\n\n- Artists: Individual workspaces\n- Canvases: Project areas within artists\n\n## Commands\n\n- `atelier artist init <name>`: Add new artist\n- `atelier artist list`: List artists\n", dirName)
		case "artist":
			return fmt.Sprintf("# AI Context for %s\n\nArtist workspace containing multiple canvases.\n\n## Canvases\n\nThis artist can have multiple canvases for different projects.\n\n## Commands\n\n- `atelier canvas init <name>`: Add new canvas\n- `atelier canvas list`: List canvases\n", dirName)
		case "canvas":
			return fmt.Sprintf("# AI Context for %s\n\nProject canvas for development work.\n\n## Purpose\n\nThis canvas is where the actual development happens.\n\n## Technologies\n\nAdd your tech stack and project details here.\n", dirName)
		}
	}

	return fmt.Sprintf("# %s\n\nDefault content for %s.\n", dirName, filename)
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Ensure the destination file has the correct permissions
	return os.Chmod(dst, 0644)
}
