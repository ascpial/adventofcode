package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

var example = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

//go:embed input.txt
var input string

type Run struct {
	red   uint
	green uint
	blue  uint
}

func max(a uint, b uint) uint {
	if a > b {
		return a
	} else {
		return b
	}
}

func (r Run) min(o Run) Run {
	return Run{
		max(r.red, o.red),
		max(r.green, o.green),
		max(r.blue, o.blue),
	}
}

type Game struct {
	ID   int
	runs []Run
}

func (g Game) power() uint {
	minRun := Run{0, 0, 0}
	for _, run := range g.runs {
		minRun = run.min(minRun)
	}
	return minRun.red * minRun.green * minRun.blue
}

func parseRun(rawRun string) Run {
	run := Run{}
	for cubes := range strings.SplitSeq(rawRun, ", ") {
		data := strings.Split(cubes, " ")
		n, _ := strconv.ParseInt(data[0], 10, 0)
		switch data[1] {
		case "red":
			run.red += uint(n)
		case "green":
			run.green += uint(n)
		case "blue":
			run.blue += uint(n)
		default:
			fmt.Printf("Not correct: %#v, %#v\n", data[1], data)
		}
	}
	return run
}

func parseGame(s string) Game {
	s = strings.TrimPrefix(s, "Game ")
	data := strings.Split(s, ": ")
	id, err := strconv.ParseInt(data[0], 10, 0)
	if err != nil {
		panic(err)
	}
	cubesData := strings.Split(data[1], "; ")
	runs := []Run{}
	for _, run := range cubesData {
		runs = append(runs, parseRun(run))
	}
	return Game{
		int(id),
		runs,
	}
}

func main() {
	puzzle := strings.Trim(input, "\n")

	games := []Game{}
	for rawGame := range strings.SplitSeq(puzzle, "\n") {
		games = append(games, parseGame(rawGame))
	}

	var powers uint = 0
	for _, game := range games {
		powers += game.power()
	}
	fmt.Printf("Total powers: %d\n", powers)
}
