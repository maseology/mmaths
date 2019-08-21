package mmaths

// Node is a
type Node struct {
	ID     int
	US, DS []*Node
	p      chan int
}

// Leaves returns a slice of leaf nodes
func Leaves(nodes map[int]*Node) []*Node {
	var out []*Node
	for _, v := range nodes {
		if len(v.US) == 0 {
			out = append(out, v)
		}
	}
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
	Rev(ord)
	return ord
}
