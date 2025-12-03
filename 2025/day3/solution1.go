package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

var example = []byte(`987654321111111
811111111111119
234234234234278
818181911112111`)

//go:embed input.txt
var input []byte

type Puzzle struct {
	Data   []byte
	Width  int
	Height int
}

func (p *Puzzle) get(x, y int) int {
	return int(p.Data[y*(p.Width+1)+x] - '0')
}

func createPuzzle(data []byte) *Puzzle {
	width := bytes.IndexByte(data, '\n')
	// fmt.Printf("buffer size: %d\n", len(data))
	height := (len(data) + 1) / (width + 1)
	// fmt.Printf("size: %d %d\n", width, height)
	return &Puzzle{
		data,
		width,
		height,
	}
}

func Star1(puzzle *Puzzle) int {
	sum := 0
	for y := range puzzle.Height {
		firstValue := 0
		firstIndex := 0
		for x := 0; x < puzzle.Width-1; x++ {
			if puzzle.get(x, y) > firstValue {
				firstValue = puzzle.get(x, y)
				firstIndex = x
			}
		}
		secondValue := 0
		for x := firstIndex + 1; x < puzzle.Width; x++ {
			if puzzle.get(x, y) > secondValue {
				secondValue = puzzle.get(x, y)
				// secondIndex = x
			}
		}
		sum += firstValue*10 + secondValue
	}
	return sum
}

func main() {
	puzzle := createPuzzle(input)
	fmt.Printf("%d\n", Star1(puzzle))
}
