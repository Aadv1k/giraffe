package main

import (
    "github.com/aadv1k/giraffe/giraffe"
)

func main() {
    var g Graph

    CreateNode(&g, 19, []int{})                   // Node 19
    CreateNode(&g, 18, []int{19})                 // Node 18 connected to Node 19
    CreateNode(&g, 17, []int{18, 19})             // Node 17 connected to Node 18 and Node 19
    CreateNode(&g, 16, []int{17, 18})             // Node 16 connected to Node 17 and Node 18
    CreateNode(&g, 15, []int{16, 17})             // Node 15 connected to Node 16 and Node 17
    CreateNode(&g, 14, []int{15, 16})             // Node 14 connected to Node 15 and Node 16
    CreateNode(&g, 13, []int{14, 15})             // Node 13 connected to Node 14 and Node 15
    CreateNode(&g, 12, []int{13, 14})             // Node 12 connected to Node 13 and Node 14
    CreateNode(&g, 11, []int{12, 13})             // Node 11 connected to Node 12 and Node 13
    CreateNode(&g, 10, []int{11, 12})             // Node 10 connected to Node 11 and Node 12
    CreateNode(&g, 9, []int{10, 11})              // Node 9 connected to Node 10 and Node 11
    CreateNode(&g, 8, []int{9, 10})               // Node 8 connected to Node 9 and Node 10
    CreateNode(&g, 7, []int{8, 9})                // Node 7 connected to Node 8 and Node 9
    CreateNode(&g, 6, []int{7, 8})                // Node 6 connected to Node 7 and Node 8
    CreateNode(&g, 5, []int{6, 7})                // Node 5 connected to Node 6 and Node 7
    CreateNode(&g, 4, []int{5, 6})                // Node 4 connected to Node 5 and Node 6
    CreateNode(&g, 3, []int{4, 5})                // Node 3 connected to Node 4 and Node 5
    CreateNode(&g, 2, []int{3, 4})                // Node 2 connected to Node 3 and Node 4
    CreateNode(&g, 1, []int{2, 3})                // Node 1 connected to Node 2 and Node 3
    CreateNode(&g, 0, []int{1, 2})                // Node 0 connected to Node 1 and Node 2

    // Visualize the graph starting from Node 0
    VisualizeNode(FindNode(g, 4))
}
