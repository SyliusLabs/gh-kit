package github

import (
	"fmt"
	"github.com/google/go-github/v57/github"
)

func (c *Client) GetWorkflowRunsForHeadSha(headSha string) (github.WorkflowRuns, error) {
	response := github.WorkflowRuns{}

	err := c.Get(
		fmt.Sprintf("repos/%s/%s/actions/runs?event=pull_request&head_sha=%s", c.Repository.Owner(), c.Repository.Name(), headSha),
		&response,
	)
	if nil != err {
		return github.WorkflowRuns{}, err
	}

	return response, nil
}

func (c *Client) GetLatestWorkflowRunForHeadSha(headSha string) (github.WorkflowRun, error) {
	workflowRuns, workflowRunsErr := c.GetWorkflowRunsForHeadSha(headSha)
	if nil != workflowRunsErr {
		return github.WorkflowRun{}, workflowRunsErr
	}

	if 0 == workflowRuns.GetTotalCount() {
		return github.WorkflowRun{}, fmt.Errorf("no workflow runs found for SHA %s", headSha)
	}

	return *workflowRuns.WorkflowRuns[workflowRuns.GetTotalCount()-1], nil
}

func (c *Client) RerunFailedJobs(runId int64) error {
	return c.Post(
		fmt.Sprintf("repos/%s/%s/actions/runs/%d/rerun-failed-jobs", c.Repository.Owner(), c.Repository.Name(), runId),
		nil,
		nil,
	)
}
