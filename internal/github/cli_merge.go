package github

import (
	"fmt"
	"strconv"
)

func (c Cli) Merge(number int, subject string, body string, strategy string, deleteBranch bool) error {
	cmd := []string{"pr", "merge", strconv.Itoa(number), fmt.Sprintf("--%s", strategy)}

	if "merge" == strategy {
		cmd = append(cmd, "--subject", subject, "--body", body)
	}

	if deleteBranch {
		cmd = append(cmd, "--delete-branch")
	}

	_, _, err := c.Exec(cmd...)

	return err
}
