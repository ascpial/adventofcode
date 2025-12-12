package main

import (
	_ "embed"
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
	puzzle := strings.Split(strings.TrimSpace(example), "\n\n")
	presents := puzzle[:5]
	maps := puzzle[5]

	allPresents := map[int16]struct{}{}
	for _, present := presents {

	}
}
