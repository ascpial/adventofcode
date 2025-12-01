package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var example = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

//go:embed input.txt
var input string

func main() {
	puzzle := strings.Replace(strings.TrimSpace(input), "R", "", -1)
	puzzle = strings.Replace(puzzle, "L", "-", -1)
	pos := 50
	counter := 0
	for line := range strings.SplitSeq(puzzle, "\n") {
		delta, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		pos = (pos + delta) % 100
		fmt.Printf("Pos: %d\n", pos)
		if pos == 0 {
			counter += 1
		}
	}
	fmt.Printf("%d\n", counter)
}
