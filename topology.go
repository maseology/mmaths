package mmaths

import (
	"fmt"

	"github.com/maseology/mmio"
)

// Node is a topological node
type Node struct {
	ID     int
	US, DS []*Node
	p      chan int
}

// Roots return nodes without downslope nodes
func Roots(g []*Node) []*Node {
	r := []*Node{}
	for _, n := range g {
		if len(n.DS) == 0 {
			r = append(r, n)
		}
	}
	return r
}

// Leaves returns a slice of leaf nodes
func Leaves(g []*Node) []*Node {
	var out []*Node
	for _, v := range g {
		if len(v.US) == 0 {
			out = append(out, v)
		}
	}
	return out
}

// Climb returns the upstream nodes
func (n *Node) Climb() []*Node {
	out := []*Node{}
	var recurs func(*Node)
	recurs = func(nn *Node) {
		out = append(out, nn)
		for _, nnn := range nn.US {
			recurs(nnn)
		}
	}
	recurs(n)
	return out
}

// OrderFromToTree returns the topological order of a set of from-to connections forming a tree graph
func OrderFromToTree(fromto map[int]int, root int) []int {
	ord := make([]int, 0, len(fromto))
	tofrom := make(map[int][]int, len(fromto))
	for k, v := range fromto {
		if _, ok := tofrom[v]; ok {
			tofrom[v] = append(tofrom[v], k)
		} else {
			tofrom[v] = []int{k}
		}
	}
	queue := make([]int, 0)
	for _, v := range tofrom[root] {
		queue = append(queue, v) // roots
	}
	for {
		if len(queue) == 0 {
			break
		}
		// pop
		x := queue[0]
		ord = append(ord, x)
		queue = queue[1:]
		// push
		if f, ok := tofrom[x]; ok { // othwise leaves
			for _, v := range f {
				queue = append(queue, v)
			}
		}
	}
	mmio.Rev(ord)
	return ord
}

// OrderedForest returns a concurrent-safe ordering of a set trees
func OrderedForest(fromto map[int]int, root int) [][]int {
	dg := NewDirectedGraph(fromto, root)
	frst := dg.Forest()
	ifrst, nord := make([][][]int, len(frst)), make([]int, len(frst))
	for i, tree := range frst {
		ifrst[i] = make([][]int, len(tree))
		nord[i] = len(tree)
		for j, ns := range tree {
			ifrst[i][j] = make([]int, len(ns))
			for k, n := range ns {
				ifrst[i][j][k] = n.ID
			}
		}
	}
	if len(frst) == 1 {
		return ifrst[0]
	}

	fmt.Println(len(frst), len(nord))
	return nil
}
