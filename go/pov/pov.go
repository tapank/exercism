package pov

type Tree struct {
	value    string
	parent   *Tree
	children []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	tr := &Tree{value: value}
	tr.children = children
	for _, ch := range children {
		ch.parent = tr
	}
	return tr
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// Element order does not matter.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if tr.parent != nil {
		result += " parent::" + tr.parent.value + ":: "
	}
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	fromTree := tr.Find(from)
	if fromTree == nil {
		return nil
	}

	if fromTree == tr {
		return tr
	}
	current, next := fromTree, fromTree.parent
	for next != nil {
		// add parent to children
		current.children = append(current.children, next)
		// remove current from next.children
		for i := 0; i < len(next.children); i++ {
			if next.children[i] == current {
				next.children = append(next.children[:i], next.children[i+1:]...)
				break
			}
		}
		// climb one step up the tree
		next.parent, next, current = current, next.parent, next
	}
	// finally, erase parent of new root
	fromTree.parent = nil
	return fromTree
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	fromTree := tr.Find(from)
	toTree := tr.Find(to)
	if fromTree == nil || toTree == nil {
		return nil
	}
	return path(fromTree, toTree)
}

// Find returns the tree node with the given value, nil if not found.
func (tr *Tree) Find(val string) *Tree {
	if tr.value == val {
		return tr
	}
	for _, tree := range tr.children {
		if t := tree.Find(val); t != nil {
			return t
		}
	}
	return nil
}

// path assumes the from and to tree nodes to be present.
func path(f, t *Tree) []string {
	// nodes from the "from node" up to root
	fToRoot := []string{}
	for node := f; node != nil; node = node.parent {
		fToRoot = append(fToRoot, node.value)
	}
	// nodes from the "to node" up to root
	tToRoot := []string{}
	for node := t; node != nil; node = node.parent {
		tToRoot = append(tToRoot, node.value)
	}
	// find commonRoot root and trim the paths
	var commonRoot string
	for lf, lt := len(fToRoot), len(tToRoot); lf > 0 && lt > 0 && (fToRoot[lf-1] == tToRoot[lt-1]); {
		commonRoot = fToRoot[lf-1]
		fToRoot = fToRoot[:lf-1]
		tToRoot = tToRoot[:lt-1]
		lf--
		lt--
	}
	// now stitch the paths together
	path := fToRoot
	path = append(path, commonRoot)
	for i := len(tToRoot) - 1; i >= 0; i-- {
		path = append(path, tToRoot[i])
	}
	return path
}
