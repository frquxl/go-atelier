package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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

		// Check if atelier exists
		atelierDir := "atelier"
		if _, err := os.Stat(atelierDir); os.IsNotExist(err) {
			fmt.Println("Error: Atelier does not exist. Run 'atelier init' first.")
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
