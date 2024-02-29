package giraffe

import (
	"testing"
)


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


	for i, cluster := range clusters {
		if len(cluster) == 0 {
				t.Errorf("Zero elements found in cluster %d, likely an error.", i)
		}
	}
}
