package commands

type actionChain []actionChainElement

func (c actionChain) execute(args []string) int {
	for _, action := range c {
		if action.IsResponsible(args) {
			return action.Execute()
		}
	}

	panic("dead end of chain")
}

type actionChainElement interface {
	IsResponsible(args []string) bool
	Execute() int
}
