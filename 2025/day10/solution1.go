package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var example = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`

//go:embed input.txt
var input string

type Machine struct {
	Target   int
	Buttons  []int
	Joltages []int
}

func ParseMachine(rawMachine string) Machine {
	parts := strings.Split(rawMachine, " ")
	rawTarget := parts[0]
	rawTarget = strings.TrimSuffix(strings.TrimPrefix(rawTarget, "["), "]")
	target := 0
	for i := len(rawTarget) - 1; i >= 0; i-- {
		target <<= 1
		if rawTarget[i] == '#' {
			target++
		}
	}

	buttons := []int{}
	for i := 1; i < len(parts)-1; i++ {
		part := strings.TrimPrefix(strings.TrimSuffix(parts[i], ")"), "(")
		button := 0
		for rawLight := range strings.SplitSeq(part, ",") {
			light, err := strconv.Atoi(rawLight)
			if err != nil {
				panic(err)
			}
			button += 1 << light
		}
		buttons = append(buttons, button)
	}

	return Machine{target, buttons, []int{}}
}

type Move struct {
	Current int
	Presses int
}

func main() {
	machines := []Machine{}
	for rawMachine := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		machines = append(machines, ParseMachine(rawMachine))
	}
	// fmt.Printf("%v\n", machines)

	minPresses := 0
	for _, machine := range machines {
		queue := []Move{{0, 0}}
		seens := map[int]struct{}{}
		found := false
		for len(queue) > 0 && !found {
			current := queue[0]
			queue = queue[1:]
			for _, button := range machine.Buttons {
				next := current.Current ^ button
				if next == machine.Target {
					found = true
					minPresses += current.Presses + 1
				}
				_, seen := seens[next]
				if !seen {
					queue = append(queue, Move{next, current.Presses + 1})
				}
			}
		}
	}
	fmt.Printf("%d\n", minPresses)
}
