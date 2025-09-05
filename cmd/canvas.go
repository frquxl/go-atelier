package cmd

import (
	"atelier-cli/pkg/engine"
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
	Long:  `Initialize a new canvas within the current artist workspace as a Git submodule. Must be run from an artist directory.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// Check if we're in an artist directory
		if _, err := os.Stat(".artist"); os.IsNotExist(err) {
			listAvailableArtists()
			return fmt.Errorf("not in an artist directory. See available artists above")
		}

		canvasName := args[0]

		// Get current working directory to construct absolute paths
		artistPath, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current working directory: %w", err)
		}

		if err = engine.CreateCanvas(artistPath, canvasName); err != nil {
			return err // Error is already formatted and cleanup is handled by the engine
		}

		fmt.Printf("Canvas '%s' initialized successfully in artist '%s'!\n", canvasName, filepath.Base(artistPath))
		return nil
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