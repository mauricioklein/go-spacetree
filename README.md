# Space Tree

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
  root := spacetree.New(scanner, indentationSymbol)
```

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

