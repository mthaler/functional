package slice

func Filter[S ~[]T, T any](items S, predicate func(item T) bool) S {
	result := make(S, 0)
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
