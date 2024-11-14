# go-version

[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/olenindenis/go-version.svg)](https://pkg.go.dev/github.com/olenindenis/go-version)
[![Go Report Card](https://goreportcard.com/badge/github.com/olenindenis/vault-extractor)](https://goreportcard.com/report/github.com/olenindenis/vault-extractor)


This package helps to include version info to your go application.

#### Installation
Make sure that Go is installed on your computer.
Type the following command in your terminal:
```sh
go get github.com/olenindenis/go-version
```
After it the package is ready to use.
#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/olenindenis/go-version"
```

#### Usage

Before running your code, you need to run the command:
```sh
make build
```
or create your own as shown in the example:
```sh
go build -ldflags=${linker_flags} -o=./main ./main.go
```

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/olenindenis/go-version"
)

func main() {
    http.HandleFunc("/version", func (w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(version.Info())
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```