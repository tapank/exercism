package linkedlist

import (
	"errors"
)

// Define the List and Element types here.
type List struct {
	head *Node
	len  int
}

type Node struct {
	val  int
	next *Node
}

func New(l []int) *List {
	list := &List{}
	for _, n := range l {
		list.Push(n)
	}
	return list
}

func (l *List) Size() int {
	return l.len
}

func (l *List) Push(element int) {
	n := &Node{val: element}
	if l.len == 0 {
		l.head = n
	} else {
		for node := l.head; ; node = node.next {
			if node.next == nil {
				node.next = n
				break
			}
		}
	}
	l.len++
}

func (l *List) Pop() (int, error) {
	if l.len == 0 {
		return 0, errors.New("list empty")
	}
	var v int
	if l.len == 1 {
		v = l.head.val
		l.head = nil
	} else {
		var prev *Node
		for node := l.head; ; prev, node = node, node.next {
			if node.next == nil {
				prev.next = nil
				v = node.val
				break
			}
		}
	}
	l.len--
	return v, nil
}

func (l *List) Array() []int {
	if l.len == 0 {
		return nil
	}
	arr := make([]int, 0, l.len)
	for node := l.head; node != nil; node = node.next {
		arr = append(arr, node.val)
	}
	return arr
}

func (l *List) Reverse() *List {
	if l.len < 2 {
		return l
	}
	var newHead *Node
	for node := l.head; node != nil; node = l.head {
		l.head = node.next
		newHead, node.next = node, newHead
	}
	l.head = newHead
	return l
}
