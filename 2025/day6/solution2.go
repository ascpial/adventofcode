package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var example = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

//go:embed input.txt
var input string

func main() {
	puzzle := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	x := 0
	totaltotal := 0
	for x < len(puzzle[0]) {
		length := 0
		for x+length+1 < len(puzzle[0]) && puzzle[len(puzzle)-1][x+length+1] == ' ' {
			length++
		}
		if x+length+1 >= len(puzzle[0]) {
			length++
		}
		// fmt.Printf("Length: %d\n", length)
		var operator func(int, int) int
		var total int
		// fmt.Printf("Operator: %q\n", puzzle[len(puzzle)-1][x])
		if puzzle[len(puzzle)-1][x] == '*' {
			operator = func(a, b int) int { return a * b }
			total = 1
		} else {
			operator = func(a, b int) int { return a + b }
			total = 0
		}
		for x2 := range length {
			v := 0
			for y := range len(puzzle) - 1 {
				if puzzle[y][x+x2] != ' ' {
					// fmt.Printf("Cur char: %d\n", int(puzzle[y][x+x2]-'0'))
					v = 10*v + int(puzzle[y][x+x2]-'0')
				}
			}
			// fmt.Printf("%d\n", v)
			total = operator(total, v)
		}

		// fmt.Printf("total: %d\n", total)
		totaltotal += total
		x += length + 1
	}
	fmt.Printf("%d\n", totaltotal)
}
