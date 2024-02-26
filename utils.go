package giraffe

func UtilFindByRef[T comparable](elem *T, arr []*T) (*T, int) {
    for i, v := range arr {
        if v == elem {
            return v, i
        }
    }
    return nil, -1
}
