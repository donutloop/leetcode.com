package kata

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNodeWrapper struct {
	p *TreeNode
	n *TreeNode
	l int
}

type check struct {
	p int
	l int
}

func isCousins(node *TreeNode, x int, y int) bool {
	if node == nil {
		return false
	}
	queue := list.New()

	queue.PushBack(&TreeNodeWrapper{n: node, l: 1})

	var xLevel *check
	var yLevel *check
	nodeWrapper := TreeNodeWrapper{}
	for queue.Len() > 0 {
		e := queue.Front()
		w := e.Value.(*TreeNodeWrapper)

		if yLevel != nil && xLevel != nil {
			break
		}

		if w.p != nil && w.n.Val == x {
			xLevel = &check{
				p: w.p.Val,
				l: w.l,
			}
		}
		if w.p != nil && w.n.Val == y {
			yLevel = &check{
				p: w.p.Val,
				l: w.l,
			}
		}
		if w.n.Left != nil {
			leftNode := nodeWrapper
			leftNode.p = w.n
			leftNode.n = w.n.Left
			leftNode.l = w.l + 1
			queue.PushBack(&leftNode)
		}
		if w.n.Right != nil {
			rightNode := nodeWrapper
			rightNode.p = w.n
			rightNode.n = w.n.Right
			rightNode.l = w.l + 1
			queue.PushBack(&rightNode)
		}
		queue.Remove(e)
	}

	return yLevel != nil && xLevel != nil && yLevel.l == xLevel.l && yLevel.p != xLevel.p
}
