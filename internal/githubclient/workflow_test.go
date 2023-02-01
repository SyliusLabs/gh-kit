package githubclient

import (
	"github.com/SyliusLabs/gh-kit/internal/test"
	"github.com/cli/go-gh"
	"github.com/google/go-github/v49/github"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestClient_GetWorkflowRunsForHeadSha(t *testing.T) {
	t.Cleanup(gock.Off)
	gock.New("https://api.github.com").
		Get("/repos/JohnnyWalker/BlackLabel/actions/runs").
		MatchParam("event", "pull_request").
		MatchParam("head_sha", "f1e7b8d0c0f7e5e6c9d6b8a8c9b9a1a2b3b4b5b6").
		Reply(200).
		JSON(`{
			"total_count": 1,
			"workflow_runs": [
				{
					"id": 30433642,
					"name": "Build",
					"node_id": "MDEyOldvcmtmbG93IFJ1bjI2OTI4OQ==",
					"check_suite_id": 42,
					"check_suite_node_id": "MDEwOkNoZWNrU3VpdGU0Mg==",
					"head_branch": "master"
				}
			]
		}`)
	restClient, _ := gh.RESTClient(nil)
	githubClient := &Client{
		RestClient: restClient,
		Repository: &test.FakeRepository{},
	}
	workflowRuns, err := githubClient.GetWorkflowRunsForHeadSha("f1e7b8d0c0f7e5e6c9d6b8a8c9b9a1a2b3b4b5b6")

	assert.NoError(t, err)
	assert.True(t, gock.IsDone())
	assert.IsType(t, workflowRuns, github.WorkflowRuns{})
	assert.Equal(t, 1, workflowRuns.GetTotalCount())

	for _, workflowRun := range workflowRuns.WorkflowRuns {
		assert.IsType(t, &github.WorkflowRun{}, workflowRun)
	}
}

func TestClient_GetLatestWorkflowRunForHeadSha(t *testing.T) {
	t.Cleanup(gock.Off)
	gock.New("https://api.github.com").
		Get("/repos/JohnnyWalker/BlackLabel/actions/runs").
		MatchParam("event", "pull_request").
		MatchParam("head_sha", "f1e7b8d0c0f7e5e6c9d6b8a8c9b9a1a2b3b4b5b6").
		Reply(200).
		JSON(`{
			"total_count": 2,
			"workflow_runs": [
				{
					"id": 30433642,
					"name": "Build",
					"node_id": "MDEyOldvcmtmbG93IFJ1bjI2OTI4OQ==",
					"check_suite_id": 42,
					"check_suite_node_id": "MDEwOkNoZWNrU3VpdGU0Mg==",
					"head_branch": "master"
				},
				{
					"id": 30433669,
					"name": "Build2",
					"node_id": "MDEyOldvcmtm5593IFJ1bjI2OTI4OQ==",
					"check_suite_id": 45,
					"check_suite_node_id": "MDEwOkNoZWffU3VpdGU0Mg==",
					"head_branch": "master"
				}
			]
		}`)
	restClient, _ := gh.RESTClient(nil)
	githubClient := &Client{
		RestClient: restClient,
		Repository: &test.FakeRepository{},
	}
	latestWorkflowRun, err := githubClient.GetLatestWorkflowRunForHeadSha("f1e7b8d0c0f7e5e6c9d6b8a8c9b9a1a2b3b4b5b6")

	assert.NoError(t, err)
	assert.True(t, gock.IsDone())
	assert.IsType(t, github.WorkflowRun{}, latestWorkflowRun)
	assert.Equal(t, int64(30433669), latestWorkflowRun.GetID())
}

func TestClient_GetLatestWorkflowRunForHeadSha_NoRuns(t *testing.T) {
	t.Cleanup(gock.Off)
	gock.New("https://api.github.com").
		Get("/repos/JohnnyWalker/BlackLabel/actions/runs").
		MatchParam("event", "pull_request").
		MatchParam("head_sha", "f1e7b8d0c0f7e5e6c9d6b8a8c9b9a1a2b3b4b5b6").
		Reply(200).
		JSON(`{
			"total_count": 0,
			"workflow_runs": []
		}`)
	restClient, _ := gh.RESTClient(nil)
	githubClient := &Client{
		RestClient: restClient,
		Repository: &test.FakeRepository{},
	}
	_, err := githubClient.GetLatestWorkflowRunForHeadSha("f1e7b8d0c0f7e5e6c9d6b8a8c9b9a1a2b3b4b5b6")

	assert.Error(t, err, "no workflow runs found for SHA f1e7b8d0c0f7e5e6c9d6b8a8c9b9a1a2b3b4b5b6")
	assert.True(t, gock.IsDone())
}

func TestClient_RerunFailedJobs(t *testing.T) {
	t.Cleanup(gock.Off)
	gock.New("https://api.github.com").
		Post("/repos/JohnnyWalker/BlackLabel/actions/runs/30433642/rerun").
		Reply(201).
		JSON(`{}`)
	restClient, _ := gh.RESTClient(nil)
	githubClient := &Client{
		RestClient: restClient,
		Repository: &test.FakeRepository{},
	}
	err := githubClient.RerunFailedJobs(30433642)

	assert.NoError(t, err)
	assert.True(t, gock.IsDone())
}
