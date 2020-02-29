package solution

/**
 * https://leetcode-cn.com/problems/move-zeroes/
 */
func moveZeroes(nums []int)  {

	l := len(nums)
	j := -1
	//j is the index of the first 0
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			j = i
			break
		}
	}

	//0 not found or 0 is the last elem, return
	if j == -1 || j == l - 1 {
		return
	}

	i := j + 1
	for i < l {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
		i++
	}
}
