package github

import (
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
)

func NewRestClient() api.RESTClient {
	client, _ := gh.RESTClient(nil)

	return client
}
