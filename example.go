package main

import (
    "github.com/aadv1k/giraffe/giraffe"
    "fmt" 
)

func main() {
    var g giraffe.Graph

    v0 := &giraffe.Vertex{Index: 0, Siblings: []*giraffe.Vertex{}}
    v1 := &giraffe.Vertex{Index: 1, Siblings: []*giraffe.Vertex{v0}}

    g.AddVertex(v0)
    g.AddVertex(v1)

    fmt.Println("Vertices:")
    for _, v := range g.Vertices {
        fmt.Println("\tVertex:", v.Index)
    }

    fmt.Println("Edges:")
    for _, e := range g.Edges {
        fmt.Printf("\tEdge: %d -> %d\n", e.Start.Index, e.End.Index)
    }
}
