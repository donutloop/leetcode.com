package _73__Binary_Search_Tree_Iterator

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []int
}

func Constructor(root *TreeNode) BSTIterator {
	stack := make([]int, 0)
	populateStack(root, &stack)
	sort.Slice(stack, func(i, j int) bool {
		return stack[i] < stack[j]
	})
	return BSTIterator{
		stack: stack,
	}
}

func populateStack(node *TreeNode, stack *[]int) {
	if node == nil {
		return
	}

	*stack = append(*stack, node.Val)

	if node.Left != nil {
		populateStack(node.Left, stack)
	}
	if node.Right != nil {
		populateStack(node.Right, stack)
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	v := this.stack[0]
	this.stack = this.stack[1:]
	return v
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.stack) == 0 {
		return false
	}
	return true
}
