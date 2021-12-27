package kata

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	list := make([]int, 0)
	inorder(root1, &list)
	inorder(root2, &list)
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
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
