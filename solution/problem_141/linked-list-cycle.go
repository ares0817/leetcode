package problem_141

/**
 * 给定一个链表，判断链表中是否有环。
 * https://leetcode-cn.com/problems/linked-list-cycle/
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	p1, p2 := head, head

	for {
		if p2.Next == nil || p2.Next.Next == nil {
			return false
		}
		p1, p2 = p1.Next, p2.Next.Next
		if p1 == p2 {
			return true
		}
	}
}
