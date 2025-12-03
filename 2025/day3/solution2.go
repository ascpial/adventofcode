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

func rowSum(puzzle *Puzzle, y int, result chan int) {
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
	result <- total
}

func main() {
	puzzle := createPuzzle(input)
	result := make(chan int, puzzle.Height)
	for y := range puzzle.Height {
		go rowSum(puzzle, y, result)
	}
	sum := 0
	counter := 0
	for val := range result {
		counter++
		sum += val
		if counter == puzzle.Height {
			close(result)
		}
	}
	fmt.Printf("%d\n", sum)
}
