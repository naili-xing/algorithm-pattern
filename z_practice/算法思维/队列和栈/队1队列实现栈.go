package main

import "database/sql/driver"

/*
	232. 用栈实现队列
	使用栈实现队列的下列操作：

	push(x) -- 将一个元素放入队列的尾部。
	pop() -- 从队列首部移除元素。
	peek() -- 返回队列首部的元素。
	empty() -- 返回队列是否为空。


	示例:

	MyQueue queue = new MyQueue();

	queue.push(1);
	queue.push(2);
	queue.peek();  // 返回 1
	queue.pop();   // 返回 1
	queue.empty(); // 返回 false
 */

type MyQueue struct {
	queue []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{queue: make([]int, 0)}

}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.queue = append(this.queue, x)

}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	ele:=this.queue[0]
	this.queue = this.queue[1:]
	return ele

}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	ele:=this.queue[0]
	return ele
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.queue)==0{
		return true
	}else{
		return false
	}
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */


