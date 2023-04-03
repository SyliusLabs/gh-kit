package cmd

import (
	"github.com/spf13/cobra"
)

type Command interface {
	GetCommand() *cobra.Command
}

func GetCommand(command Command) *cobra.Command {
	return command.GetCommand()
}

type RootCmd struct {
	cmd *cobra.Command
}

func (r *RootCmd) Execute() error {
	return r.cmd.Execute()
}

func NewRootCmd(commands []Command) *RootCmd {
	rootCmd := &RootCmd{
		cmd: &cobra.Command{
			Use:   "kit",
			Short: "Golang implementation of the Hubkit",
			Long:  `GitHub Kit created by Jakub Tobiasz for Sylius`,
		},
	}

	for _, subcommand := range commands {
		rootCmd.cmd.AddCommand(GetCommand(subcommand))
	}

	return rootCmd
}
