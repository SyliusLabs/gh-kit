package pull_request

import (
	"fmt"
	"github.com/cli/go-gh"
	"strconv"
)

func Merge(number int, subject string, strategy string) error {
	_, _, err := gh.Exec("pr", "merge", strconv.Itoa(number), "--subject", subject, fmt.Sprintf("--%s", strategy))

	return err
}
