package cmd

import (
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "atelier",
	Short: "A metaphor-driven CLI for software project management",
	Long:  `Atelier is a CLI tool that uses the atelier/artist/canvas metaphor to help manage software projects.`,
}

func init() {
	RootCmd.Version = getVersion()
}

func getVersion() string {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.Output()
	if err != nil {
		return "dev"
	}
	return strings.TrimSpace(string(output))
}
