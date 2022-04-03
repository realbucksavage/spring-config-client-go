package cloudconfig

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type defaultClient struct {
	server      string
	application string
	profile     string
	branch      string
	format      Format
	scheme      string
	basicAuth   *basicAuthInfo

	httpClient *http.Client
}

type basicAuthInfo struct {
	username string
	password string
}

func NewClient(server, application, profile string, opts ...ClientOption) (Client, error) {

	if server == "" {
		return nil, errors.New("server is required")
	}

	if application == "" {
		return nil, errors.New("application is required")
	}

	if profile == "" {
		return nil, errors.New("a base profile is required")
	}

	c := &defaultClient{
		server:      server,
		application: application,
		profile:     profile,
		format:      JSONFormat,
		scheme:      "http",
		httpClient:  &http.Client{},
	}

	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (d *defaultClient) Raw() (io.ReadCloser, error) {

	u := d.buildURL()
	if _, err := url.Parse(u); err != nil {
		return nil, errors.Wrapf(err, "invalid config url [%s]", u)
	}

	request, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	if d.basicAuth != nil {
		request.SetBasicAuth(d.basicAuth.username, d.basicAuth.password)
	}

	response, err := d.httpClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "config resolution failed")
	}

	return response.Body, nil
}

func (d *defaultClient) Decode(v interface{}) error {

	reader, err := d.Raw()
	if err != nil {
		return err
	}

	if d.format == JSONFormat {
		return json.NewDecoder(reader).Decode(v)
	}

	return yaml.NewDecoder(reader).Decode(v)
}

func (d *defaultClient) buildURL() string {

	u := fmt.Sprintf("%s://%s", d.scheme, d.server)

	if d.branch != "" {
		u = fmt.Sprintf("%s/%s", u, d.branch)
	}

	return fmt.Sprintf("%s/%s-%s.%s", u, d.application, d.profile, d.format)
}
