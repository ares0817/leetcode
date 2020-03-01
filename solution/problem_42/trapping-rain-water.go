package problem_42

type node struct {
	h   int
	max int
}

func trap(height []int) int {

	l, r := 0, len(height) - 1
	
	if r < 0 {
		return 0
	}

	for l < r {
		if height[l + 1] >= height[l] {
			l++
		} else {
			break
		}
	}
	for l < r {
		if height[r - 1] >= height[r] {
			r--
		} else {
			break
		}
	}

	ret := 0

	q := make([]*node, 0)
	q = append(q, &node{
		h:   height[l],
		max: height[l],
	})
	l++
	for l <= r {
		if height[l] < q[len(q) - 1].max {
			q = append(q, &node{
				h:   height[l],
				max: q[len(q) - 1].max,
			})
		} else {
			for len(q) > 0 {
				n := q[len(q)-1]
				q = q[:len(q)-1]
				ret += n.max - n.h
			}
			q = append(q, &node{
				h:   height[l],
				max: height[l],
			})
		}
		l++
	}
	if len(q) > 0 {
		last := q[len(q)-1]
		q = q[:len(q)-1]
		for len(q) > 0 {
			n := q[len(q)-1]
			q = q[:len(q)-1]
			if n.h < last.h {
				ret += last.h - n.h
			} else {
				last = n
			}
		}
	}

	return ret
}
