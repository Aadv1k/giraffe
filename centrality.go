package giraffe

import (
	"math"
)

func (g *Graph) GetDegree() []int {
	if len(g.Vertices) == 0 {
		return nil
	}
	
	res := make([]int, len(g.Vertices))

	for _, v := range g.Vertices {
		res[v.Index] = g.GetVertexDegree(v)
	}

	return res
}

func (g *Graph) GetVertexDegree(v *Vertex) int {
	count := 0 
	for _, e := range g.Edges {
		if e.End.Index == v.Index {
			count++ 
		}
	}
	return count
}

func (g *Graph) FindNodeBFS(start, end int) (*Vertex, []*Vertex) {
	var queue []*Vertex
	visited := make(map[int]bool)
	var visitedVertices []*Vertex

	startVertex, endVertex := g.GetVertex(start), g.GetVertex(end)

	if startVertex == nil || endVertex == nil {
		return nil, nil
	}

	queue = append(queue, startVertex)

	for len(queue) != 0 {

		// Dequeue
		cur := queue[0]
		queue = queue[1:]

		hasVisited := visited[cur.Index]
		if hasVisited {
			continue
		}


		visited[cur.Index] = true
		visitedVertices = append(visitedVertices, cur)

		if cur.Index == endVertex.Index {
			return cur, visitedVertices
		}

		// Enqueue
		queue = append(queue, cur.Siblings...)
	}

	return nil, visitedVertices
}

func (g *Graph) GetVertexBetweenness(v *Vertex) float64 {
	vShortestRoute := 0

	for i, vtx := range g.Vertices {
		// Skip the current vertex
		if v == vtx { continue }

		for j := i; j < len(g.Vertices); j++ {

			// Skip the current vertex
			if v == g.Vertices[j]  { continue }

			_, visitedVertices := g.FindNodeBFS(vtx.Index, g.Vertices[j].Index)

			for _, visitedVtx := range visitedVertices {

				// Check if it was involved in the shortest path
				if v == visitedVtx {
					vShortestRoute++
					break
				}
			}
		}
	}

	betweeness := math.Round(float64(vShortestRoute) / float64(len(g.Vertices)))
	return betweeness
}

func (g *Graph) GetBetweenness() []float64 {
	betweenness := make([]float64, len(g.Vertices))

	for i, vtx := range g.Vertices {
		betweenness[i] = g.GetVertexBetweenness(vtx)
	}

	return betweenness
}
