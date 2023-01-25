package current_repository

import (
	"github.com/cli/go-gh"
)

type CurrentRepository struct {
	Owner string
	Name  string
}

func GetCurrentRepositoryData() (CurrentRepository, error) {
	repo, repoErr := gh.CurrentRepository()

	if nil != repoErr {
		return CurrentRepository{}, repoErr
	}

	return CurrentRepository{
		repo.Owner(),
		repo.Name(),
	}, nil
}
