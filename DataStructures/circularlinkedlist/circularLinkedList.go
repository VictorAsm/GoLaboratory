package circularlinkedlist

import (
	"GoLaboratory/DataStructures/node"
	"fmt"
)

type CircularLinkedList struct {
	sentinel *node.Node
	size     int
}

func NewCircularLinkedList() *CircularLinkedList {
	sentinel := &node.Node{}

	sentinel.Next = sentinel
	sentinel.Prev = sentinel

	return &CircularLinkedList{
		sentinel: sentinel,
		size:     0,
	}
}

func (cll *CircularLinkedList) Search(key int) *node.Node {
	x := cll.sentinel.Next

	if x == cll.sentinel {
		return nil
	}

	for x != cll.sentinel && x.Key != key {
		x = x.Next
	}
	return x
}

func (cll *CircularLinkedList) Insert(value int) {
	node := &node.Node{Key: value}
	x := cll.sentinel.Prev

	node.Prev = x
	node.Next = cll.sentinel
	x.Next = node

	cll.sentinel.Prev = node
	cll.size++
}

func (cll *CircularLinkedList) Delete(node *node.Node) {
	if node == nil || node == cll.sentinel {
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	cll.size--
}

func (cll *CircularLinkedList) Print() {
	if cll.size == 0 {
		fmt.Println("Empty list")
		return
	}

	x := cll.sentinel.Next
	for x != cll.sentinel {
		fmt.Printf("[ %d ] -> ", x.Key)
		x = x.Next
	}
	fmt.Println("[ SENTINEL ]")
}

func (cll *CircularLinkedList) Size() int {
	return cll.size
}

func (cll *CircularLinkedList) IsCircular() bool {
	return cll.sentinel.Next != nil && cll.sentinel.Prev != nil &&
		cll.sentinel.Next.Prev == cll.sentinel.Prev.Next
}
