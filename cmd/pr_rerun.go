package cmd

import (
	"errors"
	"fmt"
	"github.com/SyliusLabs/gh-kit/internal/github"
	"github.com/spf13/cobra"
	"strconv"
)

func (c *PrRerunCmd) RunE(cmd *cobra.Command, args []string) error {
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
			"The pull request #%d on %s/%s is already merged by %s",
			pr.GetNumber(),
			c.ghClient.Repository.Owner(),
			c.ghClient.Repository.Name(),
			pr.GetMergedBy().GetLogin(),
		)
		return errors.New(errMsg)
	}

	workflowRun, workflowRunErr := c.ghClient.GetLatestWorkflowRunForHeadSha(pr.GetHead().GetSHA())
	if nil != workflowRunErr {
		return errors.New("🛑 The pull request does not have any workflow runs")
	}

	if workflowRun.GetStatus() != "completed" || workflowRun.GetConclusion() != "failure" {
		return errors.New("🛑 The workflow run is not completed or it did not fail")
	}

	rerunErr := c.ghClient.RerunFailedJobs(workflowRun.GetID())
	if rerunErr != nil {
		errMsg := fmt.Sprintf("🛑 The workflow failed jobs could not be rerun. Reason: %s", rerunErr.Error())
		return errors.New(errMsg)
	}

	fmt.Println("✅ The workflow failed jobs have been rerun")
	return nil
}

type PrRerunCmd struct {
	cmd      *cobra.Command
	ghClient *github.Client
}

func (c PrRerunCmd) GetCommand() *cobra.Command {
	return c.cmd
}

func NewPrRerunCmd(ghClient *github.Client) PrRerunCmd {
	cmd := &cobra.Command{
		Use:   "rerun-failed <pull_request_number>",
		Short: "Reruns a failed GitHub Workflow jobs for a given PR",
		Args:  cobra.ExactArgs(1),
	}

	prRerunCmd := new(PrRerunCmd)
	prRerunCmd.cmd = cmd
	prRerunCmd.cmd.RunE = prRerunCmd.RunE
	prRerunCmd.ghClient = ghClient

	return *prRerunCmd
}
