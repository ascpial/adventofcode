package main

import (
	_ "embed"
	"fmt"
	"strings"
)

// try to push

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
	Star1()
	Star2()
}
