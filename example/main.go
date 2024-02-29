package main

import (
	"github.com/aadv1k/giraffe"
	"fmt"
)

func main() {
	var g *giraffe.Graph
	g = giraffe.MakeClusterGraph() // Utility function to generate a sample graph

	clusters, means := g.KMeansClustering(2)

	fmt.Printf("%v\n", means)
	fmt.Printf("There are %d clusters", len(clusters))
}
