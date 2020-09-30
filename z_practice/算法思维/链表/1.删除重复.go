package main

import "fmt"

/*
83. 删除排序链表中的重复元素
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
 */


type ListNode struct {
  Val int
  Next *ListNode
}

// 指针
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{return nil}

	slow := head
	fast := head.Next

	for fast!=nil{
		if fast.Val ==slow.Val{
			fast = fast.Next
		}else{
			slow.Next = fast
			slow = fast
			fast = fast.Next
		}
	}
	slow.Next = nil
	return head
}

// 单指针

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil{return nil}
	node := head
	for node!=nil {
		for node.Next!=nil && node.Val ==node.Next.Val{
			node.Next = node.Next.Next
		}
		node = node.Next
	}
	return head
}


func main(){
	e := ListNode{3, nil}
	d := ListNode{3,&e}
	c := ListNode{2, &d}
	b := ListNode{1, &c}
	a := ListNode{1,&b}
	fmt.Println(deleteDuplicates2(&a))

}





