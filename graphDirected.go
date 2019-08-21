package mmaths

// NewDirectedGraph creates a graph from a given "fromto" map
func NewDirectedGraph(downstream map[int]int) map[int]*Node {
	nodes := make(map[int]*Node, len(downstream))
	for k := range downstream {
		nodes[k] = &Node{ID: k, p: make(chan int)}
	}
	for k, v := range downstream {
		nodes[k].DS = append(nodes[k].DS, nodes[v])
	}
	for _, v := range nodes {
		for _, d := range v.DS {
			if d != nil {
				d.US = append(d.US, v)
			}
		}
	}
	return nodes
}
