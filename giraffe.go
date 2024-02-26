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

func (g *Graph) AddEdge(e *Edge) error {
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

// Apply K-Means Cluster using degree as the centrality metric for the feature
func (g *Vertex) KMeansCluster(k int, centrality int) {
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
	_, vtx := g.FindVertex(v.Index)

	if vtx != nil {
		return fmt.Errorf("A vertex with that index %d already exists.", v.Index)
	}

	g.Vertices = append(g.Vertices, v)

	for _, sibling := range v.Siblings {
		g.Edges = append(g.Edges, &Edge{v, sibling})
	}

	return nil  
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

func HasVisited(visited []*Vertex, vtx *Vertex) bool {
		for _, visitedVtx := range visited {
			if visitedVtx.Index == vtx.Index {
				return true
			}
		}

		return false
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

// Do a Breadth-First Search to find the vertex
func (g *Graph) FindVertex(index int) (*Vertex, []*Vertex) {
	if len(g.Vertices) == 0 {
		return nil, nil
	}

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
