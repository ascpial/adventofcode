package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"time"
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

func (p Puzzle) get(x, y int) int {
	return int(p.Data[y*(p.Width+1)+x] - '0')
}

func createPuzzle(data []byte) Puzzle {
	width := bytes.IndexByte(data, '\n')
	// fmt.Printf("buffer size: %d\n", len(data))
	height := (len(data) + 1) / (width + 1)
	// fmt.Printf("size: %d %d\n", width, height)
	return Puzzle{
		data,
		width,
		height,
	}
}

func Star1(puzzle Puzzle) int {
	sum := 0
	for y := range puzzle.Height {
		firstValue := 0
		firstIndex := 0
		for x := 0; x < puzzle.Width-1; x++ {
			// fmt.Printf("%d\n", x)
			if puzzle.get(x, y) > firstValue {
				firstValue = puzzle.get(x, y)
				firstIndex = x
			}
		}
		secondValue := 0
		// secondIndex := firstIndex
		for x := firstIndex + 1; x < puzzle.Width; x++ {
			if puzzle.get(x, y) > secondValue {
				secondValue = puzzle.get(x, y)
				// secondIndex = x
			}
		}
		// fmt.Printf("Line %d: %d\n", y, firstValue*10+secondValue)
		sum += firstValue*10 + secondValue
	}
	return sum
}

func Star2(puzzle Puzzle) int {
	sum := 0
	for y := range puzzle.Height {
		values := [12]int{}
		indexes := [12]int{}
		for i := range 12 {
			// fmt.Printf("%d\n", indexes[i])
			for x := indexes[i]; x < puzzle.Width-11+i; x++ {
				if puzzle.get(x, y) > values[i] {
					values[i] = puzzle.get(x, y)
					if i < 11 {
						indexes[i+1] = x + 1
					}
				}
			}
		}
		total := 0
		for _, n := range values {
			total = total*10 + n
		}
		// fmt.Printf("Line %d: %d\n", y, total)
		sum += total
	}
	return sum
}

func main() {
	puzzle := createPuzzle(input)
	start1 := time.Now()
	star1 := Star1(puzzle)
	fmt.Printf("First star execution time: %v\n", time.Since(start1))
	start2 := time.Now()
	star2 := Star2(puzzle)
	fmt.Printf("Second star execution time: %v\n", time.Since(start2))
	fmt.Printf("%d\n", star1)
	fmt.Printf("%d\n", star2)
}
