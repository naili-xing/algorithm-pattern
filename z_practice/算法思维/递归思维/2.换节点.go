package main

/*
	24. Swap Nodes in Pairs
	Given a linked list, swap every two adjacent nodes and return its head.
	You may not modify the values in the list's nodes, only nodes itself may be changed.
	Example:
	Given 1->2->3->4, you should return the list as 2->1->4->3.
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	return swapHelper(head)
}

func swapHelper(head *ListNode) *ListNode{

	if head==nil || head.Next==nil {return nil}

	nextHead := head.Next.Next
	res := head.Next
	head.Next.Next = head
	head.Next = swapHelper(nextHead)
	return res

}


func main(){
	e := ListNode{5, nil}
	c := ListNode{3, &e}
	b := ListNode{2,&c}
	a := ListNode{1,&b}
	swapPairs(&a)
}

