package main

import "fmt"

/*
	86. 分隔链表
	给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

	你应当保留两个分区中每个节点的初始相对位置。
	示例:
	输入: head = 1->4->3->2->5->2, x = 3
	输出: 1->2->2->4->3->5
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	var left = &ListNode{Val: 0}
	var right = &ListNode{Val: 0}
	dl := left
	dr := right
	for head != nil{
		if head.Val<x{
			left.Next = head
			left = head
		}else{
			right.Next = head
			right = head
		}
		head =  head.Next
	}
	// 必须设置nil，否则引起循环
	right.Next = nil
	left.Next = dr.Next
	return dl.Next
}

func partition2(head *ListNode, x int) *ListNode {
	// 思路：将大于x的节点，放到另外一个链表，最后连接这两个链表
	// check
	if head == nil {
		return head
	}
	headDummy := &ListNode{Val: 0}
	tailDummy := &ListNode{Val: 0}
	tail := tailDummy
	head = headDummy

	headDummy.Next = head

	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			// 移除<x节点
			t := head.Next
			head.Next = head.Next.Next
			// 放到另外一个链表
			tail.Next = t
			tail = tail.Next
		}
	}
	// 拼接两个链表
	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}



func main(){
	//g := ListNode{7,nil}
	//f := ListNode{6,&d2}
	c1 := ListNode{2, nil}
	b1 := ListNode{5, &c1}
	a1 := ListNode{2, &b1}
	c := ListNode{3, &a1}
	b := ListNode{4,&c}
	a := ListNode{1,&b}
	fmt.Println(partition(&a,3))
}
