package templates

import (
	"atelier-cli/pkg/fs"
	"embed"
	"fmt"
	"path/filepath"
)

//go:embed assets/*
var TemplatesFS embed.FS

// CreateBoilerplate generates standard project files from embedded templates.
func CreateBoilerplate(basePath, projectType string) error {
	templateDir := fmt.Sprintf("assets/%s", projectType)

	// Define files to be copied and potentially renamed
	files := map[string]string{
		"README.md":    "README.md",
		"GEMINI.md":    "GEMINI.md",
		"Makefile":     "Makefile",
		"gitignore":    ".gitignore",
		"geminiignore": ".geminiignore",
	}

	for src, dest := range files {
		path := fmt.Sprintf("%s/%s", templateDir, src)
		content, err := TemplatesFS.ReadFile(path)
		if err != nil {
			// Some project types might not have all files (e.g., no Makefile), so we can ignore not found errors
			// if os.IsNotExist(err) {
			// 	continue
			// }
			return fmt.Errorf("failed to read embedded template %s: %w", path, err)
		}

		if err := fs.WriteFile(filepath.Join(basePath, dest), content); err != nil {
			return fmt.Errorf("failed to write file %s: %w", dest, err)
		}
	}

	return nil
}
