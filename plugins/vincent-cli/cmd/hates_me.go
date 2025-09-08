package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hatesMeCmd = &cobra.Command{
	Use:   "hates-me",
	Short: "A terrible secret.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Theo shot Vincent")
	},
}

func init() {
	theoCmd.AddCommand(hatesMeCmd)
}