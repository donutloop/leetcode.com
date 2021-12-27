package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	for _, v := range preorder[1:] {
		build(root, v)
	}
	return root
}

func build(node *TreeNode, v int) {
	if node.Val > v {
		if node.Left == nil {
			node.Left = &TreeNode{Val: v}
		} else {
			build(node.Left, v)
		}
	} else if node.Val < v {
		if node.Right == nil {
			node.Right = &TreeNode{Val: v}
		} else {
			build(node.Right, v)
		}
	}
}
