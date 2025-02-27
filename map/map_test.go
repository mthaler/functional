package mypap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	var m = map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}
	m2 := New[int, int](m)
	m3 := m2.Filter(func(key int, value int) bool {
		return key%2 == 0
	})
	var expected = map[int]int{2: 2, 4: 4, 6: 6}
	assert.Equal(t, expected, m3.mymap)
}

func TestMap(t *testing.T) {
	var m = map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}
	m2 := New[int, int](m)
	m3 := m2.Map(func(key int, value int) (int, int) {
		return key, value * value
	})
	var m4 = map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25, 6: 36}
	assert.Equal(t, m4, m3.mymap)
}
