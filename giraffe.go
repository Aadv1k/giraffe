package giraffe

import (
	// "log"
	"fmt"
)

type Vertex struct {
	Index    int
	Siblings []*Vertex
}

type Edge struct {
	Start *Vertex
	End   *Vertex
}

type Graph struct {
	Vertices []*Vertex
	Edges    []*Edge
}


// Iteratively return first vertex with matched index
func (g *Graph) GetVertex(index int) *Vertex {

	if len(g.Vertices) == 0 {
		return nil
	}

	for _, vtx := range g.Vertices {
		if vtx.Index == index {
			return vtx
		}
	}

	return nil
}

func (g *Graph) AddEdge(e *Edge) error {
	if e.Start == nil || e.End == nil {
		panic("AddEdge: Unexpected edge start or end node was nil")
	}

	if e.Start == e.End {
		return fmt.Errorf("Cyclic nodes not allowed.")
	}

	for _, edg := range g.Edges {
		if (edg.Start == e.Start && edg.End == e.End) || (edg.End == e.Start && edg.Start == e.End) {
			return fmt.Errorf("An edge with the exact relation {%d} -> {%d} already exists.", e.Start.Index, e.End.Index)
		}
	}

	g.Edges = append(g.Edges, e)

	for _, vtx := range g.Vertices {
		if e.Start.Index == vtx.Index {
			vtx.Siblings = append(vtx.Siblings, e.End)
		}

		if e.End.Index == vtx.Index {
			vtx.Siblings = append(vtx.Siblings, e.Start)
		}
	}

	return nil
}

// Uses Selection sort to order the siblings
func (v *Vertex) SortSiblings() {
	for i := 0; i < len(v.Siblings) - 1; i += 1 {
		minIndex := i

		for j := i + 1; j < len(v.Siblings); j+=1 {
			if v.Siblings[j].Index < v.Siblings[minIndex].Index {
				minIndex = j
			}
		}

		if minIndex != i {
				v.Siblings[minIndex], v.Siblings[i] = v.Siblings[i], v.Siblings[minIndex]
		}
	}
}

func (g *Graph) AddVertex(v *Vertex) error {
	vtx := g.GetVertex(v.Index)

	if vtx != nil {
		return fmt.Errorf("A vertex with that index %d already exists.", v.Index)
	}

	g.Vertices = append(g.Vertices, v)

	for _, sibling := range v.Siblings {
		g.Edges = append(g.Edges, &Edge{v, sibling})
	}

	return nil  
}

// Do a Depth-First Search to find the vertex
func (g *Graph) FindVertexDFS(index int) (*Vertex, []*Vertex) {
	if len(g.Vertices) <= 1 {
		return nil, nil
	}

	var stack []*Vertex
	var visited []*Vertex

	stack = append(stack, g.Vertices[0])

	for len(stack) != 0 {
		vtx := PopLast(stack)

		if HasVisited(visited, vtx) {
			continue;
		}

		for _, sib := range vtx.Siblings {
				if !HasVisited(visited, sib) {
						stack = append(stack, sib)
				}
		}

		visited = append(visited, vtx)

		if index == vtx.Index {
			return vtx, visited
		}
	}

	return nil, nil
}

func (g *Graph) FindVertexBFS(start, end int) (*Vertex, []*Vertex) {
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
