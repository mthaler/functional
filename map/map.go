package mypap

type MyMap[K comparable, V any] struct {
	mymap map[K]V
}

func New[K comparable, V any](m map[K]V) MyMap[K, V] {
	return MyMap[K, V]{mymap: m}
}

func (m *MyMap[K, V]) Filter(predicate func(K, V) bool) {
	var result map[K]V
	for key, value := range m.mymap {
		if predicate(key, value) {
			result[key] = value
		}
	}
	m.mymap = result
}
