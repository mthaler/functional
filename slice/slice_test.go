package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiler(t *testing.T) {
	s := New[int]([]int{1, 2, 3, 4, 5, 6})
	s.Filter(func(i int) bool {
		return i%2 == 0
	})
	assert.Equal(t, []int{2, 4, 6}, s.slice)
}
