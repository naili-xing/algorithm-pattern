package main

import "fmt"

/*
	234. Palindrome Linked List
	Given a singly linked list, determine if it is a palindrome.

	Example 1:

	Input: 1->2
	Output: false
	Example 2:

	Input: 1->2->2->1
	Output: true
	Follow up:
	Could you do it in O(n) time and O(1) space?
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil{return true}
	mid := findMid3(head)
	mid = reverse3(mid)

	for mid!=nil && head!=nil{
		if mid.Val != head.Val{
			return false
		}
		mid = mid.Next
		head = head.Next
	}
	return true

}

func findMid3(head *ListNode) *ListNode {

	leng := 0
	begin := head
	for begin!=nil{
		leng += 1
		begin = begin.Next
	}

	flag :=1
	for head.Next !=nil{
		if flag == leng/2{
			tem := head.Next
			head.Next = nil
			return tem
		}

		flag+=1
		head = head.Next
	}
	return head
}


func reverse3(head *ListNode) *ListNode {

	var slow *ListNode = nil
	fast := head

	for fast!=nil{
		tmp := fast.Next
		fast.Next = slow
		slow = fast
		fast  = tmp
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
	fmt.Println(isPalindrome(&a))
}