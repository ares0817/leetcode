package problem_239

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {

	nums := []int{1,3,-1,-3,5,3,6,7}
	k    := 3

	ret := maxSlidingWindow(nums, k)
	fmt.Println(ret)
}
