package configclient

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Client facilitates connection to the spring config server. The config URL is formed like this:
// http(s)://{ServerAddr}{/Branch}/{Application}{-Profile}.{json|yaml}
type Client struct {

	// ServerAddr is the address (with port) of the config server
	ServerAddr string

	// Application is the application name
	Application string

	// Profile is the profile the application is running in
	Profile string

	// Branch is used to determine the branch name, if the config server is using git as backend.
	Branch string

	// Fomat is either json or yaml
	Format string

	// Authorize determines whether to perform basic authroization or not.
	Authorize bool

	// UseHTTP controls the http/https protocol in the URL
	UseHTTPS bool

	// BasicAuth represents the username and password authroization required to complete the request
	BasicAuth Authorization
}

// Authorization represents the username password combination
type Authorization struct {
	Username string
	Password string
}

// FetchConfig connects to the config server and returns the json/yaml response as []byte
func (c *Client) FetchConfig() ([]byte, error) {
	httpClient := &http.Client{}

	url := "http"

	if c.UseHTTPS {
		url += "s"
	}

	url += "://" + c.ServerAddr
	if c.Branch != "" {
		url += "/" + c.Branch
	}

	url += "/" + c.Application
	if c.Profile != "" {
		url += "-" + c.Profile
	}

	switch strings.ToLower(c.Format) {
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
