package problem_239

func maxSlidingWindow(nums []int, k int) []int {

	ret := make([]int, 0)

	l := len(nums)
	if l == 0 {
		return ret
	}

	queue := make([]int, 0)
	queue  = append(queue, 0)


	var left, right int
	for right = 1; right < k; right++ {
		for len(queue) > 0 && nums[right] > nums[queue[len(queue) - 1]] {
			queue = queue[:len(queue) - 1]
		}
		queue = append(queue, right)
	}
	ret = append(ret, nums[queue[0]])
	left++
	for right < l {
		for len(queue) > 0 && nums[right] > nums[queue[len(queue) - 1]] {
			queue = queue[:len(queue) - 1]
		}
		queue = append(queue, right)
		if queue[0] < left {
			queue = queue[1:]
		}
		ret = append(ret, nums[queue[0]])
		right++
		left++
	}

	return ret
}
