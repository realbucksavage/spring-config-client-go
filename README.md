# spring-config-client-go

A simple package that facilitates fetching of configuration from a Spring Cloud Config Server

## Installation

```shell
go get github.com/realbucksavage/spring-config-client-go
```

## Usage

```go
import (
    "log"
    "github.com/realbucksavage/spring-config-client-go"
)

type applicationConfig struct {
    Key1 Key1Type `json:"key1" yaml:"key1"`
    // ...
}

func main() {
    client, err := cloudconfig.NewClient(
        "config:8888",
        "someapp",
        "production",
    )
    if err != nil {
        panic(err)
    }

    // get a reader to configuration
    rdr, err := client.Raw()

    // or, decode configuration directly into a struct
    var appConfig applicationConfig
    err := client.Decode(&appConfig)
}
```

The client can also be customized with these options

### Basic Auth

```go
client, err := cloudconfig.NewClient(
    server, 
    application, 
    profile, 
    cloudconfig.WithBasicAuth("username", "password"),
)
```

### Reading config as YAML instead of (default) JSON

```go
client, err := cloudconfig.NewClient(
    server,
    application,
    profile, 
    cloudconfig.WithFormat(cloudconfig.YAMLFormat)),
)
```

### Using HTTPs (or, setting the scheme of config server's URL)

```go
client, err := cloudconfig.NewClient(
    server,
    application,
    profile, 
    cloudconfig.WithScheme("https"),
)
```