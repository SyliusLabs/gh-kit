package githubclient

import (
	"fmt"
	"github.com/google/go-github/v49/github"
)

func (c *Client) GetPullRequest(number int) (github.PullRequest, error) {
	response := github.PullRequest{}

	err := c.Get(fmt.Sprintf("repos/%s/%s/pulls/%d", c.Repository.Owner(), c.Repository.Name(), number), &response)
	if nil != err {
		return github.PullRequest{}, err
	}

	return response, nil
}

func (c *Client) GetCommitsInPullRequest(number int) ([]github.RepositoryCommit, error) {
	var response []github.RepositoryCommit

	err := c.Get(
		fmt.Sprintf("repos/%s/%s/pulls/%d/commits", c.Repository.Owner(), c.Repository.Name(), number),
		&response,
	)
	if nil != err {
		return []github.RepositoryCommit{}, err
	}

	return response, nil
}
