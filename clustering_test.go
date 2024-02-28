package giraffe

import (
	"testing"
)


func MakeClusterGraph() *Graph {
	var g Graph

	g.AddVertex(&Vertex{Index: 0})
	g.AddVertex(&Vertex{Index: 1})
	g.AddVertex(&Vertex{Index: 2})
	g.AddVertex(&Vertex{Index: 3})
	g.AddVertex(&Vertex{Index: 4})


	// Cluster 1
	g.AddEdge(&Edge{Start: g.GetVertex(0), End: g.GetVertex(1)})
	g.AddEdge(&Edge{Start: g.GetVertex(1), End: g.GetVertex(2)})
	g.AddEdge(&Edge{Start: g.GetVertex(2), End: g.GetVertex(0)})

	// Cluster 2
	g.AddEdge(&Edge{Start: g.GetVertex(0), End: g.GetVertex(4)})
	g.AddEdge(&Edge{Start: g.GetVertex(4), End: g.GetVertex(3)})

	return &g
}

func ShouldBeIn(t *testing.T, v *Vertex, cluster []KClusterChild) {
    found := false

		for _, elem := range cluster {
			if elem.Vertex == v {
				found = true
			}
		}

    if !found {
        t.Errorf("Expected vertex %d to be in the cluster.", v.Index)
    }
}


func TestKMeansClustering(t *testing.T)  {
	g := MakeClusterGraph()

	K := 2
	clusters, _ := g.KMeansClustering(K)

	if clusters == nil {
		t.Errorf("Expected 2D Array of clusters, got nil")
		return
	}

	if len(clusters) != K {
		t.Errorf("Expected %d clusters, got %d", K, len(clusters))
		return
	}

	ShouldBeIn(t, g.GetVertex(0), clusters[0])
	ShouldBeIn(t, g.GetVertex(1), clusters[0])
	ShouldBeIn(t, g.GetVertex(2), clusters[0])


	ShouldBeIn(t, g.GetVertex(3), clusters[1])
	ShouldBeIn(t, g.GetVertex(4), clusters[1])
}
