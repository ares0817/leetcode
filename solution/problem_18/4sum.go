package problem_18

import (
	"sort"
	"strconv"
)

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)

	l := len(nums)

	res := make([][]int, 0)
	m := make(map[string]bool, 0)

	for i := 0; i < l - 3; i++ {
		for j := i + 1; j < l - 2; j++ {
			t := target - nums[i] - nums[j]
			left  := j + 1
			right := l - 1
			for left < right {
				if nums[left] > t / 2 {
					break
				}
				s := nums[left] + nums[right]
				if s == t {
					key := getKey(nums[i], nums[j], nums[left], nums[right])
					if _, exist := m[key]; !exist {
						res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
						m[key] = true
					}
					left++
					right--
				} else if s < t {
					left++
				} else {
					right--
				}
			}
		}
	}

	return res
}

func getKey(a1, a2, a3, a4 int) string {
	str := strconv.Itoa(a1) + "_" + strconv.Itoa(a2) + "_" + strconv.Itoa(a3) + "_" + strconv.Itoa(a4)
	return str
}