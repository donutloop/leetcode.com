package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return SearchBST(root, subRoot, subRoot.Val)
}

func SearchBST(root *TreeNode, subRoot *TreeNode, val int) bool {
	if root == nil {
		return false
	}
	if root.Val == val {
		ok := LeafsEqual(root, subRoot)
		if ok {
			return true
		}
	}
	if root.Left != nil {
		ok := SearchBST(root.Left, subRoot, val)
		if ok {
			return true
		}
	}
	if root.Right != nil {
		ok := SearchBST(root.Right, subRoot, val)
		if ok {
			return true
		}
	}
	return false
}

func LeafsEqual(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	values1 := make([]NodeValue, 0)
	values2 := make([]NodeValue, 0)
	PickLeafValues(root1, &values2, 0, 0)
	PickLeafValues(root2, &values1, 0, 0)

	if len(values2) != len(values1) {
		return false
	}

	for i := 0; i < len(values1); i++ {
		if values1[i].value != values2[i].value || values1[i].lvl != values2[i].lvl || values1[i].direction != values2[i].direction {
			return false
		}
	}
	return true
}

type NodeValue struct {
	lvl       int
	value     int
	direction int
}

func PickLeafValues(node *TreeNode, values *[]NodeValue, lvl int, direction int) {
	if node == nil {
		return
	}
	*values = append(*values, NodeValue{lvl: lvl, value: node.Val, direction: direction})

	if node.Left != nil {
		PickLeafValues(node.Left, values, lvl+1, 1)
	}
	if node.Right != nil {
		PickLeafValues(node.Right, values, lvl+1, 2)
	}
}
