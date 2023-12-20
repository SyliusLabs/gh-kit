package generator

import (
	"fmt"
	"github.com/google/go-github/v57/github"
	"strings"
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
		commitMsg := strings.Split(commit.GetCommit().GetMessage(), "\n")
		firstLineOfCommitMsg := commitMsg[0]
		msg += fmt.Sprintf("  %s\r\n", firstLineOfCommitMsg)
	}

	return msg
}

func GenerateCommitBodyForUpmerge(headBranchName string, commits []github.RepositoryCommit) string {
	msg := fmt.Sprintf("* %s:\r\n", headBranchName)

	for _, commit := range commits {
		if len(commit.Parents) > 1 {
			continue
		}

		commitMsg := strings.Split(commit.GetCommit().GetMessage(), "\n")
		firstLineOfCommitMsg := commitMsg[0]
		msg += fmt.Sprintf("  %s\r\n", firstLineOfCommitMsg)
	}

	return msg
}

func GenerateCommitMessage(category string, number int, title string, author string) string {
	return fmt.Sprintf("%s #%d %s (%s)", category, number, title, author)
}
