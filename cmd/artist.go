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

		// Check if we're in an atelier directory
		if _, err := os.Stat(".atelier"); os.IsNotExist(err) {
			fmt.Println("Error: Not in an atelier directory. Run 'atelier init <name>' first, then cd into the created directory.")
			return
		}

		// Create artist directory with marker
		artistDir := "artist-" + artist
		canvasDir := "canvas-" + artist // Default canvas with same name as artist

		fullCanvasDir := filepath.Join(artistDir, canvasDir)

		if err := os.MkdirAll(fullCanvasDir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", fullCanvasDir, err)
			return
		}

		// Create marker files
		markerFiles := map[string]string{
			filepath.Join(artistDir, ".artist"):     artist,
			filepath.Join(fullCanvasDir, ".canvas"): artist,
		}

		for markerPath, content := range markerFiles {
			if err := os.WriteFile(markerPath, []byte(content), 0644); err != nil {
				fmt.Printf("Error creating marker file %s: %v\n", markerPath, err)
				return
			}
		}

		// Create boilerplate files
		createBoilerplateFiles(artistDir, fullCanvasDir)

		fmt.Printf("Artist '%s' initialized with default canvas\n", artist)
	},
}

func init() {
	RootCmd.AddCommand(artistCmd)
	artistCmd.AddCommand(artistInitCmd)
}
