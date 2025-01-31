package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareList(t *testing.T) {

	// Equal
	a1 := []int{1}
	b1 := []int{1}
	assert.Equal(t, 0, CompareList(a1, b1))

	// A is smaller
	a2 := []int{1}
	b2 := []int{2}
	assert.Equal(t, -1, CompareList(a2, b2))

	// B is smaller
	a3 := []int{2}
	b3 := []int{1}
	assert.Equal(t, 1, CompareList(a3, b3))

	// A is smaller by length
	a4 := []int{1}
	b4 := []int{1, -1}
	assert.Equal(t, -1, CompareList(a4, b4))

	// B is smaller by length
	a5 := []int{1, -1}
	b5 := []int{1}
	assert.Equal(t, 1, CompareList(a5, b5))
}
