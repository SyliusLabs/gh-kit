package pull_request

import (
	"fmt"
	"github.com/SyliusLabs/gh-kit/internal/factory"
	"github.com/SyliusLabs/gh-kit/internal/github/current_repository"
	"github.com/cli/go-gh/pkg/api"
	"github.com/google/go-github/v49/github"
)

var client api.RESTClient

func init() {
	client = factory.CreateGitHubClient()
}

func GetData(repo current_repository.CurrentRepository, number string) (github.PullRequest, error) {
	var response = github.PullRequest{}

	err := client.Get(fmt.Sprintf("repos/%s/%s/pulls/%s", repo.Owner, repo.Name, number), &response)
	if nil != err {
		return github.PullRequest{}, err
	}

	return response, nil
}

func GetCommits(repo current_repository.CurrentRepository, number int) ([]github.RepositoryCommit, error) {
	var response []github.RepositoryCommit

	err := client.Get(fmt.Sprintf("repos/%s/%s/pulls/%d/commits", repo.Owner, repo.Name, number), &response)
	if nil != err {
		return []github.RepositoryCommit{}, err
	}

	return response, nil
}
