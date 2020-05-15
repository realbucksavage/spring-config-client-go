package configclient

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	Server      string
	Application string
	Pofile      string
	Branch      string
	Format      string
	Authorize   bool
	BasicAuth   Authorization
}

type Authorization struct {
	Username string
	Password string
}

func (c *Client) FetchConfig() ([]byte, error) {
	httpClient := &http.Client{}

	url := c.Server
	if c.Branch != "" {
		url += "/" + c.Branch
	}

	url += "/" + c.Application
	if c.Pofile != "" {
		url += "-" + c.Pofile
	}

	switch c.Format {
	case "yaml":
		url += ".yaml"
		break
	case "json":
	default:
		url += ".json"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.Authorize {
		req.SetBasicAuth(c.BasicAuth.Username, c.BasicAuth.Password)
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
