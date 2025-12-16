package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"regexp"
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
	Buttons  [][]int
	Joltages []int
}

func ParseMachine(rawMachine string) Machine {
	parts := strings.Split(rawMachine, " ")
	rawTarget := parts[0]
	rawTarget = strings.TrimSuffix(strings.TrimPrefix(rawTarget, "["), "]")

	buttons := [][]int{}
	for i := 1; i < len(parts)-1; i++ {
		part := strings.TrimPrefix(strings.TrimSuffix(parts[i], ")"), "(")
		things := []int{}
		for rawLight := range strings.SplitSeq(part, ",") {
			light, err := strconv.Atoi(rawLight)
			if err != nil {
				panic(err)
			}
			things = append(things, light)
		}
		buttons = append(buttons, things)
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

	return Machine{buttons, requirements}
}

func InlineSum(values []int) string {
	if len(values) == 1 {
		return fmt.Sprintf("b%d", values[0])
	} else {
		return fmt.Sprintf("(+ b%d %s)", values[0], InlineSum(values[1:]))
	}
}

var pattern = regexp.MustCompile(` (\d+)\)`)

func Solve(machine Machine) int {
	query := ""
	for button := range len(machine.Buttons) {
		query += fmt.Sprintf("(declare-const b%d Int)\n", button)
	}
	allJoltages := make([][]int, len(machine.Joltages))
	allButtons := []int{}
	for buttonID, button := range machine.Buttons {
		for _, joltageID := range button {
			allJoltages[joltageID] = append(allJoltages[joltageID], buttonID)
		}
		query += fmt.Sprintf("(assert (>= b%d 0))\n", buttonID)
		allButtons = append(allButtons, buttonID)
	}
	for i, joltage := range machine.Joltages {
		query += fmt.Sprintf("(assert (= %s %d))\n", InlineSum(allJoltages[i]), joltage)
	}
	query += fmt.Sprintf("(minimize %s)\n", InlineSum(allButtons))
	query += "(check-sat)\n(get-model)\n"
	// fmt.Print(query)
	fi, err := os.Create("solution2.z3")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	fi.Write([]byte(query))
	cmd := exec.Command("z3", "solution2.z3")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	allValues := pattern.FindAllStringSubmatch(string(stdout), -1)

	total := 0
	for _, values := range allValues {
		value, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		total += value
	}

	return total
}

func main() {
	machines := []Machine{}
	for rawMachine := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		machines = append(machines, ParseMachine(rawMachine))
	}

	minPresses := 0
	for _, machine := range machines {
		minPresses += Solve(machine)
	}
	fmt.Printf("%d\n", minPresses)
}
