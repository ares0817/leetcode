package problem_11

/**
 * 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 * https://leetcode-cn.com/problems/container-with-most-water/
 */
func maxArea(height []int) int {

	max := 0
	i := 0
	j := len(height) - 1

	for {
		if i >= j {
			break
		}

		m, l := min(height[i], height[j])
		s := (j - i) * m
		if s > max {
			max = s
		}

		if l == 1 {
			j--
		} else {
			i++
		}
	}

	return max
}

func min(a int, b int) (int, int) {
	if a < b {
		return a, 0
	} else {
		return b, 1
	}
}
