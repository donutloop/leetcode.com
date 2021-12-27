package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	values1 := make([]int, 0)
	values2 := make([]int, 0)
	PickLeafValues(root1, &values2)
	PickLeafValues(root2, &values1)
	if len(values2) != len(values1) {
		return false
	}
	for i := 0; i < len(values1); i++ {
		if values1[i] != values2[i] {
			return false
		}
	}
	return true
}

func PickLeafValues(node *TreeNode, values *[]int) {
	if node != nil && node.Left == nil && node.Right == nil {
		*values = append(*values, node.Val)
		return
	}
	if node == nil {
		return
	}
	if node.Left != nil {
		PickLeafValues(node.Left, values)
	}
	if node.Right != nil {
		PickLeafValues(node.Right, values)
	}
}
