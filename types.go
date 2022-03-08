package cloudconfig

import "io"

// Client facilitates connection to the spring config server. The config URL is formed like this:
// http(s)://{ServerAddr}{/Branch}/{Application}{-Profile}.{json|yaml}
type Client interface {

	// Raw returns an io.ReadCloser of raw config loaded from server
	Raw() (io.ReadCloser, error)

	// Decode reads the configuration from server, and decodes it into a struct.
	Decode(v interface{}) error
}

type Format string

var (
	JSONFormat Format = "json"
	YAMLFormat Format = "yaml"
)

func (f Format) Valid() bool {
	return f == JSONFormat || f == YAMLFormat
}
