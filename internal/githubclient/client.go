package githubclient

import (
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
	"github.com/cli/go-gh/pkg/repository"
	"io"
)

type Client struct {
	RestClient api.RESTClient
	Repository repository.Repository
}

func (c *Client) Get(path string, response interface{}) error {
	return c.RestClient.Get(path, response)
}

func (c *Client) Post(path string, body io.Reader, response interface{}) error {
	return c.RestClient.Post(path, body, response)
}

func (c *Client) Patch(path string, body io.Reader, response interface{}) error {
	return c.RestClient.Patch(path, body, response)
}

func (c *Client) Delete(path string, response interface{}) error {
	return c.RestClient.Delete(path, response)
}

func (c *Client) Put(path string, body io.Reader, response interface{}) error {
	return c.RestClient.Put(path, body, response)
}

func NewClient(restClient *api.RESTClient, repository *repository.Repository) (*Client, error) {
	if nil == restClient {
		ghRestClient, err := gh.RESTClient(nil)
		if nil != err {
			return nil, err
		}

		restClient = &ghRestClient
	}

	if nil == repository {
		repo, err := gh.CurrentRepository()
		if nil != err {
			return nil, err
		}

		repository = &repo
	}

	return &Client{
		RestClient: *restClient,
		Repository: *repository,
	}, nil
}
