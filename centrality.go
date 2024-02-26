package giraffe

import (
	"assert"
)

func (g *Graph) GetDegree() []int {
	if len(g.Vertices) == 0 {
		return nil
	}
	
	res := make([]int, len(g.Vertices))

	for _, v := range g.Vertices {
		count := 0 
		for _, e := range g.Edges {
				if e.End.Index == v.Index {
						count ++ 
				}
		}
		res[v.Index] = count
	}

	return res
}

func (g *Graph) FindNodeBFS(start, end int) (*Vertex, []*Vertex) {
	var queue []*Vertex
	visited := make(map[int]bool)
	var visitedVertices []*Vertex

	startVertex, endVertex := g.FindVertex(start), g.FindVertex(end)

	assert.NotNil(startVertex)
	assert.NotNil(endVertex)

	queue = append(queue, startVertex)

	for len(queue) != 0 {

		// Dequeue
		cur := queue[0]
		queue = queue[1:]

		hasVisited := visited[cur.Index]
		if hasVisited {
			continue
		}

		if cur.Index == endVertex.Index {
			return cur, visitedVertices
		}

		visited[cur.Index] = true
		visitedVertices = append(visitedVertices, cur)

		// Enqueue
		queue = append(queue, cur.Siblings...)
	}

	return nil, visitedVertices
}

func (g *Graph) GetBetweenness() []float64 {
    betweenness := make([]float64, len(g.Vertices))

    for _, vtx := range g.Vertices {
        for _, vtx2 := range g.Vertices {
            if vtx.Index == vtx2.Index {
                continue
            }
            _, visitedVertices := g.FindNodeBFS(vtx.Index, vtx2.Index)
            found := false
            for _, v := range visitedVertices {
                if v.Index == vtx.Index {
                    found = true
                    break
                }
            }
            if found {
                betweenness[vtx.Index]++
            }
        }
    }

    for i := range betweenness {
        betweenness[i] /= float64(len(g.Vertices) - 1)
    }

    return betweenness
}
