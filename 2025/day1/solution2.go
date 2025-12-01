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

func sign(v int) int {
	if v >= 0 {
		return 1
	} else {
		return -1
	}
}

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
		s := sign(delta)
		for _ = range delta * s {
			pos = (pos + s) % 100
			// fmt.Printf("Pos: %d\n", pos)
			if pos == 0 {
				counter += 1
				// fmt.Print("Points to zero\n")
			}
		}
	}
	fmt.Printf("%d\n", counter)
}
