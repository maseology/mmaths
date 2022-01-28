package topology

import "log"

// OrderedForest returns a concurrent-safe optimized ordering of a set trees
func OrderedForest(fromto map[int]int, root, ad int) [][]int {
	dg := NewDirectedGraph(fromto, root)
	frst := dg.Forest()

	// Bottom heavy
	ifrst, tx := make([][][]int, len(frst)), 0
	for i, tree := range frst {
		ltree := len(tree)
		ifrst[i] = make([][]int, ltree)
		if ltree > tx {
			tx = ltree
		}
		for j, ns := range tree {
			jj := ltree - j - 1
			ifrst[i][jj] = make([]int, len(ns))
			for k, n := range ns {
				ifrst[i][jj][k] = n.I[0]
			}
		}
	}
	if len(frst) == 1 {
		s := ifrst[0]
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	}

	out := make([][]int, tx)
	for i := 0; i < tx; i++ {
		ii := tx - i - 1
		out[ii] = []int{}
		for j := 0; j < len(ifrst); j++ {
			if len(ifrst[j]) > i {
				out[ii] = append(out[ii], ifrst[j][i]...)
			}
		}
	}

	// // Distribute singles
	// //  worked as planned, but didn't seem to optimize my concurrent rdrr runs, could still have value..?
	// ifrst, nord := make([][][]int, len(frst)), make([]int, len(frst))
	// for i, tree := range frst {
	// 	ifrst[i] = make([][]int, len(tree))
	// 	nord[i] = len(tree)
	// 	for j, ns := range tree {
	// 		ifrst[i][j] = make([]int, len(ns))
	// 		for k, n := range ns {
	// 			ifrst[i][j][k] = n.ID
	// 		}
	// 	}
	// }
	// if len(frst) == 1 {
	// 	return ifrst[0]
	// }

	// //optimize ifrst
	// sngl, multi, ox := []int{}, [][][]int{}, 0
	// for _, v := range ifrst {
	// 	switch len(v) {
	// 	case 1:
	// 		if len(v[0]) != 1 {
	// 			log.Fatalf("topology.go OrderedForest error 1\n")
	// 		}
	// 		sngl = append(sngl, v[0][0])
	// 	default:
	// 		multi = append(multi, v)
	// 		if len(v) > ox {
	// 			ox = len(v)
	// 		}
	// 	}
	// }
	// out := make([][]int, ox)
	// for i := 0; i < ox; i++ {
	// 	out[i] = []int{} // initialize
	// }
	// for _, v := range multi {
	// 	ii := ox - len(v)
	// 	for i := ii; i < ox; i++ {
	// 		out[i] = append(out[i], v[i-ii]...)
	// 	}
	// }

	// getMins := func() []int {
	// 	mn := math.MaxInt32
	// 	for _, v := range out {
	// 		if len(v) < mn {
	// 			mn = len(v)
	// 		}
	// 	}
	// 	ns := []int{}
	// 	for i, v := range out {
	// 		if len(v) == mn {
	// 			ns = append(ns, i)
	// 		}
	// 	}
	// 	return ns
	// }

	// ii := 0
	// for {
	// 	for _, i := range getMins() {
	// 		out[i] = append(out[i], sngl[ii])
	// 		ii++
	// 		if ii == len(sngl) {
	// 			break
	// 		}
	// 	}
	// 	if ii == len(sngl) {
	// 		break
	// 	}
	// }

	// check
	m := make(map[int]bool, len(fromto))
	for _, v := range out {
		// fmt.Println(i, len(v))
		for _, c := range v {
			if _, ok := m[c]; ok {
				log.Fatalf("topology.go OrderedForest error: duplicate node IDs found")
			}
			m[c] = false
		}
	}
	for f := range fromto {
		if _, ok := m[f]; !ok {
			log.Fatalf("topology.go OrderedForest error: missing node ID: %d\n", f)
		}
	}
	for _, v := range out {
		for _, c := range v {
			m[c] = true
			t := fromto[c]
			if t >= 0 && m[t] {
				log.Fatalf("topology.go OrderedForest error: node out of order: from:%d to:%d\n", c, t)
			}
		}
	}
	return out
}
