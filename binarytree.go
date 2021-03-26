package mmaths

type binaryNode struct {
	left  *binaryNode
	right *binaryNode
	val   float64
	indx  int
}

// IndexOf -1: place at start; n: place at end; i: place after i, before i+1
func (n *binaryNode) IndexOf(val float64) int {
	i := -2
	n.Search(&val, &i)
	return i
}

func (n *binaryNode) Search(val *float64, sid *int) {
	if *sid <= -2 {
		if *val > n.val {
			if n.right == nil {
				*sid = n.indx
			} else {
				n.right.Search(val, sid)
			}
		} else if *val < n.val {
			if n.left == nil {
				*sid = n.indx - 1
			} else {
				n.left.Search(val, sid)
			}
		} else {
			*sid = n.indx
		}
	}
}

// modified from: https://www.golangprograms.com/golang-program-to-implement-binary-tree.html
// type BinaryTree struct { // not yet used
// 	root *binaryNode
// }

// // Insert node: the first entered becomes the root
// func (t *BinaryTree) Insert(indx int, val float64) *BinaryTree {
// 	if t.root == nil {
// 		t.root = &binaryNode{val: val, indx: indx, left: nil, right: nil}
// 	} else {
// 		t.root.insert(val, indx)
// 	}
// 	return t
// }

// func (n *binaryNode) insert(val float64, indx int) {
// 	if n == nil {
// 		return
// 	} else if val <= n.val {
// 		if n.left == nil {
// 			n.left = &binaryNode{val: val, indx: indx, left: nil, right: nil}
// 		} else {
// 			n.left.insert(val, indx)
// 		}
// 	} else {
// 		if n.right == nil {
// 			n.right = &binaryNode{val: val, indx: indx, left: nil, right: nil}
// 		} else {
// 			n.right.insert(val, indx)
// 		}
// 	}
// }

// func print(w io.Writer, node *binaryNode, ns int, ch rune) {
// 	if node == nil {
// 		return
// 	}

// 	for i := 0; i < ns; i++ {
// 		fmt.Fprint(w, " ")
// 	}
// 	fmt.Fprintf(w, "%c:%v\n", ch, node.val)
// 	print(w, node.left, ns+2, 'L')
// 	print(w, node.right, ns+2, 'R')
// }
