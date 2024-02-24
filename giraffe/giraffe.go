package giraffe

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

func (g *Graph) AddEdge(e *Edge) {
	g.Edges = append(g.Edges, e)
}

func (g *Graph) AddVertex(v *Vertex) {
	g.Vertices = append(g.Vertices, v)

	for _, sibling := range v.Siblings {
		g.Edges = append(g.Edges, &Edge{v, sibling})
	}
}

// Do a Breadth-First Search to find the vertex
func (g *Graph) FindVertex(index int) *Vertex {
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

		if currentNode.Index == index {
			return currentNode
		}

		for _, sibling := range currentNode.Siblings {
			queue = append(queue, sibling)
		}
	}

	return nil
}
