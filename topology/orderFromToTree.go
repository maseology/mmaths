package topology

import "github.com/maseology/mmaths/slice"

// OrderFromToTree returns the topological order of a set of from-to connections forming a tree graph
func OrderFromToTree(fromto map[int]int, root int) []int {
	ord := make([]int, 0, len(fromto))
	tofrom := make(map[int][]int, len(fromto))
	for k, v := range fromto {
		tofrom[v] = append(tofrom[v], k)
	}
	queue := make([]int, 0)
	queue = append(queue, tofrom[root]...) // roots
	for {
		if len(queue) == 0 {
			break
		}
		// pop
		x := queue[0]
		ord = append(ord, x)
		queue = queue[1:]
		// push
		if f, ok := tofrom[x]; ok { // otherwise leaves
			queue = append(queue, f...)
		}
	}
	slice.Rev(ord)
	return ord
}
