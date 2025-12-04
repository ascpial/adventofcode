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
	Data       []bool
	Neighbours []int
	Width      int
	Height     int
}

func (p *Puzzle) InitNeighbours() {
	for y := range p.Height {
		for x := range p.Width {
			neighbours := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if p.Get(x+i, y+j) {
						neighbours++
					}
				}
			}
			if p.Get(x, y) {
				neighbours--
			}
			p.Neighbours[y*p.Width+x] = neighbours
		}
	}
}

func (p *Puzzle) Get(x, y int) bool {
	if x < 0 || x >= p.Width || y < 0 || y >= p.Height {
		return false
	}
	return p.Data[y*p.Width+x]
}

func (p *Puzzle) Empty(x, y int) {
	if !(x < 0 || x >= p.Width || y < 0 || y >= p.Height) {
		p.Data[y*p.Width+x] = false
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				xn := x + i
				yn := y + j
				if !(xn < 0 || xn >= p.Width || yn < 0 || yn >= p.Height) {
					p.Neighbours[yn*p.Width+xn]--
				}
			}
		}
	}
}

func (p *Puzzle) CanRemove(x, y int) bool {
	if p.Get(x, y) {
		return p.Neighbours[y*p.Width+x] < 4
	}
	return false
}

func (p *Puzzle) Remove(x, y int) int {
	removed := 0
	if p.CanRemove(x, y) {
		removed += 1
		p.Empty(x, y)
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				removed += p.Remove(x+i, y+j)
			}
		}
	}
	return removed
}

func (p *Puzzle) Show() {
	for y := range p.Height {
		for x := range p.Width {
			if p.Get(x, y) {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func CreatePuzzle(data []byte) *Puzzle {
	width := bytes.IndexByte(data, '\n')
	height := (len(data) + 1) / (width + 1)
	newData := make([]bool, width*height)
	for y := range height {
		for x := range width {
			newData[y*width+x] = data[y*(width+1)+x] == 64
		}
	}
	puzzle := &Puzzle{
		newData,
		make([]int, width*height),
		width,
		height,
	}
	puzzle.InitNeighbours()
	return puzzle
}

func Star2(puzzle *Puzzle) int {
	removed := 0
	for y := range puzzle.Height {
		for x := range puzzle.Width {
			removed += puzzle.Remove(x, y)
			// puzzle.Show()
			// fmt.Printf("\n")
		}
	}
	return removed
}

func main() {
	puzzle := CreatePuzzle(input)
	fmt.Printf("%d\n", Star2(puzzle))
}
