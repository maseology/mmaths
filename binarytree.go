package mmaths

type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Val   float64
	Indx  int
}

// IndexOf -1: place at start; n: place at end; i: place after i, before i+1
func (n *BinaryNode) IndexOf(val float64) int {
	i := -2
	n.Search(&val, &i)
	return i
}

func (n *BinaryNode) Search(val *float64, sid *int) {
	if *sid <= -2 {
		if *val > n.Val {
			if n.Right == nil {
				*sid = n.Indx
			} else {
				n.Right.Search(val, sid)
			}
		} else if *val < n.Val {
			if n.Left == nil {
				*sid = n.Indx - 1
			} else {
				n.Left.Search(val, sid)
			}
		} else {
			*sid = n.Indx
		}
	}
}

func (node *BinaryNode) AddNode(is *IndexedSlice, first, last int) {
	// from a uniform distribution, picks a "balanced" tree
	if first <= last {
		nid := first + (last-first)/2 // median node. same as (first+last)/2, avoids overflow
		node.Indx = is.Indx[nid]
		node.Val = is.Val[nid]

		if first <= nid-1 {
			node.Left = &BinaryNode{}
			node.Left.AddNode(is, first, nid-1)
		}
		if nid+1 <= last {
			node.Right = &BinaryNode{}
			node.Right.AddNode(is, nid+1, last)
		}
	}
}
