package topology

// Climb returns the upstream nodes
func (n *Node) Climb() []*Node {
	out := []*Node{}
	eval := map[*Node]bool{}
	var recurs func(*Node)
	recurs = func(nn *Node) {
		out = append(out, nn)
		eval[nn] = true
		for _, nnn := range nn.US {
			if _, ok := eval[nnn]; !ok {
				recurs(nnn)
			}
		}
	}
	recurs(n)
	return out
}
