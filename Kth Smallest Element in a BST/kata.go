package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	list := make([]int, 0)
	inorder(root, &list, k)
	return list[k-1]
}

func inorder(node *TreeNode, list *[]int, k int) {
	if k == len(*list) {
		return
	}
	if node == nil {
		return
	}
	if node.Left != nil {
		inorder(node.Left, list, k)
	}
	*list = append(*list, node.Val)
	if node.Right != nil {
		inorder(node.Right, list, k)
	}
}
