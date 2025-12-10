package main

import (
	"container/heap"
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

	return Machine{0, buttons, requirements}
}

type Move struct {
	Current   []int
	Presses   int
	Estimated int
}

// Define a custom type for the heap
type PriorityQueue []Move

// Implement the heap.Interface for IntHeap

// Len is the number of elements in the collection.
func (h PriorityQueue) Len() int { return len(h) }

// Less reports whether the element with index i should sort before the element with index j.
func (h PriorityQueue) Less(i, j int) bool {
	return h[i].Estimated < h[j].Estimated // Min-heap: smallest element should be at the top
}

// Swap swaps the elements with indexes i and j.
func (h PriorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap.
func (h *PriorityQueue) Push(x any) {
	*h = append(*h, x.(Move))
}

// Pop removes and returns the smallest element from the heap.
func (h *PriorityQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func HashStatus(status []int) string {
	// fmt.Printf("%v\n", status)
	// hashed := []string{}
	// for _, a := range status {
	// 	hashed = append(hashed, strconv.FormatInt(int64(a), 10))
	// }
	return fmt.Sprintf("%v", status)
}

func MinDist(current, target []int) int {
	minDist := ^int(0)
	for i := range len(current) {
		if target[i]-current[i] < minDist {
			minDist = target[i] - current[i]
		}
	}
	return minDist
}

func Compatible(current, target []int) bool {
	for i := range len(current) {
		if current[i] > target[i] {
			return false
		}
	}
	return true
}

func Add(current []int, things []int) []int {
	current2 := make([]int, len(current))
	copy(current2, current)
	for _, thing := range things {
		current2[thing]++
	}
	return current2
}

func AStar(machine Machine) int {
	// fmt.Printf("Target: %v\nButtons: %v\n", machine.Joltages, machine.Buttons)
	visited := map[string]struct{}{}

	pq := &PriorityQueue{}
	target := HashStatus(machine.Joltages)

	heap.Init(pq)
	start := make([]int, len(machine.Joltages))
	heap.Push(pq, Move{start, 0, MinDist(start, machine.Joltages)})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Move)
		// fmt.Printf("%v\n", cur)

		for _, things := range machine.Buttons {
			next := Add(cur.Current, things)
			hashedNext := HashStatus(next)
			if hashedNext == target {
				return cur.Presses + 1
			}
			if Compatible(next, machine.Joltages) {
				// fmt.Printf("%s\n", hashedNext)
				_, seen := visited[hashedNext]
				if !seen {
					heap.Push(pq, Move{next, cur.Presses + 1, cur.Presses + 1 + MinDist(next, machine.Joltages)})
				}
			}
		}
	}

	panic("not found what is going on aaaaaaaah")
}

func main() {
	machines := []Machine{}
	for rawMachine := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		machines = append(machines, ParseMachine(rawMachine))
	}

	minPresses := 0
	for _, machine := range machines {
		minPresses += AStar(machine)
	}
	fmt.Printf("%d\n", minPresses)
}
