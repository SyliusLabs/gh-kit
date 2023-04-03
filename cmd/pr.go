package cmd

import "github.com/spf13/cobra"

type PrCmd struct {
	cmd *cobra.Command
}

func (p PrCmd) GetCommand() *cobra.Command {
	return p.cmd
}

func NewPrCmd(subcommands []Command) PrCmd {
	cmd := &cobra.Command{
		Use:   "pr",
		Short: "Manage pull requests",
		Long:  `Manage pull requests`,
	}

	for _, subcommand := range subcommands {
		cmd.AddCommand(GetCommand(subcommand))
	}

	return PrCmd{cmd: cmd}
}
