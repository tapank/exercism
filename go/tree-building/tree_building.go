package tree

import (
	"errors"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
	Parent   int
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	var root *Node = nil
	nodes := map[int]*Node{}
	for _, r := range records {
		node := &Node{r.ID, []*Node{}, r.Parent}

		// validate id
		if node.ID < 0 || node.ID >= len(records) {
			return nil, errors.New("non-continuous ids")
		}
		if node.ID != 0 && node.ID <= node.Parent {
			return nil, errors.New("ids not right")
		}
		if _, ok := nodes[node.ID]; ok {
			return nil, errors.New("duplicate node")
		}

		// assign root node
		if node.ID == 0 {
			root = node
			if root.ID != root.Parent {
				return nil, errors.New("root has parent")
			}
		}
		nodes[node.ID] = node
	}

	// flag no root
	if root == nil {
		return nil, errors.New("no root")
	}

	// establish relationships
	for i := 1; i < len(records); i++ {
		if node, ok := nodes[i]; !ok {
			return nil, errors.New("missing node")
		} else if parent, ok := nodes[node.Parent]; !ok {
			return nil, errors.New("missing node")
		} else {
			parent.Children = append(parent.Children, node)
		}
	}
	return root, nil
}
