package _29__Sum_Root_to_Leaf_Numbers

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sums := make([]string, 0)
	f(root, &sums, "")
	var max int
	for _, s := range sums {
		sum, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		max = max + sum
	}
	return max
}

func f(node *TreeNode, sums *[]string, n string) {
	val := strconv.Itoa(node.Val)
	n = n + val

	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, n)
		return
	}

	if node.Left != nil {
		f(node.Left, sums, n)
	}

	if node.Right != nil {
		f(node.Right, sums, n)
	}
}
