
# collections

Utilidades de Go para tratar con colecciones.

[![GoDoc](https://godoc.org/github.com/altipla-consulting/collections?status.svg)](https://godoc.org/github.com/altipla-consulting/collections)

> Set of functions that help us work with slices and maps.


### Install

```shell
go get github.com/altipla-consulting/collections
```

This library has no external dependencies outside the Go standard library.


### Usage
```go
package main

import (
  "fmt"
  "gitthub.com/altipla-consulting/collections"

  func main() {
    goFounders := []string{"Robert Griesemer", "Rob Pike", "Ken Thompson"}
    fmt.Printf("Fracis McCabe: %t \n", collections.HasString(goFounders, "Francis McCabe"))
    fmt.Printf("RobertGriesemer: %t \n", collections.HasString(goFounders, "RobertGriesemer"))
    fmt.Printf("Robert Griesemer: %t \n", collections.HasString(goFounders, "Robert Griesemer"))
  }
)
```

Result:
```
Fracis McCabe: false 
RobertGriesemer: false 
Robert Griesemer: true 
```

### Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using ```gofmt```.

### License

[MIT License](LICENSE)
