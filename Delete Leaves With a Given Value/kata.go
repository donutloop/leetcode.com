package kata

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	for {
		ok := removeLeafNode(nil, root, target, RootNodeType)
		if !ok {
			if root.Right == nil && root.Left == nil && root.Val == target {
				return nil
			}
			return root
		}
	}
}

type NodeType int

const (
	RootNodeType  NodeType = 0
	LeftNodeType  NodeType = 1
	RightNodeType NodeType = 2
)

func removeLeafNode(parentNode *TreeNode, node *TreeNode, target int, nodeType NodeType) bool {
	if node.Right == nil && node.Left == nil && node.Val == target && parentNode != nil {
		if nodeType == LeftNodeType {
			parentNode.Left = nil
		} else {
			parentNode.Right = nil
		}
		return true
	}

	var left bool
	if node.Left != nil {
		left = removeLeafNode(node, node.Left, target, LeftNodeType)
	}
	var right bool
	if node.Right != nil {
		right = removeLeafNode(node, node.Right, target, RightNodeType)
	}
	return right == true || left == true
}
