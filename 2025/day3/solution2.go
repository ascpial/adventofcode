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

func rowTotal(puzzle *Puzzle, y int) int {
	total := 0
	minIndex := 0
	for i := range 12 {
		maxValue := 0
		for x := minIndex; x < puzzle.Width-11+i; x++ {
			value := puzzle.get(x, y)
			if value > maxValue {
				maxValue = value
				minIndex = x + 1
			}
		}
		total = total*10 + maxValue
	}
	return total
}

func main() {
	data := input
	width := bytes.IndexByte(data, '\n')
	height := (len(data) + 1) / (width + 1)
	puzzle := &Puzzle{
		data,
		width,
		height,
	}
	sum := 0
	for y := range puzzle.Height {
		sum += rowTotal(puzzle, y)
	}
	fmt.Printf("%d\n", sum)
}
