package pathFinding

import "GoLaboratory/DataStructures/graphStructure"

func BFS(G *graphStructure.Graph, s int) any {
	visited := map[int]bool{}
	Q := []int{s}
	res := []int{}

	// WHILE Q NOT EMPTY
	for len(Q) > 0 {
		// u = DEQUEUE(Q)
		current := Q[0]
		Q = Q[1:]
		res = append(res, current)

		// FOR EACH VERTEX v IN G.NEIGHBORS
		for _, v := range G.Neighbors(current) {
			// ACA TENGO LOS VECINOS DE CURRENT: [d1, d2, d3, ..., dn]
			if !visited[v] {
				Q = append(Q, v)
				visited[v] = true
			}
		}
	}
	return res
}
