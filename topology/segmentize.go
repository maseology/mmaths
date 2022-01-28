package topology

// Segmentize breaks a polyline "node" into a set of connected vertex nodes
func (n *Node) Segmentize() []*Node {
	d := n.I[0]
	nv := len(n.S) / d
	o := make([]*Node, nv)
	for i := 0; i < nv; i++ {
		a := make([]float64, d)
		for j := 0; j < d; j++ {
			a[j] = n.S[i*d+j]
		}
		o[i] = &Node{
			S: a,
			I: []int{d},
		}
	}
	for i := 0; i < nv; i++ {
		if i > 0 {
			o[i].DS = []*Node{o[i-1]}
		}
		if i < nv-1 {
			o[i].US = []*Node{o[i+1]}
		}
	}
	return o
}
