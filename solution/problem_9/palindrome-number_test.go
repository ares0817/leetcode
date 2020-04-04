package problem_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome2(t *testing.T) {

	assert.Equal(t, true, isPalindrome2(121))
	assert.Equal(t, true, isPalindrome2(12344321))
	assert.Equal(t, false, isPalindrome2(-121))
	assert.Equal(t, false, isPalindrome2(12))
}
