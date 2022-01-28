package topology

import "github.com/maseology/mmaths/slice"

// DirectedGraph is a set of nodes
type DirectedGraph struct {
	r, n []*Node
}

// NewDirectedGraph builds a (directed) topological graph
func NewDirectedGraph(fromto map[int]int, root int) DirectedGraph {
	m := make(map[int]*Node, len(fromto))
	for k := range fromto {
		m[k] = &Node{I: []int{k}}
	}
	r := []*Node{}
	for k, v := range fromto {
		if v != root {
			mk, mv := m[k], m[v]
			mk.DS = append(mk.DS, mv)
			mv.US = append(mv.US, mk)
			m[k] = mk
			m[v] = mv
		} else {
			r = append(r, m[k])
		}
	}
	n, i := make([]*Node, len(fromto)), 0
	for _, v := range m {
		n[i] = v
		i++
	}
	return DirectedGraph{r: r, n: n}
}

// Forest returns a collection of concurrent-safe trees [tree][order][nodes]
func (dg DirectedGraph) Forest() [][][]*Node {
	type col struct {
		r int
		n [][]*Node
	}
	ch := make(chan col, len(dg.r))
	for _, r := range dg.r {
		go func(r *Node) {
			us := r.Climb()
			for i, j := 0, len(us)-1; i < j; i, j = i+1, j-1 {
				us[i], us[j] = us[j], us[i] // reverse slice
			}
			xr := make(map[int]int, len(us))
			for i := 0; i < len(us); i++ {
				xr[us[i].I[0]] = i
			}
			cnt := make(map[int]int, len(us))
			incr := func(i, v int) {
				if _, ok := cnt[i]; !ok {
					cnt[i] = v + 1
				} else {
					if v+1 > cnt[i] {
						cnt[i] = v + 1
					}
				}
			}
			for _, u := range us {
				incr(u.I[0], 0)
				for _, d := range u.DS {
					incr(d.I[0], cnt[u.I[0]])
				}
			}

			mord, lord := slice.InvertMap(cnt)
			ord := make([][]*Node, len(lord)) // concurrent-safe ordering of nodes
			for i, k := range lord {
				cpy := make([]*Node, len(mord[k]))
				for ii, kk := range mord[k] {
					cpy[ii] = us[xr[kk]]
				}
				ord[i] = cpy
			}
			// ord := make([][]int, len(lord)) // concurrent-safe ordering of nodes
			// for i, k := range lord {
			// 	cpy := make([]int, len(mord[k]))
			// 	copy(cpy, mord[k])
			// 	ord[i] = cpy
			// }
			ch <- col{r: r.I[0], n: ord}
		}(r)
	}

	out := make([][][]*Node, len(dg.r))
	for i := 0; i < len(dg.r); i++ {
		c := <-ch
		out[i] = c.n
	}
	close(ch)

	// for i := 0; i < nx; i++ {
	// 	outt, ii := make([]*Node, len(roots)), 0
	// 	for _, c := range cols {
	// 		if i < len(c.ord) {
	// 			outt[ii] = c.ord[i]
	// 			ii++
	// 		}
	// 	}
	// 	out[i] = make([]*Node, len(outt))
	// 	copy(out[i], outt)
	// }

	return out
}

// // NewDirectedGraph creates a graph from a given "fromto" map
// func NewDirectedGraph(downstream map[int]int) map[int]*Node {
// 	nodes := make(map[int]*Node, len(downstream))
// 	for k := range downstream {
// 		nodes[k] = &Node{ID: k, p: make(chan int)}
// 	}
// 	for k, v := range downstream {
// 		nodes[k].DS = append(nodes[k].DS, nodes[v])
// 	}
// 	for _, v := range nodes {
// 		for _, d := range v.DS {
// 			if d != nil {
// 				d.US = append(d.US, v)
// 			}
// 		}
// 	}
// 	return nodes
// }
