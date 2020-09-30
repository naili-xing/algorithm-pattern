package main

import "fmt"

/*
206. Reverse Linked List
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var slow *ListNode = nil
	fast := head
	for fast!=nil{
		nn := fast.Next
		fast.Next = slow
		slow = fast
		fast = nn
	}
	return slow
}

func main(){
	//g := ListNode{7,nil}
	//f := ListNode{6,&d2}
	e := ListNode{5, nil}
	d := ListNode{4, &e}
	c := ListNode{3, &d}
	b := ListNode{2,&c}
	a := ListNode{1,&b}
	fmt.Println(reverseList(&a))
}