package problem_84

type node struct {
	index  int
	height int
}

func largestRectangleArea(heights []int) int {

	stack := make([]*node, 0)
	top   := 0

	maxArea := 0
	l := len(heights)

	stack = append(stack, &node{
		index:  -1,
		height: 0,
	})

	for i, h := range heights {
		for h < stack[top].height {
			//pop
			n := stack[top]
			stack = stack[:top]
			top--
			area := n.height * (i - stack[top].index - 1)
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, &node{
			index:  i,
			height: h,
		})
		top++
	}
	for top >= 0 {
		n := stack[top]
		stack = stack[:top]
		top--
		area := n.height * (l - stack[top].index - 1)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
