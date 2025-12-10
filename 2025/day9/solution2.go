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

func Show(world map[Point]int8, x0, y0, offset int) {
	for y := y0; y < y0+offset; y++ {
		for x := x0; x < x0+offset; x++ {
			v, ok := world[Point{x, y}]
			if ok {
				if v == 0 {
					fmt.Print("#")
				} else if v == -1 {
					fmt.Print("-")
				} else {
					fmt.Print("+")
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Print("\n")
	}
}

func OptionalySet(world map[Point]int8, pos Point, v int8) {
	v2, ok := world[pos]
	if !ok || v2 != 0 {
		world[pos] = v
	}
}

func main() {
	rawPoints := strings.Split(strings.TrimSpace(example), "\n")
	points := []Point{}
	for _, rawPoint := range rawPoints {
		points = append(points, parsePoint(rawPoint))
	}
	world := map[Point]int8{}
	previous := points[len(points)-1]
	for _, point := range points {
		if point.X == previous.X {
			for y := min(point.Y, previous.Y); y <= max(point.Y, previous.Y); y++ {
				world[Point{point.X, y}] = 0
				if point.Y < previous.Y {
					OptionalySet(world, Point{point.X - 1, y}, -1)
					OptionalySet(world, Point{point.X + 1, y}, 1)
				} else {
					OptionalySet(world, Point{point.X - 1, y}, 1)
					OptionalySet(world, Point{point.X + 1, y}, -1)
				}
			}
		} else {
			for x := min(point.X, previous.X); x <= max(point.X, previous.X); x++ {
				world[Point{x, point.Y}] = 0
				if point.X < previous.X {
					OptionalySet(world, Point{x, point.Y - 1}, 1)
					OptionalySet(world, Point{x, point.Y + 1}, -1)
				} else {
					OptionalySet(world, Point{x, point.Y - 1}, -1)
					OptionalySet(world, Point{x, point.Y + 1}, 1)
				}
			}
		}
		previous = point
	}
	// Show(world, 0, 0, 16)
	// looking for the external value

	pos := Point{0, points[0].Y}
	var outerDirection int8 = 0
	for outerDirection == 0 {
		v, ok := world[pos]
		if ok {
			outerDirection = v
		}
		pos.X++
	}

	maxArea := 0
	// counter := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			// counter++
			p1 := points[i]
			p2 := points[j]
			x0 := min(p1.X, p2.X)
			x1 := max(p1.X, p2.X)
			y0 := min(p1.Y, p2.Y)
			y1 := max(p1.Y, p2.Y)
			if p1.Area(p2) > maxArea {
				valid := true
				for x := x0; x <= x1 && valid; x++ {
					v, ok := world[Point{x, y0 + 1}]
					valid = !ok || v != outerDirection
					v, ok = world[Point{x, y1 - 1}]
					valid = valid && (!ok || v != outerDirection)
				}
				for y := y0; y <= y1 && valid; y++ {
					v, ok := world[Point{x0 + 1, y}]
					valid = !ok || v != outerDirection
					v, ok = world[Point{x1 - 1, y}]
					valid = valid && (!ok || v != outerDirection)
				}
				if valid {
					maxArea = p1.Area(p2)
				}
			}
			// if counter%100 == 0 {
			// 	fmt.Printf("Progress: %d/%d; current max: %d\n", counter, len(points)*len(points)/2, maxArea)
			// }
		}
	}
	fmt.Printf("%d\n", maxArea)
}
