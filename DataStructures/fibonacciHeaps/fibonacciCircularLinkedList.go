package fibonacciHeaps

import "fmt"

type FibonacciCircularLinkedList struct {
	Sentinel *FibonacciNode
	size     int
}

func NewFibonacciCircularLinkedList() *FibonacciCircularLinkedList {
	sentinel := &FibonacciNode{}

	sentinel.Next = sentinel
	sentinel.Prev = sentinel

	return &FibonacciCircularLinkedList{
		Sentinel: sentinel,
		size:     0,
	}
}

func (cll *FibonacciCircularLinkedList) Head() *FibonacciNode {
	if cll.size == 0 {
		return nil
	}
	return cll.Sentinel.Next
}

func (cll *FibonacciCircularLinkedList) Tail() *FibonacciNode {
	if cll.size == 0 {
		return nil
	}
	return cll.Sentinel.Prev
}

func (cll *FibonacciCircularLinkedList) Search(key int) *FibonacciNode {
	x := cll.Sentinel.Next

	if x == cll.Sentinel {
		return nil
	}

	for x != cll.Sentinel && x.Key != key {
		x = x.Next
	}
	return x
}

func (cll *FibonacciCircularLinkedList) Insert(node *FibonacciNode) {
	node.Next = cll.Sentinel.Next
	cll.Sentinel.Next.Prev = node
	cll.Sentinel.Next = node
	node.Prev = cll.Sentinel

	cll.size++
}

func (cll *FibonacciCircularLinkedList) Delete(node *FibonacciNode) {
	if node == nil || node == cll.Sentinel {
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	cll.size--
}

func (cll *FibonacciCircularLinkedList) Print() {
	if cll.size == 0 {
		fmt.Println("Empty FibonacciHeap.")
		return
	}

	current := cll.Sentinel.Next
	fmt.Print("Root List: ")

	for current != cll.Sentinel {
		fmt.Printf("[Key: %d, Degree: %d] => ", current.Key, current.Degree)
		current = current.Next
	}
	fmt.Println()
}

func FibonacciCircularLinkedListUnion(fcll1, fcll2 *FibonacciCircularLinkedList) *FibonacciCircularLinkedList {
	if fcll1 == nil {
		return fcll2
	}
	if fcll2 == nil {
		return fcll1
	}

	fcll1Head := fcll1.Head()
	fcll2Head := fcll2.Head()

	fcll1Tail := fcll1.Tail()
	fcll2Tail := fcll2.Tail()

	fcll1Tail.Next = fcll2Head
	fcll2Head.Prev = fcll1Tail

	fcll2Tail.Next = fcll1Head
	fcll1Head.Prev = fcll2Tail

	newCLL := &FibonacciCircularLinkedList{
		Sentinel: fcll1.Sentinel,
		size:     fcll1.size + fcll2.size,
	}
	return newCLL
}

func (cll *FibonacciCircularLinkedList) Size() int {
	return cll.size
}

func (cll *FibonacciCircularLinkedList) IsCircular() bool {
	return cll.Sentinel.Next != nil && cll.Sentinel.Prev != nil &&
		cll.Sentinel.Next.Prev == cll.Sentinel.Prev.Next
}
