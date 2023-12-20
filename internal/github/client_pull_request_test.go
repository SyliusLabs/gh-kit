package github

import (
	"github.com/SyliusLabs/gh-kit/internal/test"
	"github.com/google/go-github/v57/github"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_GetPullRequest(t *testing.T) {
	t.Cleanup(gock.Off)
	gock.New("https://api.github.com").
		Post("/repos/JohnnyWalker/BlackLabel/issues/5/comments").
		BodyString(`{"body": "hello world!"}`).
		Reply(204).
		JSON(`{}`)
	restClient, _ := gh.RESTClient(nil)
	githubClient := &Client{
		RestClient: restClient,
		Repository: &test.FakeRepository{},
	}
	err := githubClient.AddCommentToIssue(5, "hello world!")

	assert.NoError(t, err)
	assert.True(t, gock.IsDone())
}

func TestClient_GetPullRequestResultInError(t *testing.T) {
	t.Cleanup(gock.Off)
	gock.New("https://api.github.com").
		Post("/repos/JohnnyWalker/BlackLabel/issues/6/comments").
		BodyString(`{"body": "hello world!"}`).
		ReplyError(&github.ErrorResponse{Message: "Some kind of error"})
	restClient, _ := gh.RESTClient(nil)
	githubClient := &Client{
		RestClient: restClient,
		Repository: &test.FakeRepository{},
	}
	err := githubClient.AddCommentToIssue(6, "hello world!")

	assert.Error(t, err, "Some kind of error")
	assert.True(t, gock.IsDone())
}

func TestClient_GetCommitsInPullRequest(t *testing.T) {
	t.Cleanup(gock.Off)
	gock.New("https://api.github.com").
		Get("/repos/JohnnyWalker/BlackLabel/pulls/5/commits").
		Reply(200).
		JSON(`[
			{
				"sha": "f1e7b8d0c0f7e5e6c9d6b8a8c9b9a1a2b3b4b5b6",
				"commit": {
					"message": "Merge pull request #1 from JohnnyWalker/feature/1"
				}
			}
		]`)
	restClient, _ := gh.RESTClient(nil)
	githubClient := &Client{
		RestClient: restClient,
		Repository: &test.FakeRepository{},
	}
	commits, err := githubClient.GetCommitsInPullRequest(5)

	assert.NoError(t, err)
	assert.True(t, gock.IsDone())

	for _, commit := range commits {
		assert.IsType(t, commit, github.RepositoryCommit{})
	}
}

func TestClient_GetCommitsInPullRequestResultInError(t *testing.T) {
	t.Cleanup(gock.Off)
	gock.New("https://api.github.com").
		Get("/repos/JohnnyWalker/BlackLabel/pulls/5/commits").
		ReplyError(&github.ErrorResponse{Message: "Not Found"})
	restClient, _ := gh.RESTClient(nil)
	githubClient := &Client{
		RestClient: restClient,
		Repository: &test.FakeRepository{},
	}
	_, err := githubClient.GetCommitsInPullRequest(5)

	assert.Error(t, err, "Not Found")
	assert.True(t, gock.IsDone())
}
