package cmd

import (
	githubCli "github.com/SyliusLabs/gh-kit/internal/githubcli"
	githubClient "github.com/SyliusLabs/gh-kit/internal/githubclient"
	"github.com/spf13/cobra"
)

var ghCli githubCli.GhCliExecutor
var ghClient githubClient.Client

var rootCmd = &cobra.Command{
	Use:   "kit",
	Short: "Golang implementation of the Hubkit",
	Long:  `GitHub Kit created by Jakub Tobiasz for Sylius`,
}

func init() {
	var err error

	ghCli = githubCli.NewGhCli()

	var client *githubClient.Client
	client, err = githubClient.NewClient(nil, nil)
	if nil != err {
		panic(err)
	}

	ghClient = *client
}

func Run() error {
	return rootCmd.Execute()
}
