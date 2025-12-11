package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var example = `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
`

type Node struct {
	ID          string
	RawChildren []string
	Children    []*Node
}

func ParseNode(rawNode string) Node {
	parts := strings.Split(rawNode, ": ")
	if len(parts) != 2 {
		panic("uh")
	}
	return Node{
		parts[0],
		strings.Split(parts[1], " "),
		[]*Node{},
	}
}

//go:embed input.txt
var input string

type PathData struct {
	Valids   int
	FFT      int
	DAC      int
	Invalids int
}

func main() {
	nodes := map[string]*Node{}
	for rawNode := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		node := ParseNode(rawNode)
		nodes[node.ID] = &node
	}
	nodes["out"] = &Node{"out", []string{}, []*Node{}}
	for _, node := range nodes {
		for _, childID := range node.RawChildren {
			child, ok := nodes[childID]
			if !ok {
				panic(fmt.Sprintf("node %s has no child %s!!", node.ID, childID))
			}
			node.Children = append(node.Children, child)
		}
	}
	// fmt.Printf("%v\n", nodes)

	paths := map[string]*PathData{}
	paths["svr"] = &PathData{0, 0, 0, 1}
	queue := []*Node{nodes["svr"]}
	for len(queue) > 0 {
		node := queue[0]
		// fmt.Printf("todo: %d; ID: %s\n", len(queue), node.ID)
		curData := paths[node.ID]
		if curData == nil {
			panic("paniiiiiiiic!!! bannana")
		}
		queue = queue[1:]

		if node.ID != "out" {
			for _, child := range node.Children {
				data := paths[child.ID]
				if data == nil {
					queue = append(queue, child)
					data = &PathData{0, 0, 0, 0}
					paths[child.ID] = data
				}
				if child.ID == "fft" {
					data.Valids += curData.Valids + curData.DAC
					data.FFT += curData.Invalids + curData.FFT
				} else if child.ID == "dac" {
					data.Valids += curData.Valids + curData.FFT
					data.DAC += curData.Invalids + curData.DAC
				} else {
					data.Valids += curData.Valids
					data.FFT += curData.FFT
					data.DAC += curData.DAC
					data.Invalids += curData.Invalids
				}
			}
			delete(paths, node.ID)
		}
	}
	fmt.Printf("%d\n", paths["out"].Valids)
}
