package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var rerunCmd = &cobra.Command{
	Use:   "rerun-failed <pull_request_number>",
	Short: "Reruns a failed GitHub Workflow jobs for a given PR",
	Run: func(cmd *cobra.Command, args []string) {
		prNumber, err := strconv.Atoi(args[0])
		if nil != err {
			fmt.Println("🛑 The pull request number is not a valid integer")
			return
		}

		pr, prErr := ghClient.GetPullRequest(prNumber)
		if nil != prErr {
			fmt.Println(prErr.Error())
			return
		}

		if pr.GetMerged() {
			fmt.Printf(
				"The pull request #%d on %s/%s is already merged by %s\r\n",
				pr.GetNumber(),
				ghClient.Repository.Owner(),
				ghClient.Repository.Name(),
				pr.GetMergedBy().GetLogin(),
			)
			return
		}

		workflowRun, workflowRunErr := ghClient.GetLatestWorkflowRunForHeadSha(pr.GetHead().GetSHA())
		if nil != workflowRunErr {
			fmt.Println("🛑 The pull request does not have any workflow runs")
			return
		}

		if workflowRun.GetStatus() != "completed" || workflowRun.GetConclusion() != "failure" {
			fmt.Println("🛑 The workflow run is not completed or it did not fail")
			return
		}

		rerunErr := ghClient.RerunFailedJobs(workflowRun.GetID())
		if rerunErr != nil {
			fmt.Printf("🛑 The workflow failed jobs could not be rerun. Reason: %s\r\n", rerunErr.Error())
			return
		}

		fmt.Println("✅ The workflow failed jobs have been rerun")
	},
}

func init() {
	prCmd.AddCommand(rerunCmd)
}
