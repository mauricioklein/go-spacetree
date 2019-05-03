package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mauricioklein/go-spacetree/spacetree"
)

func main() {
	file, err := os.Open("../test_input/multiple_levels_with_spaces.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	indentationSymbol := "  " // File is indented by two consecutive spaces

	root := spacetree.New(scanner, indentationSymbol)

	fmt.Printf("Root: %+v\n", root)
}
