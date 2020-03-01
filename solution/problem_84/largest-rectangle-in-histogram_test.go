package problem_84

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {

	heights := []int{2, 1, 5, 6, 2, 3}
	assert.Equal(t, 10, largestRectangleArea(heights))
}
