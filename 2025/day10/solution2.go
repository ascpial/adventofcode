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

//go:embed input2.txt
var input2 string

type Machine struct {
	id       int
	Target   int
	Buttons  []int
	Joltages []int
}

func ParseMachine(id int, rawMachine string) Machine {
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

	rawRequirements := strings.Split(strings.TrimPrefix(strings.TrimSuffix(parts[len(parts)-1], "}"), "{"), ",")
	requirements := []int{}
	for _, rawRequirement := range rawRequirements {
		requirement, err := strconv.Atoi(rawRequirement)
		if err != nil {
			panic(err)
		}
		requirements = append(requirements, requirement)
	}

	return Machine{id, target, buttons, requirements}
}

type Target struct {
	machine int
	target  int
}
type PressState struct {
	presses        int
	configurations []int
}

var blinkTargets = map[Target]PressState{}

type ExploreState struct {
	current int
	presses int
	buttons int
}

func computeConfigurations(machine Machine, target int) PressState {
	states, found := blinkTargets[Target{machine.id, target}]
	if found {
		return states
	}
	if target == 0 {
		return PressState{0, []int{0}}
	}
	queue := []ExploreState{{0, 0, 0}}
	seens := map[int]struct{}{}
	foundFirst := false
	foundLast := false
	targetPresses := 0
	configs := []int{}
	for len(queue) > 0 && !foundLast {
		current := queue[0]
		queue = queue[1:]
		for buttonID, button := range machine.Buttons {
			if current.buttons>>buttonID&1 == 0 {
				next := current.current ^ button
				nextButtons := current.buttons | (1 << buttonID)
				if next == target {
					if !foundFirst {
						targetPresses = current.presses + 1
						configs = append(configs, nextButtons)
						foundFirst = true
					} else {
						if current.presses+1 == targetPresses {
							configs = append(configs, nextButtons)
						} else {
							foundLast = true
						}
					}
				} else {
					_, seen := seens[nextButtons]
					if !seen {
						queue = append(queue, ExploreState{next, current.presses + 1, nextButtons})
						seens[nextButtons] = struct{}{}
					}
				}
			}
		}
	}
	blinkTargets[Target{machine.id, target}] = PressState{targetPresses, configs}
	return blinkTargets[Target{machine.id, target}]
}

func computeModulo(joltages []int) int {
	modulo := 0
	for i, joltage := range joltages {
		if joltage%2 == 1 {
			modulo = modulo | 1<<i
		}
	}
	return modulo
}

func computeMinPresses(machine Machine, target []int) int {
	end := true
	for i := 0; i < len(target) && end; i++ {
		end = end && target[i] == 0
	}
	if end {
		return 0
	}
	moduloTarget := computeModulo(target)
	state := computeConfigurations(machine, moduloTarget)
	// fmt.Printf(" target: %v; modulo: %s; presses: %d\n", target, strconv.FormatInt(int64(moduloTarget), 2), state.presses)
	minPresses := int(^uint(0) >> 1)
	for _, configuration := range state.configurations {
		curTarget := make([]int, len(target))
		copy(curTarget, target)
		invalid := false
		for buttonID := 0; buttonID < len(machine.Buttons) && !invalid; buttonID++ {
			button := machine.Buttons[buttonID]
			if configuration>>buttonID&1 == 1 {
				for i := range len(curTarget) {
					if button>>i&1 == 1 {
						curTarget[i]--
						if curTarget[i] < 0 {
							invalid = true
						}
					}
				}
			}
		}
		if !invalid {
			// fmt.Printf("  config: %s; state: %v\n", strconv.FormatInt(int64(configuration), 2), curTarget)
			for i := range len(curTarget) {
				if curTarget[i]%2 != 0 {
					panic("cannot be divided by two, paniiiiiiiiiiiic!")
				}
				curTarget[i] = curTarget[i] / 2
			}
			// fmt.Printf("  (after division: %v)\n", curTarget)
			curPresses := computeMinPresses(machine, curTarget)*2 + state.presses
			if curPresses < minPresses {
				// fmt.Printf("  found better candidate: %d; target: %v\n", curPresses, target)
				minPresses = curPresses
			}
		}
	}
	return minPresses
}

func main() {
	machines := []Machine{}
	for i, rawMachine := range strings.Split(strings.TrimSpace(input), "\n") {
		machines = append(machines, ParseMachine(i, rawMachine))
	}

	minPresses := 0

	for _, machine := range machines {
		fmt.Printf("machine %d; buttons: [", machine.id)
		for _, button := range machine.Buttons {
			fmt.Printf("%s ", strconv.FormatInt(int64(button), 2))
		}
		fmt.Print("]\n")
		curMinPresses := computeMinPresses(machine, machine.Joltages)
		minPresses += curMinPresses
		fmt.Printf("machine %d; minPresses: %d\n", machine.id, curMinPresses)
	}

	fmt.Printf("%d\n", minPresses)
}
