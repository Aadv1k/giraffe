package giraffe

func UtilFindByRef[T comparable](elem *T, arr []*T) (*T, int) {
    for i, v := range arr {
        if v == elem {
            return v, i
        }
    }
    return nil, -1
}

func PopLast[K any](arr[]K) K {
	if len(arr) == 0 {
		return *new(K)
	}

	last := len(arr) - 1
	elem := arr[last]
	arr = append(arr[:last], arr[last+1:]...)

	return elem
}

func HasVisited(visited []*Vertex, vtx *Vertex) bool {
		for _, visitedVtx := range visited {
			if visitedVtx.Index == vtx.Index {
				return true
			}
		}

		return false
}
