package giraffe

import (
    "testing"
)

func TestAddVertex(t *testing.T) {
    var g Graph

		v0 := &Vertex{ Index: 0 }
		g.AddVertex(v0)
		err := g.AddVertex(v0)

		if err == nil {
			t.Errorf("Expected adding of duplicate vertex to throw an error")
		}
}

func TestAddEdge(t *testing.T) {
    var g Graph

		v0 := &Vertex{ Index: 0 }
		v1 := &Vertex{ Index: 1 }
		v2 := &Vertex{ Index: 2 }

		g.AddVertex(v0)
		g.AddVertex(v1)
		g.AddVertex(v2)

		err := g.AddEdge(&Edge{Start: v0, End: v0})
		g.AddEdge(&Edge{Start: v0, End: v1})

		if err == nil {
			t.Errorf("Expected cyclic edge creation to throw an error")
		}

		err = g.AddEdge(&Edge{Start: v1, End: v0})

		if err == nil {
			t.Errorf("Expected duplicate edge creation to throw an error")
		}

}



func MakeTestGraph() *Graph {
    var g Graph

    v0 := &Vertex{ Index: 0 }
    v1 := &Vertex{ Index: 1 }
    v2 := &Vertex{ Index: 2 }
    v3 := &Vertex{ Index: 3 }

    g.AddVertex(v0)
    g.AddVertex(v1)
    g.AddVertex(v2)
    g.AddVertex(v3)


    g.AddEdge(&Edge{Start: v0, End: v3})
    g.AddEdge(&Edge{Start: v0, End: v2})
    g.AddEdge(&Edge{Start: v0, End: v1})

		v0.SortSiblings()

    g.AddEdge(&Edge{Start: v2, End: v1})
    g.AddEdge(&Edge{Start: v2, End: v3})

		v2.SortSiblings()
		return &g
}

func TestFindVertex(t *testing.T) {
		g := MakeTestGraph()

    found, visited := g.FindVertex(3)

    if found == nil {
        t.Errorf("Expected to find vertex with index 3, but found nil")
    }

		const nodes_to_visit int = 4

    if len(visited) != nodes_to_visit {
        t.Errorf("Expected %d visited vertices, but got %d", nodes_to_visit, len(visited))
    }

    expectedIndices := map[int]bool{0: true, 1: true, 2: true, 3: true}
    for _, vtx := range visited {
        if !expectedIndices[vtx.Index] {
            t.Errorf("Unexpected visited vertex index: %d", vtx.Index)
        }
    }
}

func TestFindVertexDFS(t *testing.T) {
		g := MakeTestGraph()

    found, visited := g.FindVertexDFS(2)

    if found == nil {
        t.Errorf("Expected to find vertex with index 2, but found nil")
    }
		
		const nodes_to_visit int = 3

    if len(visited) != nodes_to_visit {
        t.Errorf("Expected %d visited vertices, but got %d", nodes_to_visit, len(visited))
    }

    expectedIndices := map[int]bool{0: true, 3: true, 2: true}
    for _, vtx := range visited {
        if !expectedIndices[vtx.Index] {
            t.Errorf("Unexpected visited vertex index: %d", vtx.Index)
        }
    }
}
