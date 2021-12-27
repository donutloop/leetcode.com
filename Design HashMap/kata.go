package kata

import "container/list"

type HashMapValue struct {
	key   int
	value int
}

type MyHashMap struct {
	buckets []*list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	buckets := make([]*list.List, 500)
	for i := range buckets {
		buckets[i] = list.New()
	}
	return MyHashMap{
		buckets: buckets,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	idx := key % 500
	for e := this.buckets[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(*HashMapValue).key == key {
			e.Value.(*HashMapValue).value = value
			return
		}
	}
	this.buckets[idx].PushFront(&HashMapValue{key: key, value: value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	idx := key % 500
	for e := this.buckets[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(*HashMapValue).key == key {
			return e.Value.(*HashMapValue).value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	idx := key % 500
	for e := this.buckets[idx].Front(); e != nil; e = e.Next() {
		if e.Value.(*HashMapValue).key == key {
			this.buckets[idx].Remove(e)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
