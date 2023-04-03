package github

import (
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/repository"
)

func NewRepository() repository.Repository {
	repo, _ := gh.CurrentRepository()

	return repo
}
