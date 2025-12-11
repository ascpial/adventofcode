package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var example = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
`

type Node struct {
	ID          string
	RawChildren []string
	Children    []*Node
}

type Traveling struct {
	Paths int
	Node  Node
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

	you := nodes["you"]
	seen := map[string]int{}
	seen["you"] = 1
	junk := []*Node{you}
	for len(junk) > 0 {
		node := junk[0]
		junk = junk[1:]
		// fmt.Printf("node to treat: %d; current node: %s\n", len(junk), node.ID)
		if node.ID != "out" {
			for _, child := range node.Children {
				paths, ok := seen[child.ID]
				seen[child.ID] = paths + seen[node.ID]
				// fmt.Printf("node %s has %d incoming paths\n", child.ID, seen[child.ID])
				if !ok {
					junk = append(junk, child)
				}
			}
			delete(seen, node.ID)
		}
	}
	fmt.Printf("%d\n", seen["out"])
}
