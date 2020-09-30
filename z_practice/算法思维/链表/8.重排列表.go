package main

import "fmt"

/*
	143. Reorder List
	Given a singly linked list L: L0→L1→…→Ln-1→Ln,
	reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

	You may not modify the values in the list's nodes, only nodes itself may be changed.

	Example 1:

	Given 1->2->3->4, reorder it to 1->4->2->3.
	Example 2:

	Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

 */
type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode)  {

	if head == nil{return}
	if head.Next==nil{return}

	// mid 是后变的头
	var mid *ListNode

	mid = findMid2(head)
	mid = reverseIt(mid)

	slow := head
	fast := mid

	flag := 0
	for fast!=nil{
		th:=head.Next
		tm:=mid.Next

		slow.Next = fast
		slow = fast
		if flag%2==0{
			fast = th
			head = th
		}else{
			fast =tm
			mid =tm
		}
		flag += 1
	}

	if mid!=nil{
		slow.Next = mid.Next
	}
	if head!=nil{
		slow.Next = head.Next
	}
}

func findMid2(head *ListNode)*ListNode{

	begin := head
	leng :=1
	for begin.Next!=nil{
		begin = begin.Next
		leng += 1
	}
	pos := 1
	for head.Next != nil{
		if pos == leng/2{
			tmp := head.Next
			head.Next = nil
			return tmp
		}
		head = head.Next
		pos += 1
	}
	return head
}

func reverseIt(head *ListNode)*ListNode{

	var slow *ListNode = nil
	fast := head

	for fast !=nil{
		tm := fast.Next
		fast.Next = slow
		slow = fast
		fast = tm
	}
	return slow
}

func main(){
	//e := ListNode{5, nil}
	d := ListNode{4, nil}
	//c := ListNode{3, &d}
	//b := ListNode{2,&c}
	//a := ListNode{1,&b}

	reorderList(&d)
	fmt.Println(d)
}
