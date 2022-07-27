package kata

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	_, ok := p(head, head)
	return ok
}

func p(a *ListNode, b *ListNode) (*ListNode, bool) {
	if a.Next == nil {
		if a.Val != b.Val {
			return nil, false
		}
		return b.Next, true
	}

	c, ok := p(a.Next, b)
	if !ok {
		return nil, false
	}

	if c.Val == a.Val {
		return c.Next, true
	}
	return nil, false
}
