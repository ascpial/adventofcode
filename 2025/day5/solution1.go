package main

import (
	"crypto/des"
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

type SegtreeNode struct {
	Start    int
	End      int
	Segments []Range
	Left     *SegtreeNode
	Right    *SegtreeNode
}

func (n *SegtreeNode) Leaf() bool {
	return n.Left == nil && n.Right == nil
}

func main() {
	dataAreas := strings.Split(strings.TrimSpace(example), "\n\n")
	if len(dataAreas) != 2 {
		panic("Data input unreadable")
	}
	rawRanges := strings.Split(dataAreas[0], "\n")
	rawIngredients := strings.Split(dataAreas[1], "\n")

	ranges := make([]Range, len(rawRanges))
	for i, rawRange := range rawRanges {
		ranges[i] = parseRange(rawRange)
	}

	pointsMap := map[int]struct{}{}
	for _, range_ := range ranges {
		pointsMap[range_.Start] = struct{}{}
		pointsMap[range_.End] = struct{}{}
	}
	points := make([]int, len(pointsMap))
	i := 0
	for point := range pointsMap {
		points[i] = point
		i++
	}
	slices.Sort(points)

	ingredients := make([]int, len(rawIngredients))
	for i, rawIngredient := range rawIngredients {
		ingredient, err := strconv.Atoi(rawIngredient)
		if err != nil {
			panic(err)
		}
		ingredients[i] = ingredient
	}

	fmt.Printf("%v\n", ranges)
}
