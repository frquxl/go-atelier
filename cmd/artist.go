package cmd

import (
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
	Long:  `Initialize a new artist studio within the existing atelier.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		artist := args[0]

		// Check if any atelier-* directory exists
		atelierDir := ""
		entries, err := os.ReadDir(".")
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			return
		}

		for _, entry := range entries {
			if entry.IsDir() && strings.HasPrefix(entry.Name(), "atelier-") {
				atelierDir = entry.Name()
				break
			}
		}

		if atelierDir == "" {
			fmt.Println("Error: No atelier directory found. Run 'atelier init <name>' first.")
			return
		}

		// Create artist directory
		artistDir := filepath.Join(atelierDir, artist)
		canvasDir := filepath.Join(artistDir, "canvas")

		if err := os.MkdirAll(canvasDir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", canvasDir, err)
			return
		}

		// Create boilerplate files
		createBoilerplateFiles(artistDir, canvasDir)

		fmt.Printf("Artist '%s' initialized with canvas\n", artist)
	},
}

func init() {
	RootCmd.AddCommand(artistCmd)
	artistCmd.AddCommand(artistInitCmd)
}
