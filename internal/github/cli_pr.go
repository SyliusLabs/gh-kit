package github

func (c Cli) CreatePullRequest(title string, body string, sourceBranch string, targetBranch string) error {
	cmd := []string{"pr", "create", "--title", title, "--body", body, "--base", targetBranch, "--head", sourceBranch}

	_, _, err := c.Exec(cmd...)

	return err
}
