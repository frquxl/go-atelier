package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [<artist-name> <canvas-name>]",
	Short: "Initialize a new atelier workspace",
	Long: `Initialize a new atelier workspace with the basic skeleton structure.
If no arguments are provided, defaults to creating 'van-gogh' as the artist and 'sunflowers' as the canvas.`,
	Run: func(cmd *cobra.Command, args []string) {
		artist := "van-gogh"
		canvas := "sunflowers"

		if len(args) >= 2 {
			artist = args[0]
			canvas = args[1]
		}

		// Create directories
		atelierDir := "atelier"
		artistDir := filepath.Join(atelierDir, artist)
		canvasDir := filepath.Join(artistDir, canvas)

		dirs := []string{atelierDir, artistDir, canvasDir}

		for _, dir := range dirs {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", dir, err)
				return
			}
		}

		// Initialize Git repository
		if err := exec.Command("git", "init", atelierDir).Run(); err != nil {
			fmt.Printf("Error initializing git repository: %v\n", err)
			return
		}

		// Create boilerplate files
		createBoilerplateFiles(atelierDir, artistDir, canvasDir)

		fmt.Printf("Atelier initialized with artist '%s' and canvas '%s'\n", artist, canvas)
	},
}

func createBoilerplateFiles(dirs ...string) {
	for _, dir := range dirs {
		readmePath := filepath.Join(dir, "README.md")
		geminiPath := filepath.Join(dir, "GEMINI.md")

		// Simple template content
		readmeContent := fmt.Sprintf("# %s\n\nThis is a README for %s.\n", filepath.Base(dir), dir)
		geminiContent := fmt.Sprintf("# AI Context for %s\n\nThis file contains AI context for %s.\n", filepath.Base(dir), dir)

		if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
			fmt.Printf("Error creating %s: %v\n", readmePath, err)
		}

		if err := os.WriteFile(geminiPath, []byte(geminiContent), 0644); err != nil {
			fmt.Printf("Error creating %s: %v\n", geminiPath, err)
		}
	}
}

func init() {
	RootCmd.AddCommand(initCmd)
}
