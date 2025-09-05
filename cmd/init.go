package cmd

import (
	"atelier-cli/pkg/engine"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <atelier-name> [<artist-name> <canvas-name>]",
	Short: "Initialize a new atelier workspace",
	Long: `Initialize a new atelier workspace with 3-level Git submodule structure.
Creates atelier-<atelier-name> as main repo, artist as submodule, canvas as submodule of artist.
If no artist/canvas provided, defaults to 'van-gogh' and 'sunflowers'.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) == 0 {
			return fmt.Errorf("atelier name is required")
		}

		atelierBaseName := args[0]
		artistName := "van-gogh"
		canvasName := "sunflowers"

		if len(args) >= 3 {
			artistName = args[1]
			canvasName = args[2]
		}

		// Get current working directory to construct absolute paths
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current working directory: %w", err)
		}

		// 1. Create the Atelier
		atelierPath, err := engine.CreateAtelier(wd, atelierBaseName)
		if err != nil {
			return err // Error is already formatted and cleanup is handled by the engine
		}

		// 2. Create the Artist and default Canvas
		if err = engine.CreateArtist(atelierPath, artistName, canvasName); err != nil {
			return err // Error is already formatted and cleanup is handled by the engine
		}

		fmt.Printf("Atelier '%s' initialized successfully!\n", atelierBaseName)
		fmt.Printf("  - Path: %s\n", atelierPath)
		fmt.Printf("  - Contains artist: %s\n", artistName)
		fmt.Printf("  - Which contains canvas: %s\n", canvasName)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}