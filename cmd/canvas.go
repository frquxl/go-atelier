package cmd

import (
	"fmt"
	"os"
	"path/filepath"
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
	Long:  `Initialize a new canvas within the current artist workspace. Must be run from an artist directory.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		canvas := args[0]

		// Check if we're in an artist directory
		if _, err := os.Stat(".artist"); os.IsNotExist(err) {
			fmt.Println("Error: Not in an artist directory.")
			listAvailableArtists()
			return
		}

		// Create canvas directory with marker
		canvasDir := "canvas-" + canvas

		if err := os.MkdirAll(canvasDir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", canvasDir, err)
			return
		}

		// Create marker file
		markerPath := filepath.Join(canvasDir, ".canvas")
		if err := os.WriteFile(markerPath, []byte(canvas), 0644); err != nil {
			fmt.Printf("Error creating marker file %s: %v\n", markerPath, err)
			return
		}

		// Create boilerplate files
		createBoilerplateFiles(canvasDir)

		fmt.Printf("Canvas '%s' initialized\n", canvas)
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
