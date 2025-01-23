package slice

type Slice[T any] struct {
	slice []T
}

func New[T any](slice []T) Slice[T] {
	return Slice[T]{slice: slice}
}

func (s *Slice[T]) Filter(predicate func(T) bool) {
	var result []T
	for _, item := range s.slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	s.slice = result
}
