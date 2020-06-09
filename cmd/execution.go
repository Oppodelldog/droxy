package cmd

//Run droxy command, sub-command or help dialog.
func Run(args []string) int {
	return getActionChain().execute(args)
}
