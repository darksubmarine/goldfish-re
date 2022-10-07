package goldfish_re

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_utils_growthSlice(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	assert.Len(t, s1, 5)

	assert.Panics(t, func() {
		s1[6] = 6
	})

	oldSize := len(s1)
	increaseSize := 10
	s1 = growthSlice[int](s1, increaseSize)
	assert.Len(t, s1, oldSize+increaseSize)

}
