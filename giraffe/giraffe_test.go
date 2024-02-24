package giraffe

import (
    "testing"
)

func TestFindVertex(t *testing.T) {
    var g Graph

    v0 := &Vertex{ Index: 0}
    v1 := &Vertex{ Index: 1}
    v2 := &Vertex{ Index: 2}
    v3 := &Vertex{ Index: 3}

    g.AddVertex(v0)
    g.AddVertex(v1)
    g.AddVertex(v2)
    g.AddVertex(v3)

    g.AddEdge(&Edge{Start: v0, End: v1})
    g.AddEdge(&Edge{Start: v0, End: v2})
    g.AddEdge(&Edge{Start: v0, End: v3})
    g.AddEdge(&Edge{Start: v2, End: v3})
    g.AddEdge(&Edge{Start: v2, End: v1})

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
    t.Skip("TestFindVertexDFS goes here")
}

func TestAddVertex(t *testing.T) {
    t.Skip("TestAddVertex goes here")
}

func TestDeleteVertex(t *testing.T) {
    t.Skip("TestDeleteVertex goes here")
}
