package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type leftLeaf struct {
	val   int
	level int
	kind  int
}

func findBottomLeftValue(root *TreeNode) int {
	leafs := make([]leftLeaf, 0)
	findNodes(root, 0, &leafs, 1)
	max := -1
	val := 0
	for i := 0; i < len(leafs); i++ {
		if max < leafs[i].level && leafs[i].kind == 1 {
			max = leafs[i].level
			val = leafs[i].val
		} else if max < leafs[i].level && leafs[i].kind == 2 {
			max = leafs[i].level
			val = leafs[i].val
		}
	}
	return val
}

func findNodes(node *TreeNode, level int, leafs *[]leftLeaf, kind int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*leafs = append(*leafs, leftLeaf{val: node.Val, kind: kind, level: level})
		return
	}
	level = level + 1
	if node.Left != nil {
		findNodes(node.Left, level, leafs, 1)
	}
	if node.Right != nil {
		findNodes(node.Right, level, leafs, 2)
	}
}
