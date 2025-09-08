package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/frquxl/go-atelier/pkg/engine"
	"github.com/frquxl/go-atelier/pkg/gitutil"
	"github.com/frquxl/go-atelier/pkg/util"
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

var canvasDeleteCmd = &cobra.Command{
	Use:   "delete <canvas-full-name>",
	Short: "Delete a canvas",
	Long:  `Deletes a canvas and removes it from Git tracking. Requires the full directory name (e.g., canvas-sunflowers).`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// Check if we're in an artist directory
		if _, err := os.Stat(".artist"); os.IsNotExist(err) {
			listAvailableArtists()
			return fmt.Errorf("not in an artist directory. See available artists above")
		}

		canvasFullName := args[0]

		// Get current working directory (artist path)
		artistPath, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current working directory: %w", err)
		}

		canvasPath := filepath.Join(artistPath, canvasFullName)

		// Check for uncommitted changes in the canvas
		hasUncommitted, err := gitutil.IsPathDirty(artistPath, canvasFullName)
		if err != nil {
			return fmt.Errorf("failed to check for uncommitted changes: %w", err)
		}

		// Check for unpushed changes in the canvas
		hasUnpushed, err := gitutil.HasUnpushedCommits(canvasPath)
		if err != nil {
			return fmt.Errorf("failed to check for unpushed changes: %w", err)
		}

		// First confirmation prompt
		confirmMessage := fmt.Sprintf("Are you sure you want to delete canvas '%s'? This will delete the canvas's directory and all its contents, and remove it from Git tracking.", canvasFullName)
		if !util.Confirm(confirmMessage) {
			fmt.Println("Canvas deletion cancelled.")
			return nil
		}

		// If there are uncommitted or unpushed changes, warn and require second confirmation
		if hasUncommitted || hasUnpushed {
			warningMsg := "WARNING: This canvas has "
			warnings := []string{}
			if hasUncommitted {
				warnings = append(warnings, "uncommitted changes")
			}
			if hasUnpushed {
				warnings = append(warnings, "unpushed commits")
			}
			warningMsg += strings.Join(warnings, " and ")
			warningMsg += ". Deleting will permanently lose these changes."

			fmt.Println(warningMsg)

			// Second confirmation
			confirmMessage2 := fmt.Sprintf("Are you absolutely sure you want to delete canvas '%s' despite the %s?", canvasFullName, strings.Join(warnings, " and "))
			if !util.Confirm(confirmMessage2) {
				fmt.Println("Canvas deletion cancelled.")
				return nil
			}
		}

		if err = engine.DeleteCanvas(artistPath, canvasFullName); err != nil {
			return err
		}

		// The engine.DeleteCanvas function now prints the guidance message.
		return nil
	},
}

var canvasMoveCmd = &cobra.Command{
	Use:   "move <canvas-full-name> <new-artist-full-name>",
	Short: "Move a canvas from one artist to another.",
	Long:  `Moves a canvas from its current artist to another, updating Git submodules and internal paths accordingly.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		canvasFullName := args[0]
		newArtistFullName := args[1]

		if err := engine.MoveCanvas(canvasFullName, newArtistFullName); err != nil {
			return err
		}

		fmt.Printf("Canvas '%s' moved to artist '%s' successfully!\n", canvasFullName, newArtistFullName)
		return nil
	},
}

var canvasPushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes using the git push engine",
	Long:  `Push changes at the canvas level.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if we're in a canvas directory
		if _, err := os.Stat(".canvas"); os.IsNotExist(err) {
			return fmt.Errorf("not in a canvas directory")
		}

		// Get the directory of the executable
		execPath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("could not get executable path: %w", err)
		}
		execDir := filepath.Dir(execPath)
		scriptPath := filepath.Join(execDir, "pkg/push-engine/push-engine.sh")

		// Build command arguments
		execArgs := []string{scriptPath}
		if dryRun, _ := cmd.Flags().GetBool("dry-run"); dryRun {
			execArgs = append(execArgs, "--dry-run")
		}
		if quiet, _ := cmd.Flags().GetBool("quiet"); quiet {
			execArgs = append(execArgs, "--quiet")
		}
		if force, _ := cmd.Flags().GetBool("force"); force {
			execArgs = append(execArgs, "--force")
		}

		// Execute the push engine
		command := exec.Command("bash", execArgs...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Env = append(os.Environ(), "ENGINE_ASSUME_YES=true", "AUTO_COMMIT_DEFAULT=true")

		return command.Run()
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
	canvasPushCmd.Flags().Bool("dry-run", false, "Show what would be pushed without pushing")
	canvasPushCmd.Flags().Bool("quiet", false, "Suppress verbose output")
	canvasPushCmd.Flags().Bool("force", false, "Force push (use with caution)")
	RootCmd.AddCommand(canvasCmd)
	canvasCmd.AddCommand(canvasInitCmd)
	canvasCmd.AddCommand(canvasDeleteCmd)
	canvasCmd.AddCommand(canvasPushCmd)
	canvasCmd.AddCommand(canvasMoveCmd)
}