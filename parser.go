package giraffe

import (
	"fmt"
	"strings"
)

func (g *Graph) ParseToMermaid() string {
    var builder strings.Builder
    builder.WriteString("graph TD;\n")

    for _, vertex := range g.Vertices {
        builder.WriteString(fmt.Sprintf("    %d;\n", vertex.Index))
    }

    for _, edge := range g.Edges {
        builder.WriteString(fmt.Sprintf("    %d --> %d;\n", edge.Start.Index, edge.End.Index))
    }

    return builder.String()
}

func (g *Graph) ParseToGraphvizDot() string {
    var builder strings.Builder

    builder.WriteString("digraph G {\n")

    for _, vertex := range g.Vertices {
        builder.WriteString(fmt.Sprintf("    %d;\n", vertex.Index))
    }

    for _, edge := range g.Edges {
        builder.WriteString(fmt.Sprintf("    %d -> %d;\n", edge.Start.Index, edge.End.Index))
    }

    builder.WriteString("}")
    return builder.String()
}
