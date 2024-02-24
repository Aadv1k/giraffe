# Giraffe 🦒 

experimental graphing library in go

## Usage

> The public API doesn't really exist as of now as a lot of the base stuff is being re-built every
> commit, consider running the tests or [taking the look at the source](./giraffe/giraffe.go)

```shell
go test .\giraffe\
```

## Example

```go
func main() {
    var g giraffe.Graph

    v0 := &giraffe.Vertex{Index: 0}
    v1 := &giraffe.Vertex{Index: 1}

    g.AddVertex(v0)
    g.AddVertex(v1)

    g.AddEdge(&giraffe.Edge{Start: v0, End: v1})
    
    found, visited := g.FindEdge(v1)

    fmt.Printf("Found: %d\n")

    fmt.Print("Visited: {")
    for _, vtx := range visited {
      fmt.Printf("%d, ", vtx.Index)
    }
    fmt.Println("}")
}
```
