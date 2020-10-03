package main

import "strconv"


type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	s1:=""
	s2:=""

	for l1!=nil{
		s1 = strconv.Itoa(l1.Val) + s1
		l1 = l1.Next
	}
	for l2!=nil{
		s2 = strconv.Itoa(l2.Val) + s2
		l2 = l2.Next
	}

	res := add2string(s1,s2)

	head := &ListNode{}
	cur := head
	for i:=len(res)-1;i>=0;i--{
		cur.Next = &ListNode{int(res[i])-int('0'),nil}
		cur = cur.Next
	}
	return head.Next
}



func add2string(s1,s2 string)string{
	res := ""
	add := 0
	for i,j:=len(s1)-1,len(s2)-1 ;i>=0||j>=0|| add>0; i,j = i-1,j-1{
		var ts1 = 0
		var ts2 = 0
		if i>=0{
			ts1 += int(s1[i] - '0')
		}
		if j>=0{
			ts2 += int(s2[j] - '0')
		}
		tr := ts1 + ts2 +add
		res = strconv.Itoa(tr%10) +res
		add = tr/10
	}
	return res
}

