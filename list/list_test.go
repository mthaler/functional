package mylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	l := New(1, 2, 3, 4, 5, 6)
	l2 := l.Filter(func(item any) bool {
		i := item.(int)
		return i%2 == 0
	})
	assert.Equal(t, New(2, 4, 6), l2)
}

func TestMap(t *testing.T) {
	l := New(1, 2, 3, 4, 5, 6)
	l2 := l.Map(func(i any) any {
		j := i.(int)
		return j * j
	})
	assert.Equal(t, New(1, 4, 9, 16, 25, 36), l2)
}
