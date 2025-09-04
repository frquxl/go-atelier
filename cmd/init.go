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

		// Determine template type based on directory level
		var templateType string
		switch filepath.Base(dir) {
		case "atelier":
			templateType = "atelier"
		default:
			// For artist and canvas directories, use their respective templates
			if strings.Contains(dir, "/") {
				parts := strings.Split(dir, "/")
				if len(parts) >= 2 && parts[len(parts)-2] == "atelier" {
					templateType = "artist"
				} else {
					templateType = "canvas"
				}
			} else {
				templateType = "canvas" // fallback
			}
		}

		// Copy README template
		readmeTemplate := filepath.Join("templates", templateType, "README.md")
		if err := copyFile(readmeTemplate, readmePath); err != nil {
			fmt.Printf("Error creating %s: %v\n", readmePath, err)
		}

		// Copy GEMINI template
		geminiTemplate := filepath.Join("templates", templateType, "GEMINI.md")
		if err := copyFile(geminiTemplate, geminiPath); err != nil {
			fmt.Printf("Error creating %s: %v\n", geminiPath, err)
		}
	}
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
