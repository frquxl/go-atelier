package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

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
	Run: func(cmd *cobra.Command, args []string) {
		canvas := args[0]

		// Check if we're in an artist directory
		if _, err := os.Stat(".artist"); os.IsNotExist(err) {
			fmt.Println("Error: Not in an artist directory.")
			listAvailableArtists()
			return
		}

		// Create canvas directory as Git repository
		canvasDir := "canvas-" + canvas

		// Initialize canvas as Git repository
		if err := exec.Command("git", "init", canvasDir).Run(); err != nil {
			fmt.Printf("Error initializing canvas Git repository: %v\n", err)
			return
		}

		// Change to canvas directory to set up files
		originalDir, _ := os.Getwd()
		defer os.Chdir(originalDir)

		if err := os.Chdir(canvasDir); err != nil {
			fmt.Printf("Error changing to canvas directory: %v\n", err)
			return
		}

		// Create marker file
		if err := os.WriteFile(".canvas", []byte(canvas), 0644); err != nil {
			fmt.Printf("Error creating marker file: %v\n", err)
			return
		}

		// Create boilerplate files
		createBoilerplateFiles(".")

		// Commit canvas setup
		if err := exec.Command("git", "add", ".").Run(); err != nil {
			fmt.Printf("Error staging canvas files: %v\n", err)
			return
		}

		commitMsg := fmt.Sprintf("feat: initialize canvas %s", canvas)
		if err := exec.Command("git", "commit", "-m", commitMsg).Run(); err != nil {
			fmt.Printf("Error committing canvas setup: %v\n", err)
			return
		}

		// Go back to artist directory
		os.Chdir(originalDir)

		// Add canvas as submodule to artist
		if err := exec.Command("git", "submodule", "add", "./"+canvasDir, canvasDir).Run(); err != nil {
			fmt.Printf("Error adding canvas as submodule: %v\n", err)
			return
		}
		if err := exec.Command("git", "add", canvasDir).Run(); err != nil {
			fmt.Printf("Error staging submodule: %v\n", err)
			return
		}

		if err := exec.Command("git", "commit", "-m", fmt.Sprintf("feat: add canvas %s as submodule", canvas)).Run(); err != nil {
			fmt.Printf("Error committing submodule addition: %v\n", err)
			return
		}

		fmt.Printf("Canvas '%s' initialized as submodule\n", canvas)
	},
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
