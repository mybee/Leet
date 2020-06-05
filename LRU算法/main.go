package main

import (
	"container/list"
	"fmt"
)

func main() {

	cache := Constructor(2);

	cache.Put(1, 1);
	cache.Put(2, 2);
	fmt.Println(cache.Get(1))      // 返回  1
	cache.Put(3, 3);    // 该操作会使得关键字 2 作废
	fmt.Println(cache.Get(2))     // 返回 -1 (未找到)
	cache.Put(4, 4);    // 该操作会使得关键字 1 作废
	fmt.Println(cache.Get(1))     // 返回 -1 (未找到)
	fmt.Println(cache.Get(3))      // 返回  3
	fmt.Println(cache.Get(4))       // 返回  4


}

// lru 算法核心:
//   1. hash 实现 快速查找
//   2. 双向链表(list) 实现 有序排列

type LRUCache struct {
	Cap int
	L *list.List
	M map[int](*list.Element)
}

type LRUNode struct {
	Key int
	Value int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap : capacity,
		L: list.New(),
		M: make(map[int]*list.Element),
	}
}


func (this *LRUCache) Get(key int) int {
	if e,ok := this.M[key]; ok {
		this.L.MoveToFront(e)
		return (e.Value).(LRUNode).Value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if e, ok := this.M[key]; ok {
		e.Value = LRUNode{Key:key, Value: value}
		this.L.MoveToFront(e)
	}else{
		if this.L.Len() == this.Cap {
			e := this.L.Back()
			delete(this.M, (e.Value).(LRUNode).Key)
			this.L.Remove(e)
		}
		v := LRUNode{
			Key: key,
			Value: value,
		}
		e := this.L.PushFront(v)
		this.M[key] = e
	}
}