package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var example = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

//go:embed input.txt
var input string

var puzzle = input

type Direction int
type CellType int

const (
	Left Direction = iota
	Right
	Down
	Up
	Empty     CellType = iota
	Mirror1            // /
	Mirror2            // \
	SplitterH          // -
	SplitterV          // |
)

type Cell struct {
	Type      CellType
	Energized bool
	Seen      map[Direction]bool
}

var Mirror1Dirs = map[Direction]Direction{
	Left:  Down,
	Right: Up,
	Down:  Left,
	Up:    Right,
}
var Mirror2Dirs = map[Direction]Direction{
	Left:  Up,
	Right: Down,
	Down:  Right,
	Up:    Left,
}

func (c *Cell) GoThrough(incoming Direction) []Direction {
	if !c.Seen[incoming] {
		c.Energized = true
		c.Seen[incoming] = true
		switch c.Type {
		case Empty:
			return []Direction{incoming}
		case Mirror1:
			return []Direction{Mirror1Dirs[incoming]}
		case Mirror2:
			return []Direction{Mirror2Dirs[incoming]}
		case SplitterH:
			switch incoming {
			case Left, Right:
				return []Direction{incoming}
			case Up, Down:
				return []Direction{Left, Right}
			}
		case SplitterV:
			switch incoming {
			case Up, Down:
				return []Direction{incoming}
			case Left, Right:
				return []Direction{Up, Down}
			}
		default:
			panic(fmt.Sprintf("Unrecognized type: %d", c.Type))
		}
	}
	return []Direction{}
}

func MakeCell(t CellType) *Cell {
	return &Cell{
		t,
		false,
		make(map[Direction]bool, 4),
	}
}

var contraption [][]*Cell

type Traverse struct {
	X int
	Y int
	D Direction
}

func Count(entrypoint Traverse) int {
	for _, row := range contraption {
		for _, cell := range row {
			cell.Energized = false
			cell.Seen = make(map[Direction]bool, 4)
		}
	}
	height := len(contraption)
	width := len(contraption[0])
	queue := []Traverse{entrypoint}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, dir := range contraption[cur.Y][cur.X].GoThrough(cur.D) {
			switch dir {
			case Left:
				if cur.X > 0 {
					queue = append(queue, Traverse{cur.X - 1, cur.Y, dir})
				}
			case Right:
				if cur.X < width-1 {
					queue = append(queue, Traverse{cur.X + 1, cur.Y, dir})
				}
			case Up:
				if cur.Y > 0 {
					queue = append(queue, Traverse{cur.X, cur.Y - 1, dir})
				}
			case Down:
				if cur.Y < height-1 {
					queue = append(queue, Traverse{cur.X, cur.Y + 1, dir})
				}
			}
		}
	}
	counter := 0
	for _, row := range contraption {
		for _, cell := range row {
			if cell.Energized {
				counter++
				// 	fmt.Print("#")
				// } else {
				// 	fmt.Print(".")
			}
		}
		// fmt.Print("\n")
	}
	return counter
}

func Star1() {
	fmt.Printf("%d\n", Count(Traverse{0, 0, Right}))
}

func Star2() {
	height := len(contraption)
	width := len(contraption[0])
	entrypoints := []Traverse{}
	for x := range width {
		entrypoints = append(entrypoints, Traverse{x, 0, Down})
		entrypoints = append(entrypoints, Traverse{x, height - 1, Up})
	}
	for y := range height {
		entrypoints = append(entrypoints, Traverse{0, y, Right})
		entrypoints = append(entrypoints, Traverse{width - 1, y, Left})
	}
	maxCount := 0
	for _, entrypoint := range entrypoints {
		count := Count(entrypoint)
		if count > maxCount {
			maxCount = count
		}
	}
	fmt.Printf("%d\n", maxCount)
}

func main() {
	rows := strings.Split(strings.TrimSpace(puzzle), "\n")
	height := len(rows)
	width := len(rows[0])
	contraption = make([][]*Cell, height)
	for y, line := range rows {
		for x := range width {
			switch line[x] {
			case '.':
				contraption[y] = append(contraption[y], MakeCell(Empty))
			case '/':
				contraption[y] = append(contraption[y], MakeCell(Mirror1))
			case '\\':
				contraption[y] = append(contraption[y], MakeCell(Mirror2))
			case '-':
				contraption[y] = append(contraption[y], MakeCell(SplitterH))
			case '|':
				contraption[y] = append(contraption[y], MakeCell(SplitterV))
			default:
				panic(fmt.Sprintf("Unrecognized character: %b", line[x]))
			}
		}
	}
	Star2()
	Star1()
}
