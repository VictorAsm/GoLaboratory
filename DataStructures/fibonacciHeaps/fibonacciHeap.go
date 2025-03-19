package fibonacciHeaps

import "fmt"

type FibonacciHeap struct {
	n        int
	min      *FibonacciNode
	rootList *FibonacciCircularLinkedList
}

// MakeFibonacciHeap Creates and return a new heap containing no elements. O(1)
func MakeFibonacciHeap() *FibonacciHeap {
	return &FibonacciHeap{
		n:        0,
		min:      nil,
		rootList: NewFibonacciCircularLinkedList(),
	}
}

// Insert Inserts element x, whose key has already been filled in, into heap. O(1)
func (fh *FibonacciHeap) Insert(x *FibonacciNode) {
	x.Degree = 0
	x.Parent = nil
	x.Child = nil
	x.Mark = false

	if fh.min == nil {
		fh.rootList.Insert(x)
		fh.min = x
	} else {
		fh.rootList.Insert(x)
		if x.Key < fh.min.Key {
			fh.min = x
		}
	}
	fh.n = fh.n + 1
}

// Minimum returns a pointer to the element in heap whose key is minimum. O(1)
func (fh *FibonacciHeap) Minimum() *FibonacciNode {
	return fh.min
}

// ExtractMin deletes the element from heap whose key is minimum, returning a pointer to the element. O(log n)
func (fh *FibonacciHeap) ExtractMin() {
}

// DecreaseKey assigns to element x within heap the new key value k, which we assume to be no greater than its current key value. O(1)
func (fh *FibonacciHeap) DecreaseKey(x, k any) {}

// Delete deletes element x from heap. O(log n)
func (fh *FibonacciHeap) Delete(x, k any) {}

// Print Prints the heaps in current format. O(n)
func (fh *FibonacciHeap) Print() {
	fmt.Println("Fibonacci Heap:")
	fh.rootList.Print()
	if fh.min != nil {
		fmt.Printf("Min Node: Key: %d\n", fh.min.Key)
	}
}

// FibonacciHeapUnion Union creates and returns a new heap that contains all the elements of first heap and. O(1)
func FibonacciHeapUnion(h1, h2 *FibonacciHeap) *FibonacciHeap {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	h := MakeFibonacciHeap()
	h.min = h1.min
	h.rootList = FibonacciCircularLinkedListUnion(h2.rootList, h.rootList)
	if (h1.min == nil) || (h2.min != nil && h2.min.Key < h1.min.Key) {
		h.min = h2.min
	}
	h.n = h1.n + h2.n
	return h
}

func (fh *FibonacciHeap) N() int {
	return fh.n
}
