package kata

type item struct {
	value     string
	timestamp int
}

type TimeMap struct {
	store map[string][]item
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]item)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	v, ok := this.store[key]
	if !ok {
		list := make([]item, 1)
		list[0].value = value
		list[0].timestamp = timestamp
		this.store[key] = list
		return
	}
	v = append(v, item{value: value, timestamp: timestamp})
	this.store[key] = v
}

func (this *TimeMap) Get(key string, timestamp int) string {
	list, ok := this.store[key]
	if !ok {
		return ""
	}
	max := -1
	index := 0
	min := -1
	for i, item := range list {
		if item.timestamp == timestamp {
			return item.value
		}
		if max == -1 || max < item.timestamp && timestamp > item.timestamp {
			max = item.timestamp
			index = i
		}
		if min == -1 || min > item.timestamp {
			min = item.timestamp
		}
	}
	if min > timestamp {
		return ""
	}
	return list[index].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
