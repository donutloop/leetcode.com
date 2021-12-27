package kata

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
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
	if node.Right != nil {
		traversal(node.Right, list)
	}
	*list = append(*list, node.Val)
}
