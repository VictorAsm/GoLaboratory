package pathFinding

import "GoLaboratory/DataStructures/graphStructure"

type NodeData struct {
	Distance    int
	Predecessor int
	Visited     bool
	Value       int
}

type DFSNodeData struct {
	Distance    int
	Predecessor int
	Value       int
	Time        int
	Visited     bool
}

func (n *DFSNodeData) SetDistance(distance int) {
	n.Distance = distance
}

func (n *DFSNodeData) SetPredecessor(predecessor int) {
	n.Predecessor = predecessor
}

func (n *DFSNodeData) SetValue(value int) {
	n.Value = value
}

func (n *DFSNodeData) SetTime(time int) {
	n.Time = time
}

func (n *DFSNodeData) SetVisited(visited bool) {
	n.Visited = visited
}

func BFS(G *graphStructure.Graph, s int) []NodeData {
	nodeInfo := map[int]NodeData{}
	Q := []int{s}
	inOrderVisited := []int{}

	for v := range G.Vertices {
		nodeInfo[v] = NodeData{
			-1,
			-1,
			false,
			G.Vertices[v].Val,
		}
	}

	nodeInfo[s] = NodeData{
		0,
		-1,
		true,
		G.Vertices[s].Val,
	}

	// WHILE Q NOT EMPTY
	for len(Q) > 0 {
		// u = DEQUEUE(Q)
		current := Q[0]
		Q = Q[1:]
		inOrderVisited = append(inOrderVisited, current)

		// FOR EACH VERTEX v IN G.NEIGHBORS
		for _, v := range G.Neighbors(current) {
			// ACA TENGO LOS VECINOS DE CURRENT: [d1, d2, d3, ..., dn]
			if !nodeInfo[v].Visited {
				Q = append(Q, v)
				nodeInfo[v] = NodeData{
					nodeInfo[current].Distance + 1,
					current,
					true,
					G.Vertices[v].Val,
				}
			}
		}
	}

	var result []NodeData
	for _, node := range inOrderVisited {
		result = append(result, nodeInfo[node])
	}

	return result
}

func DFSVisit(G *graphStructure.Graph, u int, time *int, nodeData map[int]*DFSNodeData, visitedOrder *[]int) {

	*time++
	nodeData[u].Distance = *time
	nodeData[u].Visited = true
	*visitedOrder = append(*visitedOrder, u)

	for _, v := range G.Neighbors(u) {
		if !nodeData[v].Visited {
			nodeData[v].Predecessor = u
			DFSVisit(G, v, time, nodeData, visitedOrder)
		}
	}

	*time++
	nodeData[u].Time = *time
}

func DFS(G *graphStructure.Graph, s int) []DFSNodeData {
	nodeInfo := map[int]*DFSNodeData{}
	inOrderVisited := []int{}

	for u := range G.Vertices {
		nodeInfo[u] = &DFSNodeData{
			Distance:    -1,
			Predecessor: -1,
			Value:       G.Vertices[u].Val,
			Time:        -1,
			Visited:     false,
		}
	}

	time := 0

	if _, exists := G.Vertices[s]; exists {
		DFSVisit(G, s, &time, nodeInfo, &inOrderVisited)
	}

	for u := range G.Vertices {
		if !nodeInfo[u].Visited {
			DFSVisit(G, u, &time, nodeInfo, &inOrderVisited)
		}
	}

	var result []DFSNodeData
	for _, node := range inOrderVisited {
		result = append(result, *nodeInfo[node])
	}

	return result
}
