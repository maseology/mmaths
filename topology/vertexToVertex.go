package topology

func VertexToVertex(nds []*Node) ([]*Node, map[int][]*Node) {
	tnds, nv := make([][]*Node, len(nds)), 0
	for i, n := range nds {
		if n.I[1] != i {
			panic("assumption fail")
		}
		tnds[i] = n.Segmentize()
		nv += len(tnds[i]) - 1
	}
	segUpDwnNode := make(map[int][]*Node, len(nds))
	for i, n := range nds {
		usv := tnds[i][len(tnds[i])-1] // upstream-most vertex
		if usv.US != nil {
			panic("assumption fail")
		}
		for _, u := range n.US {
			unds := tnds[u.I[1]][1] // downstream-most vertex (minus 1) of upstream polyline/segment
			unds.DS = []*Node{usv}  // overwrite
			usv.US = append(usv.US, unds)
		}
		segUpDwnNode[i] = []*Node{usv, tnds[i][0]}
	}

	verts := make([]*Node, nv)
	c := 0
	for i, ns := range tnds {
		for j := 1; j < len(ns); j++ {
			if j == 1 && ns[j].DS[0].DS == nil {
				ns[j].DS = nil // roots
			}
			ns[j].I = append(ns[j].I, c) // vertex ID
			ns[j].I = append(ns[j].I, nds[i].I[1:]...)
			if len(ns[j].I) != len(nds[i].I)+1 {
				panic("assumption fail")
			}
			verts[c] = ns[j]
			c++
		}
	}

	// final fix
	for _, nd := range segUpDwnNode {
		nd[1].I = nd[0].I
	}

	return verts, segUpDwnNode
}
