# gostatement

[![test_and_lint](https://github.com/fpuc/gostatement/actions/workflows/workflows.yml/badge.svg?branch=main)](https://github.com/fpuc/gostatement/actions/workflows/workflows.yml)

gostatement is an analyzer checking for occurrence of `go` statements. You may want to use a custom 
func wrapping the statement utilizing recover, logging, metrics...

## Instruction

```sh
go install github.com/fpuc/gostatement/cmd/gostatement
```

## Usage

```go
package main

import (
	"fmt"
	"os"
	"testing"
)

func main() {
    go func() {
        // foo
    }()
}
```


```console 
go vet -vettool=$(which gostatement) ./...

./main.go:4:2: go statement found
```
