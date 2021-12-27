package kata

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)
	traversal(root, &list)
	return list
}

func traversal(node *TreeNode, list *[]int) {
	if node == nil {
		return
	}
	*list = append(*list, node.Val)
	if node.Left != nil {
		traversal(node.Left, list)
	}
	if node.Right != nil {
		traversal(node.Right, list)
	}
}
