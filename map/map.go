package mypap

type MyMap[K comparable, V any] struct {
	mymap map[K]V
}

func New[K comparable, V any](m map[K]V) MyMap[K, V] {
	return MyMap[K, V]{mymap: m}
}

func (m *MyMap[K, V]) Filter(predicate func(key K, value V) bool) MyMap[K, V] {
	result := make(map[K]V)
	for key, value := range m.mymap {
		if predicate(key, value) {
			result[key] = value
		}
	}
	return New[K, V](result)
}

func (m *MyMap[K, V]) Map(f func(key K, value V) (K, V)) MyMap[K, V] {
	result := make(map[K]V)
	for key, value := range m.mymap {
		k, v := f(key, value)
		result[k] = v
	}
	return New[K, V](result)
}
