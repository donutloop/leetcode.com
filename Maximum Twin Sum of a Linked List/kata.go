package kata

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	max := new(int)
	f(head, head, max)
	return *max
}

func f(a *ListNode, b *ListNode, max *int) *ListNode {
	if a.Next == nil {
		sum := a.Val + b.Val
		if sum > *max {
			*max = sum
		}
		return b.Next
	}

	c := f(a.Next, b, max)
	sum := a.Val + c.Val
	if sum > *max {
		*max = sum
	}
	return c.Next
}
