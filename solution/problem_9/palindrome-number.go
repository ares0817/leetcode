package problem_9

import (
	"strconv"
)

func isPalindrome(x int) bool {

	str := strconv.Itoa(x)
	i, j := 0, len(str) - 1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindrome2(x int) bool {

	if x < 0 {
		return false
	}

	x1  := x
	num := 0
	for x > 0 {
		num = num * 10 + x % 10
		x = x / 10
	}

	return x1 == num
}
