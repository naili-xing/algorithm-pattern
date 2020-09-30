package main

import "fmt"

/*
	148. 排序链表
	在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

	示例 1:

	输入: 4->2->1->3
	输出: 1->2->3->4
	示例 2:

	输入: -1->5->3->4->0
	输出: -1->0->3->4->5
*/

type ListNode struct {
	Val int
	Next *ListNode
}

// nlogn 典型的归并发

func sortList(head *ListNode) *ListNode {
	if head==nil{return head}
	res := mergeSort(head)
	return res
}

func mergeSort(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	//if head.Next.Next == nil{
	//	if head.Next.Val<head.Val{
	//		head.Next.Next = head
	//		a := head.Next
	//		head.Next = nil
	//		return a
	//	}
	//	return head
	//}
	mid := findMid(head)
	left:=mergeSort(head)
	right:=mergeSort(mid)
	res := combine(left, right)
	return res
}

func combine(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	begin := dummy

	for left!=nil && right!=nil{
		if left.Val>right.Val{
			dummy.Next = right
			right = right.Next
		}else{
			dummy.Next = left
			left = left.Next
		}
		dummy = dummy.Next
	}
	if left!=nil{
		dummy.Next = left
	}
	if right!=nil{
		dummy.Next = right
	}
	return begin.Next
}


func findMid(head *ListNode) *ListNode {
	slow := head
	fast := head.Next

	for fast!=nil && fast.Next!=nil{
		if fast.Next.Next == nil && slow.Next!=nil{
			slow.Next = nil
			return fast
		}
		fast = fast.Next
		slow = slow.Next
	}
	return head
}


func main(){

	a1 := ListNode{3, nil}
	c := ListNode{1, &a1}
	b := ListNode{2,&c}
	a := ListNode{4,&b}
	fmt.Println(sortList(&a))
}
