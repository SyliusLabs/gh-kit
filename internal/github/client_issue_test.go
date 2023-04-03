package github

import (
	"github.com/SyliusLabs/gh-kit/internal/test"
	"github.com/cli/go-gh"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestClient_AddCommentToIssue(t *testing.T) {
	t.Cleanup(gock.Off)
	gock.New("https://api.github.com").
		Post("/repos/JohnnyWalker/BlackLabel/issues/5/comments").
		BodyString(`{"body": "hello world!"}`).
		Reply(204).
		JSON(`{}`)
	restClient, _ := gh.RESTClient(nil)
	githubClient := &Client{
		RestClient: restClient,
		Repository: &test.FakeRepository{},
	}
	err := githubClient.AddCommentToIssue(5, "hello world!")

	assert.NoError(t, err)
	assert.True(t, gock.IsDone())
}
