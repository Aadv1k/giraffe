package giraffe

import (
	"math/rand"
	"math"
)

type KVec2d struct {
	x int			// degree
	y float64 // betweeness
}

type KClusterChild struct {
	Vertex *Vertex
	Distance float64
}

type KCluster [][]KClusterChild
type KMean []float64

func KVec2d_Euclidean_Distance(v1, v2 KVec2d) float64 {
	return math.Sqrt(math.Pow(float64(v2.x - v1.x), 2) + math.Pow((v2.y - v1.y), 2))
}

func (g *Graph) KMeansClustering(k int) (KCluster, KMean) {
	points := make([]KVec2d, len(g.Vertices))

	lim := KVec2d{}

	for i, vtx := range g.Vertices {
		deg, bet := g.GetVertexDegree(vtx), g.GetVertexBetweenness(vtx)
		points[i] = KVec2d{x: deg, y: bet}

		// Define the upper limit for points 
		if deg > lim.x { lim.x = deg }
		if bet > lim.y { lim.y = bet }
	}

	// Initialize random "K" points within the 2D space
	kPoints := make([]KVec2d, k)
	for _, point := range kPoints {
		point.x = rand.Intn(lim.x)
		point.y = rand.Float64() * lim.y
	}

	clusters := make(KCluster, k)

	for i, point := range points {
		min := 0.0
		minIndex := 0

		for j := 0; j < k; j++ {
			d := KVec2d_Euclidean_Distance(point, kPoints[j])
			if d < min {
				d = min
				minIndex = j
			}
		}
		clusters[minIndex] = append(clusters[minIndex], KClusterChild{Vertex: g.Vertices[i], Distance: min})
	}

	means := make([]float64, k)


	for i, cluster := range clusters {
		for _, child := range cluster {
			means[i] += child.Distance
		}

		means[i] /= float64(len(cluster))
	}

	return clusters, means
}
