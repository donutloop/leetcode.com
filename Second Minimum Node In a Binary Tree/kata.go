package kata

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	list := make(map[int]struct{})
	inorder(root, list)
	if len(list) == 1 {
		return -1
	}
	a := make([]int, 0)
	for n := range list {
		a = append(a, n)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a[1]
}

func inorder(node *TreeNode, list map[int]struct{}) {
	if node == nil {
		return
	}

	if node.Left != nil {
		inorder(node.Left, list)
	}

	list[node.Val] = struct{}{}

	if node.Right != nil {
		inorder(node.Right, list)
	}
}
