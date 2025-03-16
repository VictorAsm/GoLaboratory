package circularlinkedlist

import "fmt"

type Node struct {
	key  int
	next *Node
	prev *Node
}

type CircularLinkedList struct {
	size int
	head *Node
	tail *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	head := &Node{key: 0}
	tail := &Node{key: 0}

	head.next = tail
	head.prev = tail

	tail.prev = head
	tail.next = head

	return &CircularLinkedList{
		size: 0,
		head: head,
		tail: tail,
	}
}

func (dbl *CircularLinkedList) Search(key int) *Node {
	x := dbl.head.next

	if x == dbl.tail {
		return nil
	}

	for x != dbl.tail && x.key != key {
		x = x.next
	}
	return x
}

func (dbl *CircularLinkedList) Insert(value int) {
	node := &Node{key: value}
	x := dbl.tail.prev

	node.prev = x
	node.next = dbl.tail
	x.next = node

	dbl.tail.prev = node
	dbl.size++
}

func (dbl *CircularLinkedList) Delete(value int) {
	x := dbl.head.next

	for x != dbl.tail && x.key != value {
		x = x.next
	}

	if x == dbl.tail {
		return
	}

	x.prev.next = x.next
	x.next.prev = x.prev
	dbl.size--
}

func (dbl *CircularLinkedList) Print() {
	if dbl.size == 0 {
		fmt.Println("Empty list")
	}

	x := dbl.head.next
	i := 0

	for x != dbl.tail {
		fmt.Printf("[p: %d | v: %d] => ", i, x.key)
		i++
		x = x.next
	}
}

func (dbl *CircularLinkedList) Size() int {
	return dbl.size
}

func (dbl *CircularLinkedList) IsCircular() bool {
	return dbl.tail.next == dbl.head && dbl.head.prev == dbl.tail
}
