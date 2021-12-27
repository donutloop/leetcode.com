package kata

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(node *TreeNode) []float64 {
	if node == nil {
		return nil
	}
	queue := list.New()
	type TreeNodeWrapper struct {
		n *TreeNode
		l int
	}
	queue.PushBack(&TreeNodeWrapper{n: node, l: 1})
	vals := make([]float64, 0)
	var sum float64
	lastLevel := 1
	var count float64
	for queue.Len() > 0 {
		e := queue.Front()
		w := e.Value.(*TreeNodeWrapper)
		if w.n.Left != nil {
			queue.PushBack(&TreeNodeWrapper{n: w.n.Left, l: w.l + 1})
		}
		if w.n.Right != nil {
			queue.PushBack(&TreeNodeWrapper{n: w.n.Right, l: w.l + 1})
		}
		if w.l != lastLevel {
			vals = append(vals, sum/count)
			sum = 0
			count = 0
			lastLevel = w.l
		}
		count = count + 1
		sum = sum + float64(w.n.Val)
		queue.Remove(e)
		if queue.Len() == 0 {
			vals = append(vals, sum/count)
		}
	}
	return vals
}
