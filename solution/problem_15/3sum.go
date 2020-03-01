package problem_15

import "sort"

/**
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
 * 注意：答案中不可以包含重复的三元组。
 * https://leetcode-cn.com/problems/3sum/
 */
func threeSum(nums []int) [][]int {

	arrRet := make([][]int, 0)

	lenA := len(nums)
	if lenA < 3 {
		return arrRet
	}

	sort.Ints(nums)

	i := 0;
	for {
		target := 0 - nums[i]
		j := i + 1
		k := lenA - 1
		if nums[j] > target / 2 {
			break
		}
		for j < k {
			if nums[j] + nums[k] == target {
				arrRet = append(arrRet, []int{nums[i], nums[j], nums[k]})
				k = decrK(k, nums)
				j = incrJ(j, nums, lenA)
			} else if nums[j] + nums[k] < target {
				j = incrJ(j, nums, lenA)
			} else {
				k = decrK(k, nums)
			}
		}
		i = incrI(i, nums, lenA)
		if i >= lenA - 2 {
			break
		}
	}

	return arrRet
}

func incrI(i int, A[]int, lenA int) int {

	i++
	for i < lenA - 2 && A[i] == A[i - 1] {
		i++
	}

	return i
}

func incrJ(j int, A []int, lenA int) int {

	j++
	for j < lenA - 1 && A[j] == A[j - 1] {
		j++
	}

	return j
}

func decrK(k int, A []int) int {

	k--
	for k > 0 && A[k] == A[k + 1] {
		k--
	}

	return k
}
