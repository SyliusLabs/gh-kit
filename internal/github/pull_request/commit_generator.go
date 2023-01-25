package pull_request

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
