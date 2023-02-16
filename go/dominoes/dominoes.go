package dominoes

type Domino [2]int

type DominoList struct {
	head, tail *Node
}

func (dl *DominoList) Add(node *Node) {
	if dl.head == nil {
		dl.head = node
		dl.tail = node
	} else {
		dl.tail.next = node
		dl.tail = node
	}
}

func (dl *DominoList) Array() []Domino {
	d := []Domino{}
	for node := dl.head; node != nil; node = node.next {
		d = append(d, node.domino)
	}
	return d
}

type Node struct {
	next   *Node
	v1, v2 int
	domino Domino
}

func New(d Domino) *Node {
	return &Node{
		v1:     d[0],
		v2:     d[1],
		domino: d,
	}
}

func (node *Node) Reverse() {
	node.v1, node.v2 = node.v2, node.v1
	node.domino[0], node.domino[1] = node.domino[1], node.domino[0]
}

func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 {
		return input, true
	}

	// create nodes for adding to list
	nodes := make(map[int]*Node, len(input))
	for i, d := range input {
		nodes[i] = New(d)
	}

	// initialize DominoList
	dl := &DominoList{}
	dl.Add(nodes[0])
	delete(nodes, 0)

	for len(nodes) > 0 {
		n := dl.tail.v2
		l := len(nodes)
		for i, node := range nodes {
			if node.v1 == n {
				dl.Add(node)
				delete(nodes, i)
			} else if node.v2 == n {
				node.Reverse()
				dl.Add(node)
				delete(nodes, i)
			}
		}
		if len(nodes) == l {
			return nil, false
		}
	}
	if dl.head.v1 != dl.tail.v2 {
		return nil, false
	}
	return dl.Array(), true
}
