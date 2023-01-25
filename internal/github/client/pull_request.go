package client

import (
	"fmt"
	"github.com/cli/go-gh/pkg/api"
	"github.com/google/go-github/v49/github"
	"github.com/jakubtobiasz/gh-kit/internal/factory"
	"github.com/jakubtobiasz/gh-kit/internal/github/current_repository"
)

var client api.RESTClient

func init() {
	client = factory.CreateGitHubClient()
}

func GetPullRequestData(repo current_repository.CurrentRepository, number string) (github.PullRequest, error) {
	var response = github.PullRequest{}

	err := client.Get(fmt.Sprintf("repos/%s/%s/pulls/%s", repo.Owner, repo.Name, number), &response)
	if nil != err {
		return github.PullRequest{}, err
	}

	return response, nil
}
