package topology

// Node is a topological node
type Node struct {
	S      []float64
	I      []int // I[0]: dimension; I[1]: ID
	US, DS []*Node
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

// Junctions returns a slice of nodes having more than one connection either downstream or upstream
func Junctions(g []*Node) []*Node {
	var out []*Node
	for _, v := range g {
		if len(v.US) > 1 || len(v.DS) > 1 {
			out = append(out, v)
		}
	}
	return out
}
