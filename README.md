
<p align="center">
  <img src="https://storage.googleapis.com/altipla-external-files/logos/collections.png">
</p>
<br>


[![GoDoc](https://godoc.org/github.com/altipla-consulting/collections?status.svg)](https://godoc.org/github.com/altipla-consulting/collections)
[![Build Status](https://travis-ci.org/altipla-consulting/cron.svg?branch=master)](https://travis-ci.org/altipla-consulting/cron)

Set of functions that help us work with slices and maps.


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
  
  "github.com/altipla-consulting/collections"
)

  func main() {
    goFounders := []string{"Robert Griesemer", "Rob Pike", "Ken Thompson"}
    fmt.Println("Francis McCabe:", collections.HasString(goFounders, "Francis McCabe"))
    fmt.Println("RobertGriesemer:", collections.HasString(goFounders, "RobertGriesemer"))
    fmt.Println("Robert Griesemer:", collections.HasString(goFounders, "Robert Griesemer"))
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
