package cmd

import (
	"errors"
	"fmt"
	"github.com/SyliusLabs/gh-kit/internal/extractor"
	"github.com/SyliusLabs/gh-kit/internal/generator"
	"github.com/SyliusLabs/gh-kit/internal/github"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"strconv"
)

type UpmergeMergeCmd struct {
	cmd      *cobra.Command
	ghClient *github.Client
	ghCli    *github.Cli
}

func (c UpmergeMergeCmd) GetCommand() *cobra.Command {
	return c.cmd
}

func (c *UpmergeMergeCmd) RunE(cmd *cobra.Command, args []string) error {
	prNumber, err := strconv.Atoi(args[0])
	if nil != err {
		return errors.New("🛑 The pull request number is not a valid integer")
	}

	pr, prErr := c.ghClient.GetPullRequest(prNumber)
	if nil != prErr {
		errMsg := fmt.Sprintf(
			"🛑 The pull request #%d on %s/%s seems to not exist",
			prNumber,
			c.ghClient.Repository.Owner(),
			c.ghClient.Repository.Name(),
		)
		return errors.New(errMsg)
	}

	if pr.GetMerged() {
		errMsg := fmt.Sprintf(
			"ℹ️ The pull request #%d on %s/%s is already merged by %s\r\n",
			pr.GetNumber(),
			c.ghClient.Repository.Owner(),
			c.ghClient.Repository.Name(),
			pr.GetMergedBy().GetLogin(),
		)
		return errors.New(errMsg)
	}

	headBranch, err := extractor.ExtractBranchBaseRefFromBranchName(pr.GetHead().GetRef())
	if nil != err {
		return errors.New("🛑 The pull request head branch name is not valid")
	}

	mergeSubject := fmt.Sprintf("Merge branch %s into %s", headBranch, pr.GetBase().GetRef())

	commits, commitsErr := c.ghClient.GetCommitsInPullRequest(pr.GetNumber())
	if nil != commitsErr {
		return errors.New("🛑 The pull request commits could not be fetched")
	}

	mergeBody := generator.GenerateCommitBodyForUpmerge(headBranch, commits)

	summaryMsg := `The pull request #%d on %s/%s will be merged with the following parameters:
Subject: %s
Body:
%s`
	fmt.Printf(summaryMsg, pr.GetNumber(), c.ghClient.Repository.Owner(), c.ghClient.Repository.Name(), mergeSubject, mergeBody)

	confirmPrompt := promptui.Select{
		Label: "Are you sure you want to merge it?",
		Items: []string{"yes", "no"},
	}

	_, confirmed, _ := confirmPrompt.Run()

	if "" == confirmed || "no" == confirmed {
		return errors.New("🛑 The pull request has not been merged")
	}

	deleteBranchPrompt := promptui.Select{
		Label: "Do you want to delete the intermediate branch?",
		Items: []string{"yes", "no"},
	}

	_, deleteBranch, _ := deleteBranchPrompt.Run()

	mergeErr := c.ghCli.Merge(pr.GetNumber(), mergeSubject, mergeBody, "merge", "yes" == deleteBranch)

	if nil != mergeErr {
		errMsg := fmt.Sprintf("🛑 The pull request #%d could not be merged. Error message: %s\r\n", pr.GetNumber(), mergeErr.Error())
		return errors.New(errMsg)
	}

	fmt.Printf("✅ The pull request #%d on %s/%s has been merged\r\n", pr.GetNumber(), c.ghClient.Repository.Owner(), c.ghClient.Repository.Name())
	return nil
}

func NewUpmergeMergeCmd(ghClient *github.Client, ghCli *github.Cli) UpmergeMergeCmd {
	cmd := &cobra.Command{
		Use:   "merge <pull_request_number>",
		Short: "Merge an upmerge",
		Args:  cobra.ExactArgs(1),
	}

	upmergeMergeCmd := new(UpmergeMergeCmd)
	upmergeMergeCmd.cmd = cmd
	upmergeMergeCmd.cmd.RunE = upmergeMergeCmd.RunE
	upmergeMergeCmd.ghClient = ghClient
	upmergeMergeCmd.ghCli = ghCli

	return *upmergeMergeCmd
}
