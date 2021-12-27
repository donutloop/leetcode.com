package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	nodes := inorderTraversal(root)
	tmp := &TreeNode{}
	current := tmp
	for _, n := range nodes {
		if current.Right == nil {
			current.Right = &TreeNode{Val: n}
		}
		current = current.Right
	}
	return tmp.Right
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
