package main

import "fmt"

/*
92. Reverse Linked List II
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
 */


type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil{return nil}

	dummy := &ListNode{}
	dummy.Next = head
	head = dummy
	pos := 1

	for pos<m{
		head = head.Next
		pos+=1
	}
	// 此时 1（head） 2 3 4 5
	pre := head.Next
	cur := head.Next.Next

	// 此时 1 2（pre） 3（cur） 4 5
	for pos<n{
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
		pos+=1
	}
	// 此时 1 2 3 4（pre）5（cur）
	head.Next.Next = cur
	head.Next = pre
	return dummy.Next
}

func main(){
	e := ListNode{5, nil}
	d := ListNode{4, &e}
	c := ListNode{3, &d}
	b := ListNode{2,&c}
	a := ListNode{1,&b}
	fmt.Println(reverseBetween(&a, 2,4))
}