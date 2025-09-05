package cmd

import (
	"atelier-cli/pkg/engine"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

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
		if _, err := os.Stat(".atelier"); os.IsNotExist(err) {
			listAvailableAteliers()
			return fmt.Errorf("not in an atelier directory. See available ateliers above")
		}

		artistName := args[0]
		canvasName := "example" // Artists are created with a default example canvas

		// Get current working directory to construct absolute paths
		atelierPath, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current working directory: %w", err)
		}

		if err = engine.CreateArtist(atelierPath, artistName, canvasName); err != nil {
			return err // Error is already formatted and cleanup is handled by the engine
		}

		fmt.Printf("Artist '%s' initialized successfully in atelier '%s'!\n", artistName, filepath.Base(atelierPath))
		return nil
	},
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