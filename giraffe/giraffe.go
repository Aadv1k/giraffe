package giraffe

import (
	// "log"
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


// TODO: add edges to neighbours, and vice-versa
// TODO: opt for AddSibling, and making Sibling field private
// TODO: check for possible clashing
// TODO: make sure bi-directional edges are not possible
func (g *Graph) AddEdge(e *Edge) {
	g.Edges = append(g.Edges, e)

	for _, vtx := range g.Vertices {

		if e.Start.Index == vtx.Index {
			vtx.Siblings = append(vtx.Siblings, e.End)
		}

		if e.End.Index == vtx.Index {
			vtx.Siblings = append(vtx.Siblings, e.Start)
		}
	}
}

func (g *Graph) AddVertex(v *Vertex) {
	g.Vertices = append(g.Vertices, v)

	for _, sibling := range v.Siblings {
		g.Edges = append(g.Edges, &Edge{v, sibling})
	}
}

func PopLast[K any](arr[]K) K {
	if len(arr) == 0 {
		return *new(K)
	}

	last := len(arr) - 1
	elem := arr[last]
	arr = append(arr[:last], arr[last+1:]...)

	return elem
}

// Do a Depth-First Search to find the vertex
func (g *Graph) FindVertexDFS(index int) (*Vertex, []*Vertex) {
	var stack []*Vertex
	var visited []*Vertex

	stack = append(stack, g.Vertices[0])

	for vtx := PopLast(stack); vtx != nil; {
		nodeIsVisited := false
		for _, visitedVtx := range visited {
			if visitedVtx.Index == vtx.Index {
				nodeIsVisited = true
				break
			}
		}

		if nodeIsVisited { continue }

		stack = append(stack, vtx.Siblings...)
		visited = append(visited, vtx)

		if index == vtx.Index {
			return vtx, visited
		}
	}

	return nil, nil
}

// Do a Breadth-First Search to find the vertex
func (g *Graph) FindVertex(index int) (*Vertex, []*Vertex) {
	var queue []*Vertex
	var visited []*Vertex

	queue = append(queue, g.Vertices[0])

	for len(queue) != 0 {
		currentNode := queue[0]
		queue = queue[1:]

		nodeIsVisited := false

		for _, visitedVtx := range visited {
			if visitedVtx.Index == currentNode.Index {
				nodeIsVisited = true
				break
			}
		}

		if nodeIsVisited {
			break
		}

		visited = append(visited, currentNode)

		if currentNode.Index == index {
			return currentNode, visited
		}

		for _, sibling := range currentNode.Siblings {
			queue = append(queue, sibling)
		}
	}

	return nil, nil
}
