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
	exptected := New(2, 4, 6)
	assert.Equal(t, exptected, l2)
}
