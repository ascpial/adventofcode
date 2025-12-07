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

func main() {
	puzzle := strings.Split(strings.TrimSpace(input), "\n")
	pos := map[int]struct{}{}
	pos[strings.Index(puzzle[0], "S")] = struct{}{}
	counter := 0
	for y := 1; y < len(puzzle); y++ {
		nextPos := map[int]struct{}{}
		for x := range pos {
			if puzzle[y][x] == '^' {
				nextPos[x-1] = struct{}{}
				nextPos[x+1] = struct{}{}
				counter++
			} else {
				nextPos[x] = struct{}{}
			}
		}
		pos = nextPos
	}
	fmt.Printf("%d\n", counter)
}
