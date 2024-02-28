package giraffe

import (
    "testing"
		"fmt"
)

func TestParseToMermaid(t *testing.T) {
    graph := Graph{
        Vertices: []*Vertex{
            {Index: 1},
            {Index: 2},
            {Index: 3},
        },
        Edges: []*Edge{
            {Start: &Vertex{Index: 1}, End: &Vertex{Index: 2}},
            {Start: &Vertex{Index: 2}, End: &Vertex{Index: 3}},
        },
    }

    expected := "graph TD;\n    1;\n    2;\n    3;\n    1 --> 2;\n    2 --> 3;\n"
    result := graph.ParseToMermaid()

    if result != expected {
        t.Errorf("ParseToMermaid() returned incorrect result, got: %s, want: %s", result, expected)
    }
}
