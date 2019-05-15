/*
Package spacetree provides an easy and convenient way to generate a tree based on an indented content.

Basics:

Spacetree generates a tree representation of indented content. For example:

 1
   1.1
   1.2
 2
   2.1

Would generate the tree:

 ROOT -
       |
        - 1 -
       |     |
       |      - 1.1
       |     |
       |      - 1.2
       |
        - 2 -
             |
              - 2.1

Usage:

Spacetree expects two attributes as input:
  - the content buffer (*bufio.Scanner), holding the indented content
  - the indentation symbol (string)

The result is:
  - a *spacetree.Node, representing the root of the tree
  - an error, set in case of a processing error (broken indentation, for example)

Example:

	import "github.com/mauricioklein/go-spacetree/spacetree"

	scanner := [your *bufio.Scanner containing the indented content]
	indentationSymbol := " " // single space

	root, err := spacetree.New(scanner, indentationSymbol)
*/
package spacetree
