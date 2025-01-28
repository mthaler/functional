package mylist

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	l := New(1, 2, 3, 4, 5, 6)
	l2 := l.Filter(func(item any) bool {
		i := item.(int)
		return i%2 == 0
	})
	exptected := New(2, 4, 6)
	assert.Equal(t, exptected, l2)
}

func toSlice(l MyList) []int {
	var result []int
	li := list.List(*l)
	for e := li.Front(); e != nil; e = e.Next() {
		v := e.Value
		result = append(result, v)
	}
	return result
}
