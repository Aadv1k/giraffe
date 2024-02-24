package giraffe

import (
    "testing"
)


func TestFindNode(t *testing.T) {
    var g Graph

    v0 := &Vertex{ Index = 0, Siblings = &Vertex{}}
    v1 := &Vertex{ Index = 1, Siblings = &Vertex{}}
    v2 := &Vertex{ Index = 2, Siblings = &Vertex{}}
    v3 := &Vertex{ Index = 3, Siblings = &Vertex{}}

    g.AddVertex(v0)
    g.AddVertex(v1)
    g.AddVertex(v3)
    g.AddVertex(v4)

    g.AddEdge(&Edge{from = v0, to = v1})
    g.AddEdge(&Edge{from = v0, to = v2})
    g.AddEdge(&Edge{from = v0, to = v3})

    // TODO: make sure bi-directional edges are not possible
    // TODO: add edges to neighbours, and vice-versa
    // TODO: opt for AddSibling, and making Sibling field private

    g.AddEdge(&Edge{from = v2, to = v3})
    g.AddEdge(&Edge{from = v2, to = v1})
}

func TestFindNodeDFS(t *testing.T) {
    t.Skip("TestFindNodeDFS goes here")
}

func TestAddNode(t *testing.T) {
    t.Skip("TestAddNode goes here")
}

func TestDeleteNode(t *testing.T) {
    t.Skip("TestDeleteNode goes here")
}
