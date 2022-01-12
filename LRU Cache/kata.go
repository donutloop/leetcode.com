package LRU_Cache

import (
	"time"
)

type element struct {
	value int
	time  int64
}

type LRUCache struct {
	elements map[int]*element
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		elements: make(map[int]*element),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {

	element, ok := this.elements[key]
	if !ok {
		return -1
	}

	element.time = time.Now().UnixNano()

	this.elements[key] = element

	return element.value
}

func (this *LRUCache) Put(key int, value int) {

	e, ok := this.elements[key]
	if !ok {
		if len(this.elements) == this.capacity {

			max := int64(-1)
			foundKey := -1
			for key, value := range this.elements {
				if max == -1 || max > value.time {
					foundKey = key
					max = value.time
				}
			}
			delete(this.elements, foundKey)
		}

		this.elements[key] = &element{value: value, time: time.Now().UnixNano()}
		return
	}

	e.value = value
	e.time = time.Now().UnixNano()

	this.elements[key] = e
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
