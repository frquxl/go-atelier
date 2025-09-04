package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "atelier",
	Short: "A metaphor-driven CLI for software project management",
	Long:  `Atelier is a CLI tool that uses the atelier/artist/canvas metaphor to help manage software projects.`,
}
