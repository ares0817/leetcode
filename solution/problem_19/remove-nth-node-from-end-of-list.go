package problem_19

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	back := false
	n++
	recRemoveNthFromEnd(head, &n, &back)
	if n == 1 {
		return head.Next
	}
	return head
}

func recRemoveNthFromEnd(node *ListNode, n *int, back *bool) {

	if node.Next == nil {
		*back = true
	}

	if !*back {
		recRemoveNthFromEnd(node.Next, n, back)
	}
	*n--
	if *n == 0 {
		node.Next = node.Next.Next
	}
}
