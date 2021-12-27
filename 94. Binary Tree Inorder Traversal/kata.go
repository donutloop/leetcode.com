package _4__Binary_Tree_Inorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)
	traversal(root, &list)
	return list
}

func traversal(node *TreeNode, list *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		traversal(node.Left, list)
	}
	*list = append(*list, node.Val)
	if node.Right != nil {
		traversal(node.Right, list)
	}
}
