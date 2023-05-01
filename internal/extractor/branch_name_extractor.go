package extractor

import (
	"errors"
	"fmt"
	"regexp"
)

func ExtractBranchBaseRefFromBranchName(branchName string) (string, error) {
	pattern := `upmerge/([^|]+)\|`

	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return "", err
	}

	matches := regex.FindStringSubmatch(branchName)
	headBranchRef := ""
	if len(matches) > 1 {
		headBranchRef = matches[1]
	} else {
		return "", errors.New("no matches found")
	}

	return headBranchRef, nil
}
