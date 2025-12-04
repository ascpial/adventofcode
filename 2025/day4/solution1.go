package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

var example = []byte(`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`)

//go:embed input.txt
var input []byte

type Puzzle struct {
	Data   []byte
	Width  int
	Height int
}

func (p *Puzzle) get(x, y int) bool {
	if x < 0 || x >= p.Width || y < 0 || y >= p.Height {
		return false
	}
	return p.Data[y*(p.Width+1)+x] == 64
}

func createPuzzle(data []byte) *Puzzle {
	width := bytes.IndexByte(data, '\n')
	height := (len(data) + 1) / (width + 1)
	return &Puzzle{
		data,
		width,
		height,
	}
}

func Star1(puzzle *Puzzle) int {
	accessible := 0
	for y := range puzzle.Height {
		for x := range puzzle.Width {
			neighbours := -1
			if puzzle.get(x, y) {
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						if puzzle.get(x+i, y+j) {
							neighbours++
						}
					}
				}
				if neighbours <= 3 {
					accessible++
				}
			}
		}
	}
	return accessible
}

func main() {
	puzzle := createPuzzle(input)
	fmt.Printf("%d\n", Star1(puzzle))
}
