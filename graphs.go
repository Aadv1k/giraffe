package giraffe

import (
	"fmt"
) 

/*
      0
     / \
    1   2
   /     \
  3-->4-->5
*/

func MakePyramidGraph() *Graph {
	var g Graph

	//  Yes. I should use a loop.
	g.AddVertex(&Vertex{Index: 0})
	g.AddVertex(&Vertex{Index: 1})
	g.AddVertex(&Vertex{Index: 2})
	g.AddVertex(&Vertex{Index: 3})
	g.AddVertex(&Vertex{Index: 4})
	g.AddVertex(&Vertex{Index: 5})

	g.AddEdge(&Edge{Start: g.GetVertex(0), End: g.GetVertex(1)})
	g.AddEdge(&Edge{Start: g.GetVertex(0), End: g.GetVertex(2)})
	
	g.AddEdge(&Edge{Start: g.GetVertex(1), End: g.GetVertex(3)})
	g.AddEdge(&Edge{Start: g.GetVertex(3), End: g.GetVertex(4)})
	g.AddEdge(&Edge{Start: g.GetVertex(4), End: g.GetVertex(5)})
	
	g.AddEdge(&Edge{Start: g.GetVertex(2), End: g.GetVertex(5)})

	return &g
}

func MakeClusterGraph() *Graph {
	var g Graph

	g.AddVertex(&Vertex{Index: 0})
	g.AddVertex(&Vertex{Index: 1})
	g.AddVertex(&Vertex{Index: 2})

	g.AddVertex(&Vertex{Index: 3})
	g.AddVertex(&Vertex{Index: 4})
	g.AddVertex(&Vertex{Index: 5})


	// Cluster 1
	g.AddEdge(&Edge{Start: g.GetVertex(0), End: g.GetVertex(1)})
	g.AddEdge(&Edge{Start: g.GetVertex(1), End: g.GetVertex(2)})
	g.AddEdge(&Edge{Start: g.GetVertex(2), End: g.GetVertex(0)})

	// Cluster 2
	g.AddEdge(&Edge{Start: g.GetVertex(3), End: g.GetVertex(0)})

	g.AddEdge(&Edge{Start: g.GetVertex(3), End: g.GetVertex(4)})
	g.AddEdge(&Edge{Start: g.GetVertex(3), End: g.GetVertex(5)})

	g.AddEdge(&Edge{Start: g.GetVertex(5), End: g.GetVertex(3)})
	g.AddEdge(&Edge{Start: g.GetVertex(5), End: g.GetVertex(4)})

	g.AddEdge(&Edge{Start: g.GetVertex(4), End: g.GetVertex(3)})
	g.AddEdge(&Edge{Start: g.GetVertex(4), End: g.GetVertex(5)})

	return &g
}

func PrintVertex(vtx *Vertex) {
	fmt.Printf("Vertex: %d\n", vtx.Index)
}

func PrintVertices(vertices []*Vertex) {
	fmt.Print("{ ")
	for i, vtx := range vertices {
		if i == len(vertices) - 1 {
			fmt.Printf("%d }\n", vtx.Index)
			continue
		}
		fmt.Printf("%d, ", vtx.Index)
	}

}
