package main

import (
	_ "embed"
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

func main() {
	Star1()
}
