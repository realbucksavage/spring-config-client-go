# spring-config-client-go

A simple package that facilitates fetching of configuration from a Spring Cloud Config Server

## Installation

```shell
$ go get github.com/realbucksavage/spring-config-client-go
```

## Usage

```go
import (
    "log"
    "github.com/realbucksavage/spring-config-client-go"
)

func main() {
    client := &configclient.Client{
        ServerAddr:  "config:8888",
        Application: "myapp",
        Profile:     "prod",
        Format:      "json",
    }

    bytes, err := client.FetchConfig()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Config as JSON: %s\n", string(bytes))
}
```
