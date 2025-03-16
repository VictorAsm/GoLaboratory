package fibonacciHeaps

type FibonacciHeap struct {
}

// MakeHeap Creates and return a new heap containing no elements. O(1)
func (fh *FibonacciHeap) MakeHeap() {}

// Insert Inserts element x, whose key has already been filled in, into heap. O(1)
func (fh *FibonacciHeap) Insert(x any) {}

// Minimum returns a pointer to the element in heap whose key is minimum. O(1)
func (fh *FibonacciHeap) Minimum() {}

// ExtractMin deletes the element from heap whose key is minimum, returning a pointer to the element. O(log n)
func (fh *FibonacciHeap) ExtractMin() {}

// DecreaseKey assigns to element x within heap the new key value k, which we assume to be no greater than its current key value. O(1)
func (fh *FibonacciHeap) DecreaseKey(x, k any) {}

// Delete deletes element x from heap. O(log n)
func (fh *FibonacciHeap) Delete(x, k any) {}

// Print Prints the heaps in current format. O(n)
func (fh *FibonacciHeap) Print() {}

// FibonacciHeapUnion Union creates and returns a new heap that contains all the elements of first heap and. O(1)
func FibonacciHeapUnion(h1, h2 *FibonacciHeap) {}
