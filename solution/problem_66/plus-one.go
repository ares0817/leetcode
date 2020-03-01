package problem_66

/**
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 * https://leetcode-cn.com/problems/plus-one/
 */
func plusOne(digits []int) []int {

	l := len(digits)
	i := 0
	for i < l {
		if digits[i] != 9 {
			break
		}
		i++
	}

	//all num is 9
	if i == l {
		//add a zero
		digits = append(digits, 0)
		//first is 1
		digits[0] = 1
		//other is 0
		for i = 1; i <= l; i++ {
			digits[i] = 0
		}
		return digits
	}

	i = l - 1
	for i >= 0 {
		digits[i] = digits[i] + 1
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
		i--
	}

	return digits
}