package topology

func VertexToVertex(nds []*Node) []*Node {
	tnds, nv := make([][]*Node, len(nds)), 0
	for i, n := range nds {
		if n.I[1] != i {
			panic("assumption fail")
		}
		tnds[i] = n.Segmentize()
		nv += len(tnds[i]) - 1
	}
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
	}

	verts := make([]*Node, nv)
	c := 0
	for _, ns := range tnds {
		for j := 1; j < len(ns); j++ {
			if j == 1 && ns[j].DS[0].DS == nil {
				ns[j].DS = nil // roots
			}
			ns[j].I = append(ns[j].I, c)
			if len(ns[j].I) != 2 {
				panic("assumption fail")
			}
			verts[c] = ns[j]
			c++
		}
	}

	return verts
}
