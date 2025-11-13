package main

import (
	_ "embed"
	"fmt"
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
	// puzzle := example

	lefts := []int{}
	rights := []int{}

	for line := range strings.SplitSeq(puzzle, "\n") {
		values := strings.SplitN(line, "   ", 2)
		left, _ := strconv.ParseInt(values[0], 10, 0)
		right, _ := strconv.ParseInt(values[1], 10, 0)
		lefts = append(lefts, int(left))
		rights = append(rights, int(right))
	}

	occurences := make(map[int]int)
	for _, e := range rights {
		i, _ := occurences[e]
		occurences[e] = i + 1
	}

	var similarity uint = 0
	for i := 0; i < len(lefts); i++ {
		left := lefts[i]
		count, _ := occurences[left]
		similarity += uint(left * count)
	}

	fmt.Printf("Similarity: %d\n", similarity)
}
