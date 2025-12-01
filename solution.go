package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var example = ``

//go:embed input.txt
var input string

func Star1() uint {
	return 0
}

func Star2() uint {
	return 0
}

func main() {
	var puzzle = strings.TrimSpace(example)

	fmt.Printf("%d\n", Star1())
	// fmt.Printf("%d\n", Star2())
}
