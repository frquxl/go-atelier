package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vincent",
	Short: "A small CLI to remember Vincent van Gogh.",
	Long:  `A small CLI to remember Vincent van Gogh, through his letters to his brother Theo.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
