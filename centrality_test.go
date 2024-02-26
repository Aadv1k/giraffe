package giraffe

import (
	"testing"
	"log"
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

func TestDegree(t *testing.T) {
	g := MakePyramidGraph()
	deg := g.GetDegree()

	if deg == nil {
		t.Errorf("Method expected to return array, got nil")
		return
	}

	node := 5
	degOfVertex := deg[node]
	expectedDeg := 2

	if degOfVertex != expectedDeg {
		t.Errorf("%d expected a degree of %d, got %d", node, expectedDeg, degOfVertex)
	}
}

func TestBetweeness(t *testing.T) {
	t.Skip()
}
