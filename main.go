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

	fmt.Println("Grafo creado con los siguientes vértices y conexiones:")
	for Key, _ := range g.Vertices {
		fmt.Printf("Vértice %d: Vecinos -> %v\n", Key, g.Neighbors(Key))
	}

	fmt.Println("\n")

	resultBFS := pathFinding.BFS(g, 0)

	fmt.Println("Información de los nodos:")
	for _, nodeData := range resultBFS {
		fmt.Printf("Nodo %d:\n", nodeData.Value)
		fmt.Printf("\tPredecesor: %d\n", nodeData.Predecessor)
		fmt.Printf("\tDistancia: %d\n", nodeData.Distance)
		fmt.Printf("\tVisitado: %v\n", nodeData.Visited)
		fmt.Printf("\tValor: %d\n", nodeData.Value)
		fmt.Println()
	}

	resultDFS := pathFinding.DFS(g, 0)

	fmt.Println("Información de los nodos:")
	for _, nodeData := range resultDFS {
		fmt.Printf("Nodo %d:\n", nodeData.Value)
		fmt.Printf("\tPredecesor: %d\n", nodeData.Predecessor)
		fmt.Printf("\tDistancia: %d\n", nodeData.Distance)
		fmt.Printf("\tVisitado: %v\n", nodeData.Visited)
		fmt.Printf("\tValor: %d\n", nodeData.Value)
		fmt.Printf("\tTime: %d\n", nodeData.Time)
		fmt.Println()
	}

}
