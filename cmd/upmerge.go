package cmd

import "github.com/spf13/cobra"

type UpmergeCmd struct {
	cmd *cobra.Command
}

func (p UpmergeCmd) GetCommand() *cobra.Command {
	return p.cmd
}

func NewUpmergeCmd(subcommands []Command) UpmergeCmd {
	cmd := &cobra.Command{
		Use:   "upmerge",
		Short: "Manage upmerges",
		Long:  `Manage upmerges`,
	}

	for _, subcommand := range subcommands {
		cmd.AddCommand(GetCommand(subcommand))
	}

	return UpmergeCmd{cmd: cmd}
}
