package factory

import (
	"fmt"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
)

func CreateGitHubClient() api.RESTClient {
	client, clientErr := gh.RESTClient(nil)

	if nil != clientErr {
		panic(fmt.Sprintf("Could not initialize GitHub client: %s", clientErr.Error()))
	}

	return client
}
