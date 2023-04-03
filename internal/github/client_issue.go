package github

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) AddCommentToIssue(issueNumber int, message string) error {
	req := struct {
		Body string `json:"body"`
	}{
		Body: message,
	}
	buffer := new(bytes.Buffer)
	_ = json.NewEncoder(buffer).Encode(req)

	return c.Post(
		fmt.Sprintf("repos/%s/%s/issues/%d/comments", c.Repository.Owner(), c.Repository.Name(), issueNumber),
		buffer,
		nil,
	)
}
