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

		// Copy README template (find templates relative to executable)
		readmeTemplate := findTemplatePath(templateType, "README.md")
		if err := copyFile(readmeTemplate, readmePath); err != nil {
			fmt.Printf("Error creating %s: %v\n", readmePath, err)
		}

		// Copy GEMINI template
		geminiTemplate := findTemplatePath(templateType, "GEMINI.md")
		if err := copyFile(geminiTemplate, geminiPath); err != nil {
			fmt.Printf("Error creating %s: %v\n", geminiPath, err)
		}
	}
}

func findTemplatePath(templateType, filename string) string {
	// Get the executable path to find templates relative to CLI location
	execPath, err := os.Executable()
	if err != nil {
		// Fallback to current directory if we can't get executable path
		return filepath.Join("templates", templateType, filename)
	}

	// Go up one directory from the executable to find templates
	execDir := filepath.Dir(execPath)
	return filepath.Join(execDir, "templates", templateType, filename)
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
