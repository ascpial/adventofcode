package main

import (
	"fmt"
)

func Count(entrypoint Traverse) int {
	for _, row := range contraption {
		for _, cell := range row {
			cell.Energized = false
			cell.Seen = make(map[Direction]bool, 4)
		}
	}
	height := len(contraption)
	width := len(contraption[0])
	queue := []Traverse{entrypoint}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, dir := range contraption[cur.Y][cur.X].GoThrough(cur.D) {
			switch dir {
			case Left:
				if cur.X > 0 {
					queue = append(queue, Traverse{cur.X - 1, cur.Y, dir})
				}
			case Right:
				if cur.X < width-1 {
					queue = append(queue, Traverse{cur.X + 1, cur.Y, dir})
				}
			case Up:
				if cur.Y > 0 {
					queue = append(queue, Traverse{cur.X, cur.Y - 1, dir})
				}
			case Down:
				if cur.Y < height-1 {
					queue = append(queue, Traverse{cur.X, cur.Y + 1, dir})
				}
			}
		}
	}
	counter := 0
	for _, row := range contraption {
		for _, cell := range row {
			if cell.Energized {
				counter++
				// 	fmt.Print("#")
				// } else {
				// 	fmt.Print(".")
			}
		}
		// fmt.Print("\n")
	}
	return counter
}

func Star1() {
	fmt.Printf("%d\n", Count(Traverse{0, 0, Right}))
}
