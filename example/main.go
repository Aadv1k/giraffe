package main

import (
	"github.com/aadv1k/giraffe"
	"fmt"
)

func main() {
    var g giraffe.Graph

    v0 := &giraffe.Vertex{Index: 0}
    v1 := &giraffe.Vertex{Index: 1}

    g.AddVertex(v0)
    g.AddVertex(v1)

    g.AddEdge(&giraffe.Edge{Start: v0, End: v1})
    
		// Find the provided index using Breadth-First Search 
    found, visited := g.FindVertex(1)

    fmt.Printf("Found: %d\n", found.Index)

		// Returns an array of visited vertices
    fmt.Print("Visited: { ")
    for i, vtx := range visited {
			if i == len(visited) - 1 {
				fmt.Printf("%d }\n", vtx.Index)
				continue
			}
      fmt.Printf("%d, ", vtx.Index)
    }
}
