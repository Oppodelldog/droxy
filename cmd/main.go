package main

import (
	"os"

	"github.com/Oppodelldog/droxy/cmd/commands"
)

func main() {
	os.Exit(commands.Run(os.Args))
}
