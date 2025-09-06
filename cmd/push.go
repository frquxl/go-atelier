package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes using the git push engine",
	Long:  `Push changes at the atelier level, recursing into all artists and canvases.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if we're in an atelier directory
		if _, err := os.Stat(".atelier"); os.IsNotExist(err) {
			return fmt.Errorf("not in an atelier directory")
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

func init() {
	pushCmd.Flags().Bool("dry-run", false, "Show what would be pushed without pushing")
	pushCmd.Flags().Bool("quiet", false, "Suppress verbose output")
	pushCmd.Flags().Bool("force", false, "Force push (use with caution)")
	RootCmd.AddCommand(pushCmd)
}
