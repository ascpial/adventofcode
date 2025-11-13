package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var example = `3   4
4   3
2   5
1   3
3   9
3   3`

//go:embed input.txt
var input string

func abs(n int) uint {
	if n < 0 {
		return uint(-n)
	} else {
		return uint(n)
	}
}

func main() {
	puzzle := strings.Trim(input, "\n")

	lefts := []int{}
	rights := []int{}

	for line := range strings.SplitSeq(puzzle, "\n") {
		values := strings.SplitN(line, "   ", 2)
		left, _ := strconv.ParseInt(values[0], 10, 0)
		right, _ := strconv.ParseInt(values[1], 10, 0)
		lefts = append(lefts, int(left))
		rights = append(rights, int(right))
	}

	slices.SortFunc(lefts, func(a int, b int) int { return b - a })
	slices.SortFunc(rights, func(a int, b int) int { return b - a })

	var similarity uint = 0
	for i := 0; i < len(lefts); i++ {
		similarity += abs(lefts[i] - rights[i])
	}

	fmt.Printf("Similarity: %d\n", similarity)
}
