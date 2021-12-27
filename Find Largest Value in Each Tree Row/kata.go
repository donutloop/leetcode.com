package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	set := make(map[int]*largest)
	findAll(root, set, 0)
	a := make([]int, len(set))
	var i int
	for _, v := range set {
		a[v.Level] = v.Val
		i++
	}
	return a
}

type largest struct {
	Val   int
	Level int
}

func findAll(node *TreeNode, set map[int]*largest, l int) {
	if node == nil {
		return
	}
	v, ok := set[l]
	if ok {
		if node.Val >= v.Val {
			v.Val = node.Val
			set[l] = v
		}
	} else {
		set[l] = &largest{Val: node.Val, Level: l}
	}
	if node.Left != nil {
		findAll(node.Left, set, l+1)
	}
	if node.Right != nil {
		findAll(node.Right, set, l+1)
	}
}
