package cmd

import (
	"fmt"
	"os"
	"os/exec"
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
	Run: func(cmd *cobra.Command, args []string) {
		artist := args[0]

		// Check if we're in an atelier directory
		if _, err := os.Stat(".atelier"); os.IsNotExist(err) {
			fmt.Println("Error: Not in an atelier directory.")
			listAvailableAteliers()
			return
		}

		// Create artist directory as Git repository
		artistDir := "artist-" + artist

		// Initialize artist as Git repository
		if err := exec.Command("git", "init", artistDir).Run(); err != nil {
			fmt.Printf("Error initializing artist Git repository: %v\n", err)
			return
		}

		// Change to artist directory to set up initial commit
		if err := os.Chdir(artistDir); err != nil {
			fmt.Printf("Error changing to artist directory: %v\n", err)
			return
		}

		// Create artist marker file
		if err := os.WriteFile(".artist", []byte(artist), 0644); err != nil {
			fmt.Printf("Error creating artist marker file: %v\n", err)
			return
		}

		// Create initial artist boilerplate
		createBoilerplateFiles(".")

		// Commit artist setup
		if err := exec.Command("git", "add", ".").Run(); err != nil {
			fmt.Printf("Error staging artist files: %v\n", err)
			return
		}

		if err := exec.Command("git", "commit", "-m", fmt.Sprintf("feat: initialize artist %s", artist)).Run(); err != nil {
			fmt.Printf("Error committing artist setup: %v\n", err)
			return
		}

		// Go back to atelier directory to add submodule
		os.Chdir("..")

		// Add artist as submodule to atelier
		if err := exec.Command("git", "submodule", "add", "./"+artistDir, artistDir).Run(); err != nil {
			fmt.Printf("Error adding artist as submodule: %v\n", err)
			return
		}

		// Change to artist directory to set up canvas
		originalDir, _ := os.Getwd()
		if err := os.Chdir(artistDir); err != nil {
			fmt.Printf("Error changing to artist directory for canvas setup: %v\n", err)
			return
		}

		// Create default canvas within artist
		canvasDir := "canvas-example"
		if err := exec.Command("git", "init", canvasDir).Run(); err != nil {
			fmt.Printf("Error initializing canvas Git repository: %v\n", err)
			return
		}

		// Change to canvas directory to set up initial commit
		if err := os.Chdir(canvasDir); err != nil {
			fmt.Printf("Error changing to canvas directory: %v\n", err)
			return
		}

		// Create marker files
		markerFiles := map[string]string{
			"../.artist": artist,
			".canvas":    "example",
		}

		for markerPath, content := range markerFiles {
			if err := os.WriteFile(markerPath, []byte(content), 0644); err != nil {
				fmt.Printf("Error creating marker file %s: %v\n", markerPath, err)
				return
			}
		}

		// Create boilerplate files for canvas
		createBoilerplateFiles("..", ".")

		// Commit canvas setup
		if err := exec.Command("git", "add", ".").Run(); err != nil {
			fmt.Printf("Error staging canvas files: %v\n", err)
			return
		}

		if err := exec.Command("git", "commit", "-m", "feat: initialize canvas example").Run(); err != nil {
			fmt.Printf("Error committing canvas setup: %v\n", err)
			return
		}

		// Go back to artist directory
		os.Chdir("..")

		// Add canvas as submodule to artist
		if err := exec.Command("git", "submodule", "add", "./"+canvasDir, canvasDir).Run(); err != nil {
			fmt.Printf("Error adding canvas as submodule: %v\n", err)
			return
		}

		// Commit artist setup with submodule if there are changes
		if err := exec.Command("git", "add", ".").Run(); err != nil {
			fmt.Printf("Error staging artist files: %v\n", err)
			return
		}

		// Only commit if there are changes to commit
		if err := exec.Command("git", "diff", "--cached", "--quiet").Run(); err != nil {
			// There are changes to commit
			commitMsg := fmt.Sprintf("feat: initialize artist %s with default canvas", artist)
			if err := exec.Command("git", "commit", "-m", commitMsg).Run(); err != nil {
				fmt.Printf("Error committing artist setup: %v\n", err)
				return
			}
		}

		// Go back to atelier directory and commit submodule addition
		os.Chdir(originalDir)
		if err := exec.Command("git", "add", artistDir).Run(); err != nil {
			fmt.Printf("Error staging submodule: %v\n", err)
			return
		}

		// Only commit if there are changes to commit
		if err := exec.Command("git", "diff", "--cached", "--quiet").Run(); err != nil {
			// There are changes to commit
			if err := exec.Command("git", "commit", "-m", fmt.Sprintf("feat: add artist %s as submodule", artist)).Run(); err != nil {
				fmt.Printf("Error committing submodule addition: %v\n", err)
				return
			}
		}

		fmt.Printf("Artist '%s' initialized as submodule with default canvas\n", artist)
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
