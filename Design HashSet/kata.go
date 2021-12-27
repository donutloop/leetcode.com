package kata

import (
	"container/list"
)

type MyHashSet struct {
	buckets []*list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	buckets := make([]*list.List, 500)
	for i := range buckets {
		buckets[i] = list.New()
	}
	return MyHashSet{
		buckets: buckets,
	}
}

func (this *MyHashSet) Add(key int) {
	idx := key % 500
	for e := this.buckets[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return
		}
	}
	this.buckets[idx].PushFront(key)
}

func (this *MyHashSet) Remove(key int) {
	idx := key % 500
	for e := this.buckets[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.buckets[idx].Remove(e)
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	idx := key % 500
	for e := this.buckets[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
