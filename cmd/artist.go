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
		var canvasName string
		if withCanvas, _ := cmd.Flags().GetBool("with-canvas"); withCanvas {
			canvasName = "example" // Artists are created with a default example canvas when flag is set
		} else {
			canvasName = "" // No canvas created by default
		}

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

var artistDeleteCmd = &cobra.Command{
	Use:   "delete <artist-full-name>",
	Short: "Delete an artist studio",
	Long:  `Deletes an artist studio and removes it from Git tracking. Requires the full directory name (e.g., artist-van-gogh).`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// Check if we're in an atelier directory
		if _, err := os.Stat(".atelier"); os.IsNotExist(err) {
			listAvailableAteliers()
			return fmt.Errorf("not in an atelier directory. See available ateliers above")
		}

		artistFullName := args[0]

		// Get current working directory (atelier path)
		atelierPath, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get current working directory: %w", err)
		}

		artistPath := filepath.Join(atelierPath, artistFullName)

		// Check for uncommitted changes in the artist
		hasUncommitted, err := gitutil.IsPathDirty(atelierPath, artistFullName)
		if err != nil {
			return fmt.Errorf("failed to check for uncommitted changes: %w", err)
		}

		// Check for unpushed changes in the artist
		hasUnpushed, err := gitutil.HasUnpushedCommits(artistPath)
		if err != nil {
			return fmt.Errorf("failed to check for unpushed changes: %w", err)
		}

		// First confirmation prompt
		confirmMessage := fmt.Sprintf("Are you sure you want to delete artist '%s'? This will delete the artist's directory and all its contents, and remove it from Git tracking.", artistFullName)
		if !util.Confirm(confirmMessage) {
			fmt.Println("Artist deletion cancelled.")
			return nil
		}

		// If there are uncommitted or unpushed changes, warn and require second confirmation
		if hasUncommitted || hasUnpushed {
			warningMsg := "WARNING: This artist has "
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

			// Show detailed breakdown of canvas statuses
			fmt.Println("\nCanvas status breakdown:")
			showCanvasStatuses(atelierPath, artistFullName)

			// Second confirmation
			confirmMessage2 := fmt.Sprintf("Are you absolutely sure you want to delete artist '%s' despite the %s?", artistFullName, strings.Join(warnings, " and "))
			if !util.Confirm(confirmMessage2) {
				fmt.Println("Artist deletion cancelled.")
				return nil
			}
		}

		if err = engine.DeleteArtist(atelierPath, artistFullName); err != nil {
			return err
		}

		// The engine.DeleteArtist function now prints the guidance message.
		return nil
	},
}

var artistPushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes using the git push engine",
	Long:  `Push changes at the artist level, recursing into all canvases.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if we're in an artist directory
		if _, err := os.Stat(".artist"); os.IsNotExist(err) {
			return fmt.Errorf("not in an artist directory")
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

func showCanvasStatuses(atelierPath, artistFullName string) {
	artistPath := filepath.Join(atelierPath, artistFullName)

	entries, err := os.ReadDir(artistPath)
	if err != nil {
		fmt.Printf("  Error reading artist directory: %v\n", err)
		return
	}

	hasCanvases := false
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "canvas-") {
			hasCanvases = true
			canvasName := entry.Name()
			canvasPath := filepath.Join(artistPath, canvasName)

			// Check for uncommitted changes
			hasUncommitted, err := gitutil.IsPathDirty(artistPath, canvasName)
			if err != nil {
				fmt.Printf("  %s: error checking status - %v\n", canvasName, err)
				continue
			}

			// Check for unpushed changes
			hasUnpushed, err := gitutil.HasUnpushedCommits(canvasPath)
			if err != nil {
				fmt.Printf("  %s: error checking status - %v\n", canvasName, err)
				continue
			}

			// Build status message
			statusParts := []string{}
			if hasUncommitted {
				statusParts = append(statusParts, "uncommitted changes")
			}
			if hasUnpushed {
				statusParts = append(statusParts, "unpushed commits")
			}

			if len(statusParts) > 0 {
				statusMsg := strings.Join(statusParts, " and ")
				fmt.Printf("  %s: %s\n", canvasName, statusMsg)
			} else {
				fmt.Printf("  %s: clean\n", canvasName)
			}
		}
	}

	if !hasCanvases {
		fmt.Println("  No canvases found in this artist.")
	}
}

func init() {
	artistPushCmd.Flags().Bool("dry-run", false, "Show what would be pushed without pushing")
	artistPushCmd.Flags().Bool("quiet", false, "Suppress verbose output")
	artistPushCmd.Flags().Bool("force", false, "Force push (use with caution)")
	artistInitCmd.Flags().Bool("with-canvas", false, "Create a default example canvas with the artist")
	RootCmd.AddCommand(artistCmd)
	artistCmd.AddCommand(artistInitCmd)
	artistCmd.AddCommand(artistDeleteCmd)
	artistCmd.AddCommand(artistPushCmd)
}
