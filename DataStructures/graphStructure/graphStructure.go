package graphStructure

type Graph struct {
	Vertices map[int]*Vertex
}

type Vertex struct {
	Val   int
	Edges map[int]*Edge
}

type Edge struct {
	Weight int
	Vertex *Vertex
}

type GraphOption func(this *Graph)

func (g *Graph) AddVertex(key, val int) {
	g.Vertices[key] = &Vertex{Val: val, Edges: map[int]*Edge{}}
}

func (g *Graph) AddEdge(srcKey, destKey int, weight int) {
	if _, ok := g.Vertices[srcKey]; !ok {
		return
	}
	if _, ok := g.Vertices[destKey]; !ok {
		return
	}
	g.Vertices[srcKey].Edges[destKey] = &Edge{Weight: weight, Vertex: g.Vertices[destKey]}
}

func (g *Graph) Neighbors(key int) []int {
	result := []int{}

	for _, edge := range g.Vertices[key].Edges {
		result = append(result, edge.Vertex.Val)
	}

	return result
}

func NewGraph(opts ...GraphOption) *Graph {
	g := &Graph{Vertices: map[int]*Vertex{}}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

func WithAdjacencyList(list map[int][]int) GraphOption {
	return func(this *Graph) {
		for vertex, edges := range list {
			if _, ok := this.Vertices[vertex]; !ok {
				this.AddVertex(vertex, vertex)
			}

			for _, edge := range edges {
				if _, ok := this.Vertices[edge]; !ok {
					this.AddVertex(edge, edge)
				}

				this.AddEdge(vertex, edge, 0)
			}
		}
	}
}
