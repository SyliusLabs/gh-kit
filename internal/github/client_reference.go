package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v49/github"
)

func (c *Client) GetReference(ref string) (github.Reference, error) {
	response := github.Reference{}

	err := c.Get(fmt.Sprintf("repos/%s/%s/git/ref/%s", c.Repository.Owner(), c.Repository.Name(), ref), &response)
	if nil != err {
		return github.Reference{}, err
	}

	return response, nil
}

func (c *Client) CreateReference(ref string, sha string) error {
	req := struct {
		Ref string `json:"ref"`
		Sha string `json:"sha"`
	}{
		Ref: fmt.Sprintf("refs/%s", ref),
		Sha: sha,
	}
	buffer := new(bytes.Buffer)
	_ = json.NewEncoder(buffer).Encode(req)

	return c.Post(
		fmt.Sprintf("repos/%s/%s/git/refs", c.Repository.Owner(), c.Repository.Name()),
		buffer,
		nil,
	)
}
