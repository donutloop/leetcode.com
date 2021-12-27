package kata

type MapSum struct {
	values map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{values: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.values[key] = val
}

func (this *MapSum) Sum(prefix string) (sum int) {
	for key, value := range this.values {
		if len(key) >= len(prefix) && key[:len(prefix)] == prefix {
			sum = sum + value
		}
	}
	return
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
