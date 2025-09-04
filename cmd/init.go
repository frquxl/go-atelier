package cmd

import (
	"embed"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var templatesFS embed.FS

var initCmd = &cobra.Command{
	Use:   "init <atelier-name> [<artist-name> <canvas-name>]",
	Short: "Initialize a new atelier workspace",
	Long: `Initialize a new atelier workspace with 3-level Git submodule structure.
Creates atelier-<atelier-name> as main repo, artist as submodule, canvas as submodule of artist.
If no artist/canvas provided, defaults to 'van-gogh' and 'sunflowers'.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: atelier name is required")
			return
		}

		atelierBaseName := args[0]
		artist := "van-gogh"
		canvas := "sunflowers"

		if len(args) >= 3 {
			artist = args[1]
			canvas = args[2]
		}

		// Always create atelier-XXXXX format
		atelierDir := "atelier-" + atelierBaseName
		artistDir := "artist-" + artist
		canvasDir := "canvas-" + canvas

		// Create atelier directory and initialize as Git repo
		if err := os.MkdirAll(atelierDir, 0755); err != nil {
			fmt.Printf("Error creating atelier directory: %v\n", err)
			return
		}

		// Initialize atelier as main Git repository
		if err := exec.Command("git", "init", atelierDir).Run(); err != nil {
			fmt.Printf("Error initializing atelier Git repository: %v\n", err)
			return
		}

		// Change to atelier directory
		originalDir, _ := os.Getwd()
		defer os.Chdir(originalDir)

		if err := os.Chdir(atelierDir); err != nil {
			fmt.Printf("Error changing to atelier directory: %v\n", err)
			return
		}

		// Create atelier marker file
		if err := os.WriteFile(".atelier", []byte(atelierBaseName), 0644); err != nil {
			fmt.Printf("Error creating atelier marker file: %v\n", err)
			return
		}

		// Create artist as Git repository
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

		// Go back to atelier directory
		os.Chdir("..")

		// Add artist as submodule to atelier
		if err := exec.Command("git", "submodule", "add", "./"+artistDir, artistDir).Run(); err != nil {
			fmt.Printf("Error adding artist as submodule: %v\n", err)
			return
		}

		// Change to artist directory to set up canvas
		if err := os.Chdir(artistDir); err != nil {
			fmt.Printf("Error changing to artist directory: %v\n", err)
			return
		}

		// Create canvas as Git repository
		if err := exec.Command("git", "init", canvasDir).Run(); err != nil {
			fmt.Printf("Error initializing canvas Git repository: %v\n", err)
			return
		}

		// Change to canvas directory to set up initial commit
		if err := os.Chdir(canvasDir); err != nil {
			fmt.Printf("Error changing to canvas directory: %v\n", err)
			return
		}

		// Create canvas marker file
		if err := os.WriteFile(".canvas", []byte(canvas), 0644); err != nil {
			fmt.Printf("Error creating canvas marker file: %v\n", err)
			return
		}

		// Create initial canvas boilerplate
		createBoilerplateFiles(".")

		// Commit canvas setup
		if err := exec.Command("git", "add", ".").Run(); err != nil {
			fmt.Printf("Error staging canvas files: %v\n", err)
			return
		}

		if err := exec.Command("git", "commit", "-m", fmt.Sprintf("feat: initialize canvas %s", canvas)).Run(); err != nil {
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

		// Go back to atelier directory to create boilerplate files
		os.Chdir("..")
		// Create boilerplate files for atelier level
		createBoilerplateFiles(".")
		// Go back to artist directory for final commits
		os.Chdir(artistDir)

		// Commit canvas submodule addition to artist
		if err := exec.Command("git", "add", canvasDir).Run(); err != nil {
			fmt.Printf("Error staging canvas submodule: %v\n", err)
			return
		}

		if err := exec.Command("git", "commit", "-m", fmt.Sprintf("feat: add canvas %s as submodule", canvas)).Run(); err != nil {
			fmt.Printf("Error committing canvas submodule: %v\n", err)
			return
		}

		// Go back to atelier directory and commit artist submodule
		os.Chdir("..")
		if err := exec.Command("git", "add", artistDir).Run(); err != nil {
			fmt.Printf("Error staging artist submodule: %v\n", err)
			return
		}

		if err := exec.Command("git", "commit", "-m", fmt.Sprintf("feat: add artist %s as submodule", artist)).Run(); err != nil {
			fmt.Printf("Error committing artist submodule: %v\n", err)
			return
		}

		fmt.Printf("Atelier '%s' initialized with 3-level Git submodule structure\n", atelierBaseName)
		fmt.Printf("├── Artist '%s' (submodule)\n", artist)
		fmt.Printf("│   └── Canvas '%s' (submodule)\n", canvas)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}

func createBoilerplateFiles(dirs ...string) {
	for _, dir := range dirs {
		// Determine template type based on directory name
		templateType := getTemplateType(dir)

		// Create README.md from embedded template
		readmeDest := filepath.Join(dir, "README.md")
		readmeContent := readEmbeddedTemplate(templateType, "README.md")
		if readmeContent != "" {
			os.WriteFile(readmeDest, []byte(readmeContent), 0644)
		} else {
			// Fallback to inline generation
			readmeContent := getDefaultContent(templateType, "README.md", dir)
			os.WriteFile(readmeDest, []byte(readmeContent), 0644)
		}

		// Create GEMINI.md from embedded template
		geminiDest := filepath.Join(dir, "GEMINI.md")
		geminiContent := readEmbeddedTemplate(templateType, "GEMINI.md")
		if geminiContent != "" {
			os.WriteFile(geminiDest, []byte(geminiContent), 0644)
		} else {
			// Fallback to inline generation
			geminiContent := getDefaultContent(templateType, "GEMINI.md", dir)
			os.WriteFile(geminiDest, []byte(geminiContent), 0644)
		}
	}
}

func getTemplateType(dir string) string {
	// Normalize relative paths like "." or ".." to absolute, so Base resolves correctly
	if abs, err := filepath.Abs(dir); err == nil {
		dir = abs
	}
	baseName := filepath.Base(dir)

	// Check for exact prefix matches in priority order
	// Canvas first (most specific)
	if strings.HasPrefix(baseName, "canvas-") {
		return "canvas"
	}
	// Artist second
	if strings.HasPrefix(baseName, "artist-") {
		return "artist"
	}
	// Atelier last (least specific)
	if strings.HasPrefix(baseName, "atelier-") {
		return "atelier"
	}

	// For directories that don't have standard prefixes,
	// check if they contain atelier/artist/canvas keywords
	// Check in reverse order to avoid false positives
	if strings.Contains(baseName, "canvas") && !strings.Contains(baseName, "artist") && !strings.Contains(baseName, "atelier") {
		return "canvas"
	}
	if strings.Contains(baseName, "artist") && !strings.Contains(baseName, "atelier") {
		return "artist"
	}
	if strings.Contains(baseName, "atelier") {
		return "atelier"
	}

	return "atelier" // Default fallback
}

func readEmbeddedTemplate(templateType, filename string) string {
	templatePath := fmt.Sprintf("templates/%s/%s", templateType, filename)
	content, err := templatesFS.ReadFile(templatePath)
	if err != nil {
		return ""
	}
	return string(content)
}

func findTemplatePath(templateType, filename string) string {
	// Try relative to current working directory first (for local development)
	templatePath := filepath.Join("templates", templateType, filename)
	if _, err := os.Stat(templatePath); err == nil {
		return templatePath
	}

	// Get the executable path to find templates relative to the binary
	execPath, err := os.Executable()
	if err != nil {
		return ""
	}

	// Get the directory containing the executable
	execDir := filepath.Dir(execPath)

	// Try relative to executable directory (for installed binary)
	templatePath = filepath.Join(execDir, "templates", templateType, filename)
	if _, err := os.Stat(templatePath); err == nil {
		return templatePath
	}

	// Try relative to executable directory with different paths
	// Handle case where binary is in a subdirectory
	parentDir := filepath.Dir(execDir)
	templatePath = filepath.Join(parentDir, "templates", templateType, filename)
	if _, err := os.Stat(templatePath); err == nil {
		return templatePath
	}

	// Try going up more levels from executable (for GOPATH installs)
	grandParentDir := filepath.Dir(parentDir)
	templatePath = filepath.Join(grandParentDir, "templates", templateType, filename)
	if _, err := os.Stat(templatePath); err == nil {
		return templatePath
	}

	// For development, try relative to the source directory
	// This handles the case where we're running from the source directory
	cwd, _ := os.Getwd()
	templatePath = filepath.Join(cwd, "templates", templateType, filename)
	if _, err := os.Stat(templatePath); err == nil {
		return templatePath
	}

	// Template not found
	return ""
}

func getDefaultContent(templateType, filename, dir string) string {
	dirName := filepath.Base(dir)

	switch filename {
	case "README.md":
		switch templateType {
		case "atelier":
			return fmt.Sprintf("# %s\n\nWelcome to your Atelier workspace!\n\nThis directory contains your software projects organized as:\n- Atelier: Main workspace (this level)\n- Artists: Project groups (Git submodules)\n- Canvases: Individual projects (Git submodules)\n\nEach canvas is an independent Git repository for isolated development.", dirName)
		case "artist":
			return fmt.Sprintf("# %s\n\nArtist workspace containing multiple project canvases.\n\nEach canvas is an independent Git repository. Work in any canvas directory to develop independently.\n\n## Canvases\nNavigate to any canvas-* directory to start working on a specific project.", dirName)
		case "canvas":
			return fmt.Sprintf("# %s\n\nProject canvas for independent development.\n\nThis is a complete Git repository where you can develop your software project.\n\n## Getting Started\n1. This directory has its own Git repository\n2. Work here independently of other projects\n3. Commit changes: git add . && git commit -m 'your message'", dirName)
		}
	case "GEMINI.md":
		switch templateType {
		case "atelier":
			return fmt.Sprintf("# AI Context: %s Atelier\n\nThis is the root level of an atelier workspace.\n\n## Architecture\n- Atelier: Main repository (this level)\n- Artists: Git submodules containing project groups\n- Canvases: Git submodules containing individual projects\n\n## Development\nEach canvas is an independent Git repository for isolated development.", dirName)
		case "artist":
			return fmt.Sprintf("# AI Context: %s Artist\n\nArtist workspace grouping related project canvases.\n\n## Structure\n- This artist contains multiple canvas submodules\n- Each canvas is an independent Git repository\n- Work in canvas directories for project development\n\n## Context\nUse this level to organize related projects thematically.", dirName)
		case "canvas":
			return fmt.Sprintf("# AI Context: %s Canvas\n\nIndividual project canvas with independent Git repository.\n\n## Development\n- This directory has its own Git history\n- Develop software independently here\n- Commit changes affect only this project\n\n## Architecture\nComplete, self-contained software project with own repository.", dirName)
		}
	}

	return fmt.Sprintf("# %s\n\nDefault content for %s.\n", dirName, filename)
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	// Ensure the destination file has the correct permissions
	return os.Chmod(dst, 0644)
}
