package github

import (
	"fmt"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/repository"
	"github.com/google/go-github/v57/github"
)

func (c *Client) Compare(base string, target string) (github.CommitsComparison, error) {
	var response github.CommitsComparison

	err := c.Get(
		fmt.Sprintf("repos/%s/%s/compare/%s...%s", c.Repository.Owner(), c.Repository.Name(), target, base),
		&response,
	)
	if nil != err {
		return github.CommitsComparison{}, err
	}

	return response, nil
}

func NewRepository() repository.Repository {
	repo, _ := gh.CurrentRepository()

	return repo
}
