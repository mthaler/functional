package slice

type Slice[T any] struct {
	slice []T
}

func New[T any](slice []T) Slice[T] {
	return Slice[T]{slice: slice}
}

func (s *Slice[T]) Filter(predicate func(T) bool) Slice[T] {
	var result []T
	for _, item := range s.slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return New[T](result)
}

func (s *Slice[T]) Map(f func(T) T) {
	var result []T
	for _, item := range s.slice {
		result = append(result, f(item))
	}
	s.slice = result
}
