package cmd

import (
	"errors"
	"fmt"
	"github.com/SyliusLabs/gh-kit/internal/generator"
	"github.com/SyliusLabs/gh-kit/internal/github"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"strconv"
)

var allowedMergeStrategies = map[string]bool{"merge": true, "squash": true, "rebase": true}

type PrMergeCmd struct {
	cmd      *cobra.Command
	ghClient *github.Client
	ghCli    *github.Cli
}

func (c PrMergeCmd) GetCommand() *cobra.Command {
	return c.cmd
}

func (c *PrMergeCmd) RunE(cmd *cobra.Command, args []string) error {
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

	mergeStrategy, _ := cmd.Flags().GetString("strategy")
	if "" == mergeStrategy {
		strategyPrompt := promptui.Select{
			Label: "Select merge strategy",
			Items: []string{"merge", "squash", "rebase"},
		}

		_, mergeStrategy, _ = strategyPrompt.Run()
	}

	if !allowedMergeStrategies[mergeStrategy] || "" == mergeStrategy {
		return errors.New("🛑 The merge strategy is not allowed or is empty. Allowed strategies are: merge, squash, rebase")
	}

	mergeCategory, _ := cmd.Flags().GetString("category")

	if "" == mergeCategory {
		categoryPrompt := promptui.Select{
			Label: "Select merge category",
			Items: []string{"minor", "feature", "bugfix", "docs", "refactor", "test", "chore"},
		}

		_, mergeCategory, _ = categoryPrompt.Run()
	}

	if "" == mergeCategory {
		return errors.New("⚠️ The Pull Request category must be provided")
	}

	mergeSubject := generator.GenerateCommitMessage(mergeCategory, pr.GetNumber(), pr.GetTitle(), pr.GetUser().GetLogin())

	commits, commitsErr := c.ghClient.GetCommitsInPullRequest(pr.GetNumber())
	if nil != commitsErr {
		return errors.New("🛑 The pull request commits could not be fetched")
	}

	mergeBody := generator.GenerateCommitBody(pr, commits)

	summaryMsg := `The pull request #%d on %s/%s will be merged with the following parameters:
Subject: %s
Strategy: %s
Category: %s
Body:
%s`
	fmt.Printf(summaryMsg, pr.GetNumber(), c.ghClient.Repository.Owner(), c.ghClient.Repository.Name(), mergeSubject, mergeStrategy, mergeCategory, mergeBody)

	confirmPrompt := promptui.Select{
		Label: "Are you sure you want to merge it?",
		Items: []string{"yes", "no"},
	}

	_, confirmed, _ := confirmPrompt.Run()

	if "" == confirmed || "no" == confirmed {
		return errors.New("🛑 The pull request has not been merged")
	}

	deleteBranchPrompt := promptui.Select{
		Label: "Do you want to delete the branch?",
		Items: []string{"yes", "no"},
	}

	_, deleteBranch, _ := deleteBranchPrompt.Run()

	mergeErr := c.ghCli.Merge(pr.GetNumber(), mergeSubject, mergeBody, mergeStrategy, "yes" == deleteBranch)

	if nil != mergeErr {
		errMsg := fmt.Sprintf("🛑 The pull request #%d could not be merged. Error message: %s\r\n", pr.GetNumber(), mergeErr.Error())
		return errors.New(errMsg)
	}

	fmt.Printf("✅ The pull request #%d on %s/%s has been merged\r\n", pr.GetNumber(), c.ghClient.Repository.Owner(), c.ghClient.Repository.Name())

	if noThankYou, _ := cmd.Flags().GetBool("skip-thankyou"); noThankYou {
		fmt.Println("⏩ The thank you comment has not been added")
		return nil
	}

	commentErr := c.ghClient.AddCommentToIssue(pr.GetNumber(), fmt.Sprintf("Thank you, @%s!", pr.GetUser().GetLogin()))
	if nil != commentErr {
		errMsg := fmt.Sprintf("⚠️ The comment could not be added. Error message: %s\r\n", commentErr.Error())
		return errors.New(errMsg)
	}

	fmt.Printf("✅ The thank you comment has been added to the pull request #%d\r\n", pr.GetNumber())
	return nil
}

func NewPrMergeCmd(ghClient *github.Client, ghCli *github.Cli) PrMergeCmd {
	cmd := &cobra.Command{
		Use:   "merge <pull_request_number>",
		Short: "Merge a pull request",
		Args:  cobra.ExactArgs(1),
	}

	cmd.PersistentFlags().StringP("category", "c", "", "Category of the pull request")
	cmd.PersistentFlags().StringP("strategy", "s", "", "Merge strategy to be used")
	cmd.PersistentFlags().Bool("skip-thankyou", false, "Do not add a thank you comment to the pull request")

	prMergeCmd := new(PrMergeCmd)
	prMergeCmd.cmd = cmd
	prMergeCmd.cmd.RunE = prMergeCmd.RunE
	prMergeCmd.ghClient = ghClient
	prMergeCmd.ghCli = ghCli

	return *prMergeCmd
}
