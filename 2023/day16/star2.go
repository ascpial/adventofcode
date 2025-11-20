package main

import "fmt"

func Star2() {
	height := len(contraption)
	width := len(contraption[0])
	entrypoints := []Traverse{}
	for x := range width {
		entrypoints = append(entrypoints, Traverse{x, 0, Down})
		entrypoints = append(entrypoints, Traverse{x, height - 1, Up})
	}
	for y := range height {
		entrypoints = append(entrypoints, Traverse{0, y, Right})
		entrypoints = append(entrypoints, Traverse{width - 1, y, Left})
	}
	maxCount := 0
	for _, entrypoint := range entrypoints {
		count := Count(entrypoint)
		if count > maxCount {
			maxCount = count
		}
	}
	fmt.Printf("%d\n", maxCount)
}
