package problem_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {

	assert.Equal(t, -1, myAtoi("-000000000000001"))
	assert.Equal(t, 2147483647, myAtoi("111111111111"))
	assert.Equal(t, -2147483648, myAtoi("-111111111111"))
	assert.Equal(t, 0, myAtoi("abcd"))
	assert.Equal(t, -123, myAtoi("     -123aabc"))
	assert.Equal(t, 123, myAtoi("  123avc45"))
	assert.Equal(t, 123, myAtoi("  +123avc45"))
	assert.Equal(t, 123, myAtoi("  00000000000000123avc45"))
}
