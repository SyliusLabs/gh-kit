package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "kit",
	Short: "Golang implementation of the Hubkit",
	Long:  `GitHub Kit created by Jakub Tobiasz for Sylius`,
}

func Run() error {
	return rootCmd.Execute()
}
