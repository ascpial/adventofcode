package main

import (
	_ "embed"
	"fmt"
	// "os"
	"slices"
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

type Segment struct {
	p1 Point
	p2 Point
}

func (s Segment) Vertical() bool {
	return s.p1.X == s.p2.X
}

func raymarching(verticalSegments []Segment, p Point) bool {
	inside := false
	x := verticalSegments[0].p1.X
	i := 0
	for i < len(verticalSegments) && x < p.X {
		segment := verticalSegments[i]
		if segment.p1.Y <= p.Y && p.Y <= segment.p2.Y {
			inside = !inside
		}
		x = verticalSegments[i+1].p1.X
		i++
	}
	return inside
}

func main() {
	rawPoints := strings.Split(strings.TrimSpace(input), "\n")
	points := []Point{}
	for _, rawPoint := range rawPoints {
		points = append(points, parsePoint(rawPoint))
	}

	// 	template := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	// <!-- Created with Inkscape (http://www.inkscape.org/) -->
	//
	// <svg
	//    viewBox="0 0 %d %d"
	//    version="1.1"
	//    id="svg1"
	//    xmlns="http://www.w3.org/2000/svg"
	//    xmlns:svg="http://www.w3.org/2000/svg">
	//   <defs
	//      id="defs1" />
	//   <g
	//      id="layer1">
	//     <path
	//        style="fill:none;stroke:#000000;stroke-width:100;stroke-opacity:1"
	//        d="M %s Z"
	//        id="path2" />
	// 		<circle
	//        style="fill:#fd0303;fill-opacity:1;stroke:none;stroke-width:100;stroke-opacity:1"
	//        id="path3"
	//        cx="%d"
	//        cy="%d"
	//        r="300" />
	//     <circle
	//        style="fill:#fd0303;fill-opacity:1;stroke:none;stroke-width:1.09934;stroke-opacity:1"
	//        id="path3-3"
	//        cx="%d"
	//        cy="%d"
	//        r="300" />
	//   </g>
	// </svg>`
	//
	// 	coords := []string{}
	// 	maxX := 0
	// 	maxY := 0
	// 	for _, point := range points {
	// 		coords = append(coords, fmt.Sprintf("%d,%d", point.X, point.Y))
	// 		if point.X > maxX {
	// 			maxX = point.X
	// 		}
	// 		if point.Y > maxY {
	// 			maxY = point.Y
	// 		}
	// 	}
	// 	f, err := os.Create("debug.svg")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer f.Close()

	segments := []Segment{}
	previous := points[len(points)-1]
	for _, point := range points {
		segments = append(segments, Segment{Point{
			min(point.X, previous.X),
			min(point.Y, previous.Y),
		}, Point{
			max(point.X, previous.X),
			max(point.Y, previous.Y),
		}})
		previous = point
	}

	verticalSegments := []Segment{}
	for _, segment := range segments {
		if segment.Vertical() {
			verticalSegments = append(verticalSegments, segment)
		}
	}
	slices.SortFunc(verticalSegments, func(a, b Segment) int { return a.p1.X - b.p1.X })

	maxArea := 0
	// var maxPoints Segment
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pa := points[i]
			pb := points[j]
			if pa.Area(pb) > maxArea {
				x1 := min(pa.X, pb.X)
				x2 := max(pa.X, pb.X)
				y1 := min(pa.Y, pb.Y)
				y2 := max(pa.Y, pb.Y)

				// find if the rectangle can be in the inner area
				inside := raymarching(verticalSegments, Point{(x1 + x2) / 2, (y1 + y2) / 2})

				if inside { // this rectangle is canditate
					stillCanditate := true
					for k := 0; k < len(segments) && stillCanditate; k++ {
						segment := segments[k]
						stillCanditate = x2 <= segment.p1.X || segment.p2.X <= x1 || y2 <= segment.p1.Y || segment.p2.Y <= y1
					}
					if stillCanditate {
						maxArea = pa.Area(pb)
						// maxPoints = Segment{pa, pb}
					}
				}
			}
		}
	}
	fmt.Printf("%d\n", maxArea)
	// fmt.Fprintf(f, template, maxX, maxY, strings.Join(coords, " "), maxPoints.p1.X, maxPoints.p1.Y, maxPoints.p2.X, maxPoints.p2.Y)
}
