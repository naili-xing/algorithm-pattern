package main
import "fmt"


/*
	21. 合并两个有序链表
	将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

	示例：

	输入：1->2->4, 1->3->4
	输出：1->1->2->3->4->4
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var cur *ListNode =&ListNode{}

	dummy := cur
	for l1!=nil && l2!= nil{
		if l1.Val>l2.Val{
			cur.Next = l2
			cur = l2
			l2 = l2.Next
		}else{
			cur.Next = l1
			cur = l1
			l1 = l1.Next
		}
	}
	if l1!=nil{
		cur.Next = l1
	}
	if l2!=nil{
		cur.Next = l2
	}
	return dummy.Next
}


func main(){
	//g := ListNode{7,nil}
	//f := ListNode{6,&d2}
	c1 := ListNode{5, nil}
	b1 := ListNode{4, &c1}
	a1 := ListNode{3, &b1}

	c := ListNode{3, nil}
	b := ListNode{2,&c}
	a := ListNode{1,&b}

	fmt.Println(mergeTwoLists(&a, &a1))
}