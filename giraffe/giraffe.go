package giraffe

import (
    "fmt"
)

type Node struct {
    index          int
    total_siblings int
    siblings       []*Node
}

type Graph []*Node

// Find a particular node
func FindNode(g Graph, index int) *Node {
    for _, node := range g {
        if node.index == index {
            return node
        }
    }
    return nil
}

// Create node with or without children
func CreateNode(g *Graph, index int, siblings []int) *Node {
    if foundNode := FindNode(*g, index); foundNode != nil {
        fmt.Printf("A node already exists with index %d\n", foundNode.index)
        return foundNode
    }

    node := &Node{index: index, siblings: []*Node{}}

    for _, siblingIndex := range siblings {
        if foundNode := FindNode(*g, siblingIndex); foundNode == nil {
            fmt.Printf("Sibling node %d doesn't exist!\n", siblingIndex)
            continue
        }
        sibling := FindNode(*g, siblingIndex)
        node.siblings = append(node.siblings, sibling)
    }

    *g = append(*g, node)
    return node
}

func VisualizeNode(n *Node) {
    fmt.Printf("%d\n", n.index)
    for _, sibling := range n.siblings {
        VisualizeNodeSiblings(sibling, "", true)
    }
}

func VisualizeNodeSiblings(n *Node, prefix string, directChild bool) {
    if !directChild {
        fmt.Printf("%s└───%d\n", prefix, n.index)
    }  else {
        fmt.Printf("└───%d\n", n.index)

    }


    for _, sibling := range n.siblings {
        if directChild {
            VisualizeNodeSiblings(sibling, prefix+"    ", false)
            continue
        }
        VisualizeNodeSiblings(sibling, prefix+"|   ", false)
    }
}
