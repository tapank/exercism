package linkedlist

import (
	"errors"
)

type Node struct {
	Value interface{}
	p, n  *Node
}

type List struct {
	head, tail *Node
}

func NewList(args ...interface{}) *List {
	if args == nil {
		return nil
	}
	l := &List{}
	var first, last *Node
	for _, v := range args {
		node := &Node{Value: v}
		if first == nil {
			first = node
		}
		node.p = last
		if last != nil {
			last.n = node
		}
		last = node
	}
	l.head, l.tail = first, last
	return l
}

func (n *Node) Next() *Node {
	return n.n
}

func (n *Node) Prev() *Node {
	return n.p
}

// Unshift adds element at the front
func (l *List) Unshift(v interface{}) {
	node := &Node{Value: v}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.head.p = node
		node.n = l.head
		l.head = node
	}
}

// Push adds element at the back
func (l *List) Push(v interface{}) {
	node := &Node{Value: v}
	if l.head == nil {
		l.head, l.tail = node, node
	} else {
		l.tail.n = node
		node.p = l.tail
		l.tail = node
	}
}

// Shift removes first node and returns its value
func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("shift: list empty")
	}
	node := l.head
	if node.n == nil {
		l.head, l.tail = nil, nil
	} else {
		l.head = node.n
		l.head.p = nil
	}
	return node.Value, nil
}

// Pop removes last node and returns its value
func (l *List) Pop() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("pop: list empty")
	}

	node := l.tail
	if l.head == l.tail {
		l.head, l.tail = nil, nil
	} else {
		l.tail = node.p
		l.tail.n = nil
	}
	return node.Value, nil
}

func (l *List) Reverse() {
	if l.head == l.tail {
		return
	}
	node := l.head
	for node != nil {
		node.n, node.p = node.p, node.n
		node = node.p
	}
	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
