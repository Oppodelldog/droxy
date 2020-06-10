package commands

import (
	"fmt"

	"github.com/Oppodelldog/droxy/version"
	"github.com/spf13/cobra"
)

// newRoot returns a new cobra command, which contains help display and sub-commands.
func newRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "droxy",
		Short: "docker proxy commands by configuration",
		Long: fmt.Sprintf(`
     _                             
    | |                            
  __| |  ____   ___   _   _  _   _ 
 / _  | / ___) / _ \ ( \ / )| | | |
( (_| || |    | |_| | ) X ( | |_| |
 \____||_|     \___/ (_/ \_) \__  |
                            (____/ 
Version: %s
About  : droxy creates commands that proxy to docker`, version.Number),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.AddCommand(newSymlinkCommandWrapper())
	rootCmd.AddCommand(newHardlinkCommandWrapper())
	rootCmd.AddCommand(newCloneCommandWrapper())

	return rootCmd
}
