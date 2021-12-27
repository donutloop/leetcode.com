package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	c := make(map[int]int)
	nums := make([]int, 0)
	collectModes(root, c)
	max := -1
	for _, n := range c {
		if max == -1 || n >= max {
			max = n
		}
	}
	for num, counter := range c {
		if counter == max {
			nums = append(nums, num)
		}
	}
	return nums
}

func collectModes(node *TreeNode, c map[int]int) {
	if node == nil {
		return
	}
	c[node.Val]++

	if node.Left != nil {
		collectModes(node.Left, c)
	}
	if node.Right != nil {
		collectModes(node.Right, c)
	}
}
