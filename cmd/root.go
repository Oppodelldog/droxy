package cmd

import "github.com/spf13/cobra"

var Root = &cobra.Command{
	Use:   "docker-command-proxy",
	Short: "tool to create and run custom-command-containers",
	Long:  `yeah`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
