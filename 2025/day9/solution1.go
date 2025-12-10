package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var example = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

//go:embed input.txt
var input string

type Point struct {
	X int
	Y int
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func (p1 Point) Area(p2 Point) int {
	return (Abs(p1.X-p2.X) + 1) * (Abs(p1.Y-p2.Y) + 1)
}

func parsePoint(rawPoint string) Point {
	coords := strings.Split(rawPoint, ",")
	if len(coords) != 2 {
		panic("blblblbll")
	}
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		panic(err)
	}
	return Point{x, y}
}

func main() {
	rawPoints := strings.Split(strings.TrimSpace(input), "\n")
	points := []Point{}
	for _, rawPoint := range rawPoints {
		points = append(points, parsePoint(rawPoint))
	}
	maxArea := 0
	for i := range len(points) {
		for j := i + 1; j < len(points); j++ {
			curArea := points[i].Area(points[j])
			if curArea > maxArea {
				maxArea = curArea
			}
		}
	}
	fmt.Printf("%d\n", maxArea)
}
