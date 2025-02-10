package main

import (
	"container/list"
)

type LRUCache struct {
	capacity int
    mp map[int]*list.Element
	ll *list.List
}

type entry struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		capacity: capacity,
		mp: make(map[int]*list.Element),
		ll: list.New(),
	}
}


func (this *LRUCache) Get(key int) int {
    if v, ok := this.mp[key]; ok {
		this.ll.MoveToFront(v)
		return v.Value.(entry).value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if v, ok := this.mp[key]; ok {
		v.Value = entry{key, value}
		this.ll.MoveToFront(v)
	} else {
		this.mp[key] = this.ll.PushFront(entry{key, value})
		if this.ll.Len() > this.capacity {
			v := this.ll.Back()
			delete(this.mp, v.Value.(entry).key)
			this.ll.Remove(v)
		}
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {

}