package main

import (
	_ "embed"
	"fmt"
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

func main() {
	dataAreas := strings.Split(strings.TrimSpace(input), "\n\n")
	if len(dataAreas) != 2 {
		panic("Data input unreadable")
	}
	rawRanges := strings.Split(dataAreas[0], "\n")
	rawIngredients := strings.Split(dataAreas[1], "\n")

	ranges := make([]Range, len(rawRanges))
	for i, rawRange := range rawRanges {
		ranges[i] = parseRange(rawRange)
	}

	counter := 0
	for _, rawIngredient := range rawIngredients {
		ingredient, err := strconv.Atoi(rawIngredient)
		if err != nil {
			panic(err)
		}
		found := false
		for i := 0; i < len(ranges) && !found; i++ {
			found = ranges[i].In(ingredient)
		}
		if found {
			counter++
		}
	}

	fmt.Printf("%v\n", counter)
}
