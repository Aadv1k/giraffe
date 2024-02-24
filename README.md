# Giraffe 🦒 

experimental graphing library in go

## Usage

```shell
go run .
```

## Example

```go
// ...

func main() {
    var g Graph

    CreateNode(&g, 5, []int{})                    // Node 5
    CreateNode(&g, 4, []int{})                    // Node 4
    CreateNode(&g, 3, []int{4, 5})                // Node 3 connected to Node 4 and Node 5
    CreateNode(&g, 2, []int{3, 4})                // Node 2 connected to Node 3 and Node 4
    CreateNode(&g, 1, []int{2, 3})                // Node 1 connected to Node 2 and Node 3
    CreateNode(&g, 0, []int{1, 2})                // Node 0 connected to Node 1 and Node 2

    n0 := FindNode(&g, 0) // returns pointer to an instance of Node
    
    // Recursively traverse, and pretty-print the nodes
    VisualizeNode(n0)
}
```