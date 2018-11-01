package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Oppodelldog/droxy/proxyfile"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type action struct {
	isResponsibleFunc func(args []string) bool
	executeFunc       func() int
}

func (a *action) IsResponsible(args []string) bool {
	return a.isResponsibleFunc(args)
}

func (a *action) Execute() int {
	return a.executeFunc()
}

func getActionChain() actionChain {
	return []actionChainElement{
		newSubCommandAction(newRoot()),
		newHelpDisplayAction(),
		newRevealItsDroxyAction(),
		newDroxyCommandAction(),
	}
}

func newSubCommandAction(cmd executer) actionChainElement {
	return &action{
		isResponsibleFunc: func(args []string) bool { return shallExecuteSubCommand(args, newRoot()) },
		executeFunc:       func() int { return execSubCommand(cmd) },
	}
}

func newDroxyCommandAction() actionChainElement {
	return &action{
		isResponsibleFunc: func([]string) bool { return true },
		executeFunc:       func() int { return executeDroxyCommand(os.Args) },
	}
}

func newRevealItsDroxyAction() actionChainElement {
	return &action{
		isResponsibleFunc: shallRevealItsDroxy,
		executeFunc:       revealTheTruth,
	}
}

func newHelpDisplayAction() actionChainElement {
	return &action{
		isResponsibleFunc: shallDisplayHelp,
		executeFunc:       func() int { return displayHelp(newRoot()) },
	}
}

func shallExecuteSubCommand(args []string, rootCmd *cobra.Command) bool {
	return len(args) >= 2 && isSubCommand(args[1], rootCmd.Commands())
}

func shallDisplayHelp(args []string) bool {
	return len(args) >= 1 && filepath.Base(args[0]) == proxyfile.GetCommandName()
}

func shallRevealItsDroxy(args []string) bool {
	for _, arg := range args {
		if arg == "--is-it-droxy" {

			return true
		}
	}

	return false
}

const theTruth = "YES-IT-IS"

func revealTheTruth() int {
	fmt.Println(theTruth)
	return 0
}

type executer interface {
	Execute() error
}

func execSubCommand(cmd executer) int {
	err := cmd.Execute()
	if err != nil {
		logrus.Info(err)
	}

	return 0
}

type helper interface {
	Help() error
}

func displayHelp(cmd helper) int {
	err := cmd.Help()
	if err != nil {
		logrus.Info(err)
	}

	return 0
}

func isSubCommand(s string, commands []*cobra.Command) bool {
	var subCommandNames []string
	for _, subCommand := range commands {
		subCommandNames = append(subCommandNames, subCommand.Name())

	}

	return stringInSlice(s, subCommandNames)
}

func stringInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}

	return false
}
