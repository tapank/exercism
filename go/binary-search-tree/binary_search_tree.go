package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
func (bst *BinarySearchTree) Insert(i int) {
	current := bst
	for {
		var next *BinarySearchTree
		if i <= current.data {
			if next = current.left; next == nil {
				current.left = NewBst(i)
				break
			}
		} else {
			if next = current.right; next == nil {
				current.right = NewBst(i)
				break
			}
		}
		current = next
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	sorted := []int{}
	if bst.left != nil {
		sorted = append(sorted, bst.left.SortedData()...)
	}
	sorted = append(sorted, bst.data)
	if bst.right != nil {
		sorted = append(sorted, bst.right.SortedData()...)
	}
	return sorted
}
