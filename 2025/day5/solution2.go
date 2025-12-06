package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

//go:embed input.txt
var input string

type Range struct {
	Start int
	End   int
}

func (r Range) In(v int) bool {
	return r.Start <= v && v <= r.End
}

func parseRange(rawRange string) Range {
	ends := strings.Split(rawRange, "-")
	if len(ends) != 2 {
		panic("Unreadable range")
	}
	start, err := strconv.Atoi(ends[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(ends[1])
	if err != nil {
		panic(err)
	}
	return Range{start, end}
}

type Point struct {
	Value int
	end   bool
}

func main() {
	dataAreas := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(dataAreas) != 2 {
		panic("Data input unreadable")
	}
	rawRanges := strings.Split(dataAreas[0], "\n")

	ranges := make([]Range, len(rawRanges))
	for i, rawRange := range rawRanges {
		ranges[i] = parseRange(rawRange)
	}

	points := []Point{}
	for _, range_ := range ranges {
		points = append(points, Point{range_.Start - 1, false})
		points = append(points, Point{range_.End, true})
	}
	slices.SortFunc(points, func(a, b Point) int { return a.Value - b.Value })

	// fmt.Printf("%v\n", points)

	counter := 0
	totalLength := 0
	// last := -1
	for i := range len(points) {
		// fmt.Printf("point: %d, count: %d\n", points[i].Value, counter)
		if counter > 0 {
			totalLength += points[i].Value - points[i-1].Value
		}
		if points[i].end {
			counter--
		} else {
			// if points[i].Value != last && counter == 0 {
			// 	totalLength++
			// }
			counter++
		}
	}
	fmt.Printf("%v\n", totalLength)
}
