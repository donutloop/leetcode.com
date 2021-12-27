package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	list := inorderTraversal(root)
	for i := 1; i < len(list); i++ {
		if !(list[i-1] < list[i]) {
			return false
		}
	}
	return true
}

func inorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)
	inorder(root, &list)
	return list
}

func inorder(node *TreeNode, list *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		inorder(node.Left, list)
	}
	*list = append(*list, node.Val)
	if node.Right != nil {
		inorder(node.Right, list)
	}
}
