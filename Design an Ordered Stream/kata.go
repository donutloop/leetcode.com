package Design_an_Ordered_Stream

type OrderedStream struct {
	values       []*string
	currentIndex int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{values: make([]*string, n)}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	if this.values[idKey-1] == nil {
		this.values[idKey-1] = new(string)
		*this.values[idKey-1] = value
	} else {
		*this.values[idKey-1] = value
	}

	rightValues := make([]string, 0)
	for this.currentIndex < len(this.values) {
		if this.values[this.currentIndex] == nil {
			break
		}
		rightValues = append(rightValues, *this.values[this.currentIndex])
		this.currentIndex++
	}
	return rightValues
}
