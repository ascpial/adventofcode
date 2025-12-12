package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var example = `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2
`

//go:embed input.txt
var input string

func main() {
	puzzle := strings.Split(strings.TrimSpace(input), "\n\n")
	presents := puzzle[:6]
	maps := puzzle[6]

	presentsSizes := make([]int, 6)
	for _, present := range presents {
		parts := strings.Split(present, ":\n")
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		size := 0
		for _, c := range parts[1] {
			if c == '#' {
				size++
			}
		}
		presentsSizes[id] = size
	}

	total := 0
	for rawArea := range strings.SplitSeq(maps, "\n") {
		parts := strings.Split(rawArea, ": ")
		rawCoords := strings.Split(parts[0], "x")
		width, err := strconv.Atoi(rawCoords[0])
		if err != nil {
			panic(err)
		}
		height, err := strconv.Atoi(rawCoords[1])
		if err != nil {
			panic(err)
		}
		numberOfThings := 0
		totalArea := 0
		for i, rawNumber := range strings.Split(parts[1], " ") {
			number, err := strconv.Atoi(rawNumber)
			if err != nil {
				panic(err)
			}
			totalArea += number * presentsSizes[i]
			numberOfThings += number
		}
		// fmt.Printf("%d\n", width*height-numberOfThings*9)
		if width*height-numberOfThings*9 < 0 && width*height-totalArea >= 0 {
			panic(fmt.Sprintf("%s not implemented error", rawArea))
		}
		if width*height-totalArea > 0 {
			total++
		}
	}
	fmt.Printf("%d\n", total)
}
