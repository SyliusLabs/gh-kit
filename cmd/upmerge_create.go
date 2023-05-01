package cmd

import (
	"errors"
	"fmt"
	"github.com/SyliusLabs/gh-kit/internal/github"
	"github.com/spf13/cobra"
)

type UpmergeCreateCmd struct {
	cmd      *cobra.Command
	ghClient *github.Client
	ghCli    *github.Cli
}

func (c UpmergeCreateCmd) GetCommand() *cobra.Command {
	return c.cmd
}

func (c *UpmergeCreateCmd) RunE(cmd *cobra.Command, args []string) error {
	base := args[0]
	target := args[1]

	compare, err := c.ghClient.Compare(base, target)
	if nil != err {
		return errors.New(fmt.Sprintf("🛑 An error occurred while comparing %s and %s: %s", base, target, err.Error()))
	}

	if 0 == compare.GetAheadBy() {
		fmt.Printf("⚠️ The %s branch is already up to date with %s \r\n", target, base)
		return nil
	}

	baseReference, err := c.ghClient.GetReference(fmt.Sprintf("heads/%s", base))
	if nil != err {
		return errors.New(fmt.Sprintf("🛑 The base reference %s does not exist", base))
	}

	branchName := fmt.Sprintf("upmerge/%s|%s/%s", base, target, baseReference.GetObject().GetSHA()[0:7])
	err = c.ghClient.CreateReference(fmt.Sprintf("heads/%s", branchName), baseReference.GetObject().GetSHA())
	if nil != err {
		return errors.New(fmt.Sprintf("🛑 The upmerge branch could not be created: %s", err.Error()))
	}

	prTitle := fmt.Sprintf("Upmerge %s into %s", base, target)
	err = c.ghCli.CreatePullRequest(prTitle, prTitle, branchName, target)
	if nil != err {
		return errors.New(fmt.Sprintf("🛑 The upmerge pull request could not be created: %s", err.Error()))
	}

	fmt.Println("✅ The upmerge pull request has been created")
	return nil
}

func NewUpmergeCreateCmd(ghClient *github.Client, ghCli *github.Cli) UpmergeCreateCmd {
	cmd := &cobra.Command{
		Use:   "create <base> <target>",
		Short: "Create an upmerge pull request",
		Long:  `Create an upmerge pull request`,
		Args:  cobra.ExactArgs(2),
	}

	upmergeCreateCmd := new(UpmergeCreateCmd)
	upmergeCreateCmd.cmd = cmd
	upmergeCreateCmd.cmd.RunE = upmergeCreateCmd.RunE
	upmergeCreateCmd.ghClient = ghClient
	upmergeCreateCmd.ghCli = ghCli

	return *upmergeCreateCmd
}
