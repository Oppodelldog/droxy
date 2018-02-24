package cmd

import (
	"fmt"

	"github.com/Oppodelldog/droxy/version"
	"github.com/spf13/cobra"
)

// Root is the root cobra command.
// Basically it contains information to display the help message, but has no logic for execution.
// The logical part for execution is handled in main.
var Root = &cobra.Command{
	Use:   "droxy",
	Short: "docker proxy commands by configuration",
	Long: fmt.Sprintf(`     _                             
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
