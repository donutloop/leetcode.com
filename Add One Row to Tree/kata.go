package kata

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}
	add(root, 0, v, d-1)
	return root
}

func add(node *TreeNode, count, v, d int) {
	count = count + 1

	if count == d {
		tmp := node.Right
		node.Right = &TreeNode{Val: v}
		node.Right.Right = tmp
		tmp = node.Left
		node.Left = &TreeNode{Val: v}
		node.Left.Left = tmp
		return
	}

	if node.Right != nil {
		add(node.Right, count, v, d)
	}

	if node.Left != nil {
		add(node.Left, count, v, d)
	}
}
