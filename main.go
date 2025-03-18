package main

import (
	"GoLaboratory/Algorithms/pathFinding"
	"GoLaboratory/DataStructures/graphStructure"
	"fmt"
)

func main() {
	adjacencyList := map[int][]int{
		0: {4, 2, 1},
		1: {},
		2: {3, 4},
		3: {1},
		4: {5, 2},
		5: {1, 2},
	}

	g := graphStructure.NewGraph(graphStructure.WithAdjacencyList(adjacencyList))

	//fmt.Println("Grafo creado con los siguientes vértices y conexiones:")
	//for Key, _ := range g.Vertices {
	//	fmt.Printf("Vértice %d: Vecinos -> %v\n", Key, g.Neighbors(Key))
	//}

	fmt.Println("BFS Result")
	pathBFS := pathFinding.BFS(g, 0)

	fmt.Printf("Key %v\n", pathBFS)
}
