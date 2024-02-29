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


func (g *Graph) GetVertexBetweenness(v *Vertex) float64 {
	vShortestRoute := 0

	for i, vtx := range g.Vertices {
		// Skip the current vertex
		if v == vtx { continue }

		for j := i; j < len(g.Vertices); j++ {

			// Skip the current vertex
			if v == g.Vertices[j]  { continue }

			_, visitedVertices := g.FindVertexBFS(vtx.Index, g.Vertices[j].Index)

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
