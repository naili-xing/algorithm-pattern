package main

import (
	"fmt"
)

/*
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。
当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
你是否可以在 O(1) 时间复杂度内完成这两种操作？
 */

type LRUCache struct {
	mapper map[int]*DLinkedList
	cur int
	capacity int
	head, tail *DLinkedList
}


func Constructor(capacity int) LRUCache {
	head:= &DLinkedList{0,0,nil,nil}
	tail:= &DLinkedList{0,0,nil,nil}
	head.next = tail
	tail.next = head
	return LRUCache{make(map[int]*DLinkedList),0, capacity,head, tail}
}

func (this *LRUCache) Get(key int) int {
	if v, ok:= this.mapper[key];ok{
		this.moveToHead(v)
		return v.value
	}else{
		return -1
	}
}

func (this *LRUCache) Put(key int, value int)  {
	if _, ok:= this.mapper[key];!ok{
		node := &DLinkedList{key, value, nil, nil}
		if this.cur >= this.capacity{
			v := this.removeTail()
			delete(this.mapper, v)
			this.cur--
		}
		this.addToHead(node)
		this.mapper[key] = node
		this.cur++
	}else{
		this.mapper[key].value = value
		this.moveToHead(this.mapper[key])
	}
}

type DLinkedList struct{
	key, value int
	pre, next  *DLinkedList
}


func (this *LRUCache) addToHead(node *DLinkedList){
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *DLinkedList){
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *DLinkedList){
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) removeTail() int{
	node := this.tail.pre
	this.removeNode(node)
	return node.key
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main(){
	cache := Constructor(1)

	cache.Put(2,1)
	a :=cache.Get(2)       // 返回  1
	cache.Put(3, 2)    // 该操作会使得关键字 2 作废
	b :=cache.Get(2)       // 返回 -1 (未找到)
	e :=cache.Get(3)       // 返回  4
	fmt.Println(a,b,e)
}
//["LRUCache","put","get","put","get","get"]
//[[1],[2,1],[2],[3,2],[2],[3]]