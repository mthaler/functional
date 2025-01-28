package mylist

import (
	"container/list"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	l := New(1, 2, 3, 4, 5, 6)
	l2 := l.Filter(func(item any) bool {
		i := item.(int)
		return i%2 == 0
	})
	s2 := l2.toSlice()
	s2 = sort.Ints(s2)
	exptected := New(2, 4, 6)
	assert.Equal(t, exptected.toSlice(), l2)
}

func (l *MyList) toSlice() []int {
	var result []int
	li := list.List(*l)
	for e := li.Front(); e != nil; e = e.Next() {
		v := e.Value
		result = append(result, v.(int))
	}
	return result
}
