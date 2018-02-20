package cmd

import "github.com/spf13/cobra"

// Root is the root cobra command.
// Basically it contains information to display the help message, but has no logic for execution.
// The logical part for execution is handled in main.
var Root = &cobra.Command{
	Use:   "droxy",
	Short: "tool to create and run custom-command-containers",
	Long:  `yeah`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
