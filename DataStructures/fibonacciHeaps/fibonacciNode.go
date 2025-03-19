package fibonacciHeaps

type FibonacciNode struct {
	Key    int
	Next   *FibonacciNode
	Prev   *FibonacciNode
	Degree int
	Parent *FibonacciNode
	Child  *FibonacciNode
	Mark   bool
}
