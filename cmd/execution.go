package cmd

//Run droxy command, subcommand or help dialog
func Run(args []string) int {
	return getActionChain().execute(args)
}
