package graphStructure

import (
	"fmt"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"math"
)

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

func (g *Graph) Draw(filename string) error {
	p := plot.New()
	p.Title.Text = "Grafo"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := make(plotter.XYs, len(g.Vertices))
	nodePositions := make(map[int]plotter.XY)

	i := 0
	for vertex := range g.Vertices {
		angle := 2 * math.Pi * float64(i) / float64(len(g.Vertices))
		pts[i].X = math.Cos(angle)
		pts[i].Y = math.Sin(angle)
		nodePositions[vertex] = pts[i]
		i++
	}

	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}

	scatter.GlyphStyle.Color = color.RGBA{R: 255, A: 255}
	p.Add(scatter)

	lines := make([]plotter.XYs, 0)
	for u := range g.Vertices {
		for _, v := range g.Neighbors(u) {
			lines = append(lines, plotter.XYs{
				nodePositions[g.Vertices[u].Val],
				nodePositions[g.Vertices[v].Val],
			})
		}
	}

	for _, line := range lines {
		
	}

}
