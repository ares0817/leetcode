package problem_13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {

	assert.Equal(t, 1476, romanToInt("MCDLXXVI"))
	assert.Equal(t, 3, romanToInt("III"))
	assert.Equal(t, 4, romanToInt("IV"))
	assert.Equal(t, 9, romanToInt("IX"))
	assert.Equal(t, 58, romanToInt("LVIII"))
	assert.Equal(t, 1994, romanToInt("MCMXCIV"))
}
