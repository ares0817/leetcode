package problem_16

func threeSumClosest(nums []int, target int) int {

	closest := 0
	min := int(^uint(0) >> 1)
	l := len(nums)

	for i := 0; i < l - 2; i++ {
		for j := i + 1; j < l - 1; j++ {
			for k := j + 1; k < l; k++ {
				sum := nums[i] + nums[j] + nums[k]
				gap := abs(sum - target)
				if gap < min {
					closest = sum
					min = gap
				}
			}
		}
	}

	return closest
}

func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}