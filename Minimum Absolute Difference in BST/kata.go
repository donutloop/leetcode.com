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
func getMinimumDifference(root *TreeNode) int {
	ps := make(PrioritySlice, 0)
	traverse(root, &ps)
	globalMin := -1
	for i := 1; i < len(ps); i++ {
		min := ps[i] - ps[i-1]
		if globalMin == -1 || min < globalMin {
			globalMin = min
		}

		if globalMin == 0 {
			return globalMin
		}
	}
	return globalMin
}

func traverse(node *TreeNode, ps *PrioritySlice) {
	if node == nil {
		return
	}
	*ps = append(*ps, node.Val)
	sort.Sort(ps)
	if node.Left != nil {
		traverse(node.Left, ps)
	}
	if node.Right != nil {
		traverse(node.Right, ps)
	}
}

type PrioritySlice []int

func (pq PrioritySlice) Len() int { return len(pq) }

func (pq PrioritySlice) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PrioritySlice) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
