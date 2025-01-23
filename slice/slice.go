package slice

type Slice[T any] struct {
	Slice []T
}

func New[T any](slice []T) Slice[T] {
	return Slice[T]{Slice: slice}
}

func (s Slice[T]) Filter(predicate func(T) bool) Slice[T] {
	var result []T
	for _, item := range s.Slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return Slice[T]{Slice: result}
}
