package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strings"
)

type Direction uint

const (
	Left Direction = iota
	Right
	Down
	Up
)

var example = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

var example2 = `111111111111
999999999991
999999999991
999999999991
999999999991`

type Pos struct {
	X, Y int
}

func (p Pos) InBound(width int, height int) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < width && p.Y < height
}

func (p Pos) Up() Pos {
	return Pos{p.X, p.Y - 1}
}

func (p Pos) Down() Pos {
	return Pos{p.X, p.Y + 1}
}

func (p Pos) Left() Pos {
	return Pos{p.X - 1, p.Y}
}

func (p Pos) Right() Pos {
	return Pos{p.X + 1, p.Y}
}

//go:embed input.txt
var input string

type Status struct {
	pos       Pos
	distance  uint
	direction Direction
}

func (s Status) Next() []Status {
	d := []Status{}
	if s.distance < 10 {
		if s.direction == Left {
			d = append(d, Status{s.pos.Left(), s.distance + 1, s.direction})
		} else if s.direction == Right {
			d = append(d, Status{s.pos.Right(), s.distance + 1, s.direction})
		} else if s.direction == Up {
			d = append(d, Status{s.pos.Up(), s.distance + 1, s.direction})
		} else {
			d = append(d, Status{s.pos.Down(), s.distance + 1, s.direction})
		}
	}
	if (s.direction == Left || s.direction == Right) && s.distance >= 4 {
		d = append(d, Status{s.pos.Up(), 1, Up}, Status{s.pos.Down(), 1, Down})
	}
	if (s.direction == Up || s.direction == Down) && s.distance >= 4 {
		d = append(d, Status{s.pos.Left(), 1, Left}, Status{s.pos.Right(), 1, Right})
	}
	return d
}

type Visited struct {
	cost   uint
	source Status
}

// Define a custom type for the heap
type PriorityQueue []Visited

// Implement the heap.Interface for IntHeap

// Len is the number of elements in the collection.
func (h PriorityQueue) Len() int { return len(h) }

// Less reports whether the element with index i should sort before the element with index j.
func (h PriorityQueue) Less(i, j int) bool {
	return h[i].cost < h[j].cost // Min-heap: smallest element should be at the top
}

// Swap swaps the elements with indexes i and j.
func (h PriorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap.
func (h *PriorityQueue) Push(x any) {
	*h = append(*h, x.(Visited))
}

// Pop removes and returns the smallest element from the heap.
func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Star1(city map[Pos]uint, width int, height int) uint {
	visited := make(map[Status]Visited)

	pq := &PriorityQueue{}

	heap.Init(pq)
	heap.Push(pq, Visited{0, Status{Pos{0, 0}, 0, Right}})
	heap.Push(pq, Visited{0, Status{Pos{0, 0}, 0, Down}})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Visited)
		// fmt.Printf("%#v\n", cur.source.pos)

		for _, next := range cur.source.Next() {
			if next.pos.InBound(width, height) {
				previousNext, exists := visited[next]
				if !exists || previousNext.cost > cur.cost+city[next.pos] {
					visited[next] = Visited{
						cur.cost + city[next.pos],
						cur.source,
					}
					heap.Push(pq, Visited{
						cur.cost + city[next.pos],
						next,
					})
				}
			}
		}
	}

	minCost := ^uint(0)
	// var minEnd Status
	for _, dir := range []Direction{Left, Right, Up, Down} {
		for _, distance := range []int{4, 5, 6, 7, 8, 9, 10} {
			s, ok := visited[Status{Pos{width - 1, height - 1}, uint(distance), dir}]
			if ok {
				if s.cost < minCost {
					minCost = s.cost
					// minEnd = s.source
				}
			}
		}
	}

	// route := make(map[Pos]struct{})
	// route[Pos{width - 1, height - 1}] = struct{}{}
	//
	// for !(minEnd.pos.X == 0 && minEnd.pos.Y == 0) {
	// 	route[minEnd.pos] = struct{}{}
	// 	minEnd = visited[minEnd].source
	// }
	//
	// for y := range height {
	// 	for x := range width {
	// 		_, ok := route[Pos{x, y}]
	// 		if ok {
	// 			fmt.Print(".")
	// 		} else {
	// 			fmt.Printf("%d", city[Pos{x, y}])
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }

	return minCost
}

func Star2() uint {
	return 0
}

func main() {
	var puzzle = strings.TrimSpace(input)

	city := make(map[Pos]uint)
	lines := strings.Split(puzzle, "\n")
	height := len(lines)
	width := len(lines[0])
	for y, line := range lines {
		for x := range len(line) {
			city[Pos{x, y}] = uint(line[x] - 48)
		}
	}

	// fmt.Printf("%d, %d\n", width, height)

	fmt.Printf("%d\n", Star1(city, width, height))
	// fmt.Printf("%d\n", Star2())
}
