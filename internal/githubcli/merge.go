package githubcli

import (
	"fmt"
	"strconv"
)

func Merge(c GhCliExecutor, number int, subject string, body string, strategy string) error {
	cmd := []string{"pr", "merge", strconv.Itoa(number), fmt.Sprintf("--%s", strategy)}

	if "merge" == strategy {
		cmd = append(cmd, "--subject", subject, "--body", body)
	}

	_, _, err := c.Exec(cmd...)

	return err
}
