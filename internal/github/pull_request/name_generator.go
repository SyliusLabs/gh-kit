package pull_request

import "fmt"

func GenerateSubjectWithCategory(category string, number int, title string, author string) string {
	return fmt.Sprintf("%s #%d %s (%s)", category, number, title, author)
}
