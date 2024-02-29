package giraffe

import (
	"testing"
	"math"
)

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
	g := MakePyramidGraph()
	bet := g.GetBetweenness()

	totalVertices := len(g.Vertices)
	sumOfBetweenness := 0.0

	for _, elem := range bet {
		sumOfBetweenness += elem
	}

	tolerance := 1e-6

	if math.Abs(float64(totalVertices)-sumOfBetweenness) > tolerance {
		t.Errorf("Expected the total sum of all centralities to be equal to %d but got %.2f", totalVertices, sumOfBetweenness)
	}
}
