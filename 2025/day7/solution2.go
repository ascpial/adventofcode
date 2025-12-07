package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var example = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`

//go:embed input.txt
var input string

type Pos struct {
	X int
	Y int
}

func explore(puzzle []string, paths map[Pos]int, pos Pos) int {
	if pos.Y >= len(puzzle) {
		return 0
	}
	counter, ok := paths[pos]
	if !ok {
		if puzzle[pos.Y][pos.X] == '^' {
			counter = explore(puzzle, paths, Pos{pos.X - 1, pos.Y + 1}) + explore(puzzle, paths, Pos{pos.X + 1, pos.Y + 1}) + 1
		} else {
			counter = explore(puzzle, paths, Pos{pos.X, pos.Y + 1})
		}
		paths[pos] = counter
	}
	return counter
}

func main() {
	puzzle := strings.Split(strings.TrimSpace(input), "\n")
	fmt.Printf("%d\n", explore(puzzle, map[Pos]int{}, Pos{strings.Index(puzzle[0], "S"), 1})+1)
}
