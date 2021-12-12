package _50__Evaluate_Reverse_Polish_Notation

import "container/list"

func evalRPN(tokens []string) int {
	nodes := Parse(tokens)
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
	Kind  int
	Value int
}

const (
	Multiply byte = 42
	Dvision  byte = 47
	Plus     byte = 43
	Minus    byte = 45
)

func Parse(tokens []string) *list.List {

	nodes := list.New()
	for i := 0; i < len(tokens); i++ {
		if len(tokens[i]) == 1 && tokens[i][0] == Minus {
			nodes.PushFront(Node{Kind: MinusKind})
		} else if len(tokens[i]) == 1 && tokens[i][0] == Multiply {
			nodes.PushFront(Node{Kind: MultiplyKind})
		} else if len(tokens[i]) == 1 && tokens[i][0] == Plus {
			nodes.PushFront(Node{Kind: PlusKind})
		} else if len(tokens[i]) == 1 && tokens[i][0] == Dvision {
			nodes.PushFront(Node{Kind: DvisionKind})
		} else if len(tokens[i]) == 1 && tokens[i][0] >= 48 && tokens[i][0] <= 57 {
			n := Node{Kind: NumberKind}
			var sum int
			sum = int(tokens[i][0]) - 48
			n.Value = sum
			nodes.PushFront(n)
		} else if len(tokens[i]) == 2 && tokens[i][0] == Minus {
			n := Node{Kind: NumberKind}
			var sum int
			sum = int(tokens[i][1]) - 48
			n.Value = sum * -1
			nodes.PushFront(n)
		} else if len(tokens[i]) > 1 && tokens[i][0] >= 48 && tokens[i][0] <= 57 {
			n := Node{Kind: NumberKind}
			var sum int
			j := len(tokens[i]) - 1
			norm := 1
			for j >= 0 {
				if tokens[i][j] >= 48 && tokens[i][j] <= 57 {
					sum += (int(tokens[i][j]) - 48) * norm
					j--
					norm = norm * 10
				}
			}
			n.Value = sum
			nodes.PushFront(n)
		} else if len(tokens[i]) > 2 && tokens[i][0] == Minus {
			n := Node{Kind: NumberKind}
			var sum int
			j := len(tokens[i]) - 1
			norm := 1
			for j >= 1 {
				if tokens[i][j] >= 48 && tokens[i][j] <= 57 {
					sum += (int(tokens[i][j]) - 48) * norm
					j--
					norm = norm * 10
				}
			}
			n.Value = sum * -1
			nodes.PushFront(n)
		}

	}

	return nodes
}

func Eval(nodes *list.List) int {
	if nodes.Len() == 0 {
		return 0
	}

	for nodes.Len() != 1 {
		doEval(nodes)
	}

	value := nodes.Back()
	sum := value.Value.(Node).Value
	return sum
}

func doEval(nodes *list.List) {
	for e := nodes.Back(); e != nil; e = e.Prev() {
		node := e.Value.(Node)
		if node.Kind == MultiplyKind {
			eA := e.Next()
			nodeA := eA.Value.(Node)
			eb := eA.Next()
			nodeB := eb.Value.(Node)
			sum := nodeB.Value * nodeA.Value

			nodes.Remove(eb)
			nodes.Remove(eA)
			nodes.InsertBefore(Node{Kind: NumberKind, Value: sum}, e)
			nodes.Remove(e)
			return
		} else if node.Kind == DvisionKind {
			eA := e.Next()
			nodeA := eA.Value.(Node)
			eB := eA.Next()
			nodeB := eB.Value.(Node)
			sum := nodeB.Value / nodeA.Value

			nodes.Remove(eB)
			nodes.Remove(eA)
			nodes.InsertBefore(Node{Kind: NumberKind, Value: sum}, e)
			nodes.Remove(e)
			return
		} else if node.Kind == PlusKind {
			eA := e.Next()
			nodeA := eA.Value.(Node)
			eb := eA.Next()
			nodeB := eb.Value.(Node)
			sum := nodeB.Value + nodeA.Value
			nodes.Remove(eb)
			nodes.Remove(eA)
			nodes.InsertBefore(Node{Kind: NumberKind, Value: sum}, e)
			nodes.Remove(e)
			return
		} else if node.Kind == MinusKind {
			eA := e.Next()
			nodeA := eA.Value.(Node)
			eb := eA.Next()
			nodeB := eb.Value.(Node)
			sum := nodeB.Value - nodeA.Value

			nodes.Remove(eb)
			nodes.Remove(eA)
			nodes.InsertBefore(Node{Kind: NumberKind, Value: sum}, e)
			nodes.Remove(e)
			return
		}

	}
}
