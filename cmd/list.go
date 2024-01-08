package cmd

import (
	"github.com/spf13/cobra"
)

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "show all gpt client",
	Run:   runList,
}

func runList(_ *cobra.Command, args []string) {
	// TODO: show models list
}
