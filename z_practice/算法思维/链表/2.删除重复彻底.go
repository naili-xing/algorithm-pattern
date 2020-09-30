package main

import "fmt"

/*
82. 删除排序链表中的重复元素 II
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3

 */

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates3(head *ListNode) *ListNode {

	if head==nil{return head}
	if head.Next==nil{return head}

	res  := make(map[int]bool)
	node := head
	for node!=nil&& node.Next!=nil{
		if node.Val==node.Next.Val{
			res[node.Val]=true
		}
		node = node.Next
	}

	// 创建了一个node ，通过fast， 改变node的next指向
	begin:=&ListNode{0,head}
	fast := begin

	for fast.Next!=nil{
		if _, ok:=res[fast.Next.Val];ok{
			fast.Next = fast.Next.Next
		}else{
			fast = fast.Next
		}
	}
	return begin.Next
}

func main(){
	//d2 := ListNode{5,nil}
	//d := ListNode{4,&d2}
	//c := ListNode{4, &d}
	//b2 := ListNode{3, &c}
	//b := ListNode{3, &b2}
	//a2 := ListNode{2,&b}
	//a := ListNode{1,&a2}
	b := ListNode{1, nil}
	a := ListNode{1,&b}
	fmt.Println(deleteDuplicates3(&a))
}