package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var example = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
`

//go:embed input.txt
var input string

type Point struct {
	X int
	Y int
	Z int
}

func (p1 Point) Dist(p2 Point) int {
	return (p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y) + (p1.Z-p2.Z)*(p1.Z-p2.Z)
}

type Pair struct {
	P1 Point
	P2 Point
	D  int
}

func parsePoint(rawPoint string) Point {
	coords := strings.Split(rawPoint, ",")
	if len(coords) != 3 {
		panic("whhhhhat?")
	}
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		panic(err)
	}
	z, err := strconv.Atoi(coords[2])
	if err != nil {
		panic(err)
	}
	return Point{x, y, z}
}

func main() {
	rawPuzzle := strings.Split(strings.TrimSpace(input), "\n")
	puzzle := []Point{}
	for _, rawPoint := range rawPuzzle {
		puzzle = append(puzzle, parsePoint(rawPoint))
	}

	lengths := []Pair{}
	for i := range len(puzzle) {
		for j := i + 1; j < len(puzzle); j++ {
			lengths = append(lengths, Pair{puzzle[i], puzzle[j], puzzle[i].Dist(puzzle[j])})
		}
	}
	slices.SortFunc(lengths, func(A, B Pair) int { return A.D - B.D })
	// fmt.Printf("%v\n", lengths)

	circuits := map[int][]Point{}
	circuitsIndex := map[Point]int{}
	for i, p := range puzzle {
		circuits[i] = []Point{p}
		circuitsIndex[p] = i
	}

	i := 0
	var lastPair Pair
	for len(circuits) > 1 {
		pair := lengths[i]
		lastPair = pair
		// fmt.Printf("%v\n", pair)
		newCircuit := circuitsIndex[pair.P1]
		oldCircuit := circuitsIndex[pair.P2]
		if newCircuit != oldCircuit {
			for _, p := range circuits[oldCircuit] {
				circuitsIndex[p] = newCircuit
			}
			circuits[newCircuit] = slices.Concat(circuits[newCircuit], circuits[oldCircuit])
			delete(circuits, oldCircuit)
		}
		i++
		// fmt.Printf("%d, %d\n", len(circuits[newCircuit]), len(circuits))
	}

	fmt.Printf("%d\n", lastPair.P1.X*lastPair.P2.X)
}
