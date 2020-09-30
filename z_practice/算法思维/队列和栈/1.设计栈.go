package main

/*
	155. 最小栈
	设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

	push(x) —— 将元素 x 推入栈中。
	pop() —— 删除栈顶的元素。
	top() —— 获取栈顶元素。
	getMin() —— 检索栈中的最小元素。

	示例:

	输入：
	["MinStack","push","push","push","getMin","pop","top","getMin"]
	[[],[-2],[0],[-3],[],[],[],[]]

	输出：
	[null,null,null,null,-3,null,0,-2]

	解释：
	MinStack minStack = new MinStack();
	minStack.push(-2);
	minStack.push(0);
	minStack.push(-3);
	minStack.getMin();   --> 返回 -3.
	minStack.pop();
	minStack.top();      --> 返回 0.
	minStack.getMin();   --> 返回 -2.
 */


type MinStack struct {
	min []int
	stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return  MinStack{min: make([]int, 0),
					 stack: make([]int, 0)}
}


func (this *MinStack) Push(x int)  {
	min := this.GetMin()
	if min>x{
		this.min = append(this.min, x)
	}else{
		this.min = append(this.min, min)
	}
	this.stack = append(this.stack, x)

}
func (this *MinStack) Pop()  {
	if len(this.stack) == 0 {return}
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}


func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 1 << 31
	}
	min := this.min[len(this.min)-1]
	return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */


