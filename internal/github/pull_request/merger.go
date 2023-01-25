package pull_request

import (
	"fmt"
	"github.com/cli/go-gh"
	"strconv"
)

func Merge(number int, subject string, body string, strategy string) error {
	cmd := []string{"pr", "merge", strconv.Itoa(number), fmt.Sprintf("--%s", strategy)}

	if "merge" == strategy {
		cmd = append(cmd, "--subject", subject, "--body", body)
	}

	_, _, err := gh.Exec(cmd[:]...)

	return err
}
