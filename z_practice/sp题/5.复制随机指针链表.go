package main


type Node struct {
	Val int
	Next *Node
	Random *Node
}


func copyRandomList(head *Node) *Node {
	// 2次循环遍历，第一遍创建node链表，并且赋予next
	// 第二遍赋予random

	var slow *Node = &Node{}
	fast := head
	begin := slow

	sfMaper := make(map[*Node]*Node)

	for fast!=nil{
		newNode := &Node{fast.Val,nil, nil}
		slow.Next = newNode
		sfMaper[fast] = slow.Next
		slow = slow.Next
		fast = fast.Next
	}

	slow = begin.Next
	fast = head

	for fast!=nil{
		slow.Random = sfMaper[fast.Random]
		slow = slow.Next
		fast = fast.Next
	}
	return begin.Next
}


func main(){
//[[7,null],[13,0],[11,4],[10,2],[1,0]]
//输出：
//[[7,4],[13,2],[11,0],[10,null],[1,null]]
}