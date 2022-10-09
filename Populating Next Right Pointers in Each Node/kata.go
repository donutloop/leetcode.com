package Populating_Next_Right_Pointers_in_Each_Node

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	f(root.Left, root.Right)
	return root
}

func f(l *Node, r *Node) {
	if l == nil || r == nil {
		return
	}
	l.Next = r
	f(r.Left, r.Right)
	f(l.Right, r.Left)
	f(l.Left, l.Right)
	return
}
