package kata

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	paths := make([]string, 0)
	allPaths(root, strconv.Itoa(root.Val)+"->", &paths)
	return paths
}

func allPaths(node *TreeNode, s string, all *[]string) {
	if node.Left == nil && node.Right == nil {
		*all = append(*all, s[:len(s)-2])
		return
	}
	if node.Left != nil {
		allPaths(node.Left, s+strconv.Itoa(node.Left.Val)+"->", all)
	}
	if node.Right != nil {
		allPaths(node.Right, s+strconv.Itoa(node.Right.Val)+"->", all)
	}
}
