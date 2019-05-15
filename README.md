# Space Tree

[![Build Status](https://travis-ci.org/mauricioklein/go-spacetree.svg?branch=master)](https://travis-ci.org/mauricioklein/go-spacetree)
![GitHub release](https://img.shields.io/github/release/mauricioklein/go-spacetree.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Maintainability](https://api.codeclimate.com/v1/badges/65187a309a09d5999bac/maintainability)](https://codeclimate.com/github/mauricioklein/go-spacetree/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/65187a309a09d5999bac/test_coverage)](https://codeclimate.com/github/mauricioklein/go-spacetree/test_coverage)
[![GoDoc](https://godoc.org/github.com/mauricioklein/go-spacetree?status.svg)](https://godoc.org/github.com/mauricioklein/go-spacetree)

Generate a tree structure from an indented text.

## Installation

```go
$ go get -u github.com/mauricioklein/go-spacetree/spacetree
```

## How to use

```go
  import (
    "os"
    "bufio"

    "github.com/mauricioklein/go-spacetree/spacetree"
  )

  // Generate a *bufio.Scanner with the indented content from a file
  file, err := os.Open("PUT THE FILEPATH HERE")

  if err != nil {
    panic(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  // Define the indentation symbol used on the file
  identationSymbol := "\t" // file is indented by tabs

  // Generate the tree
  // Returns a *spacetree.Node, holding the tree's root node
  root, err := spacetree.New(scanner, indentationSymbol)
```

Runnable examples can be found [here](examples/).

## Tests

```bash
$ go test -v ./...
```

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

## License

Spacetree is licensed under the [MIT](https://opensource.org/licenses/MIT) License.

