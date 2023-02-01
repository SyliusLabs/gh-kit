package generator

import (
	"fmt"
	"github.com/google/go-github/v49/github"
)

func GenerateCommitBody(pr github.PullRequest, commits []github.RepositoryCommit) string {
	msg := fmt.Sprintf(`This PR was merged into the %s branch.

Discussion
----------

%s

Commits
-------
`, pr.GetBase().GetRef(), pr.GetBody())

	for _, commit := range commits {
		msg += fmt.Sprintf("\r\n%s %s", commit.GetSHA(), commit.GetCommit().GetMessage())
	}

	return msg
}

func GenerateCommitMessage(category string, number int, title string, author string) string {
	return fmt.Sprintf("%s #%d %s (%s)", category, number, title, author)
}
