package kata

import (
	"container/list"
)

func calculate(s string) int {
	nodes := Parse([]byte(s))
	return Eval(nodes)
}

const (
	NumberKind = iota
	PlusKind
	MinusKind
	MultiplyKind
	DvisionKind
)

type Node struct {
	Identifier byte
	Kind       int
	Value      int
	Priority   int
}

const (
	Multiply byte = 42
	Dvision  byte = 47
	Plus     byte = 43
	Minus    byte = 45
)

func Parse(expression []byte) *list.List {

	nodes := list.New()
	for i := len(expression) - 1; i >= 0; i-- {
		if expression[i] == Minus {
			nodes.PushBack(Node{Kind: MinusKind})
		} else if expression[i] == Multiply {
			nodes.PushBack(Node{Kind: MultiplyKind, Priority: 1})
		} else if expression[i] == Plus {
			nodes.PushBack(Node{Kind: PlusKind})
		} else if expression[i] == Dvision {
			nodes.PushBack(Node{Kind: DvisionKind, Priority: 1})
		} else if expression[i] >= 48 && expression[i] <= 57 {
			n := Node{Kind: NumberKind}
			var sum int
			sum = int(expression[i]) - 48
			i--
			norm := 10
			for i >= 0 {
				if expression[i] >= 48 && expression[i] <= 57 {
					sum += (int(expression[i]) - 48) * norm
					i--
					norm = norm * 10
				} else {
					i++
					break
				}
			}
			n.Value = sum
			nodes.PushBack(n)
		}
	}

	return nodes
}

func Eval(nodes *list.List) int {
	if nodes.Len() == 0 {
		return 0
	}

	for e := nodes.Back(); e != nil; e = e.Prev() {
		node := e.Value.(Node)
		if node.Priority == 1 {
			if node.Kind == MultiplyKind {
				eA := e.Prev()
				nodeA := eA.Value.(Node)
				eb := e.Next()
				nodeB := eb.Value.(Node)
				sum := nodeA.Value * nodeB.Value

				nodes.Remove(eb)
				nodes.Remove(eA)
				newNode := nodes.InsertBefore(Node{Kind: NumberKind, Value: sum}, e)
				nodes.Remove(e)
				e = newNode
			}
			if node.Kind == DvisionKind {
				eA := e.Prev()
				nodeA := eA.Value.(Node)
				eb := e.Next()
				nodeB := eb.Value.(Node)
				sum := nodeB.Value / nodeA.Value

				nodes.Remove(eb)
				nodes.Remove(eA)
				newNode := nodes.InsertBefore(Node{Kind: NumberKind, Value: sum}, e)
				nodes.Remove(e)
				e = newNode
			}
		}
	}

	value := nodes.Back()
	sum := value.Value.(Node).Value
	value = value.Prev()
	for e := value; e != nil; e = e.Prev() {
		node := e.Value.(Node)
		if node.Kind == PlusKind {
			e = e.Prev()
			node = e.Value.(Node)
			sum += node.Value

		} else if node.Kind == MinusKind {
			e = e.Prev()
			node = e.Value.(Node)
			sum -= node.Value
		}
	}
	return sum
}
