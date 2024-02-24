package main

import (
    "github.com/aadv1k/giraffe/giraffe"
    "fmt" 
)

func main() {
    var g giraffe.Graph

    v0 := &giraffe.Vertex{Index: 0, Siblings: []*giraffe.Vertex{}}
    v1 := &giraffe.Vertex{Index: 1, Siblings: []*giraffe.Vertex{}}
    v2 := &giraffe.Vertex{Index: 2, Siblings: []*giraffe.Vertex{}}
    v3 := &giraffe.Vertex{Index: 3, Siblings: []*giraffe.Vertex{}}
    v4 := &giraffe.Vertex{Index: 4, Siblings: []*giraffe.Vertex{}}

    g.AddVertex(v0)
    g.AddVertex(v1)
    g.AddVertex(v2)
    g.AddVertex(v3)
    g.AddVertex(v4)

    g.AddEdge(&giraffe.Edge{Start: v0, End: v1})
    g.AddEdge(&giraffe.Edge{Start: v1, End: v2})
    g.AddEdge(&giraffe.Edge{Start: v2, End: v3})
    g.AddEdge(&giraffe.Edge{Start: v3, End: v4})
    g.AddEdge(&giraffe.Edge{Start: v4, End: v0})
    g.AddEdge(&giraffe.Edge{Start: v0, End: v2})
    g.AddEdge(&giraffe.Edge{Start: v1, End: v3})
    g.AddEdge(&giraffe.Edge{Start: v2, End: v4})
    g.AddEdge(&giraffe.Edge{Start: v3, End: v0})
    g.AddEdge(&giraffe.Edge{Start: v4, End: v1})

		found, visited := g.FindVertex(5)
		fmt.Printf("BFS: [%d] ", found.Index)
		for _, v := range visited {
			fmt.Printf("%d, ", v.Index)
		}
		fmt.Println()
}
