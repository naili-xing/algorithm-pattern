package main
import "fmt"


type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	b := head
	leng := 0
	for b!=nil{
		leng +=1
		b = b.Next
	}

	dummy := &ListNode{}
	dummy.Next = head
	begin := dummy
	cur := 0

	fmt.Println(leng-n-1)
	for dummy!=nil {
		if cur == leng-n{
			dummy.Next = dummy.Next.Next
		}
		cur +=1
		dummy = dummy.Next
	}
	return begin.Next
}

func main(){

	c := ListNode{3, nil}
	b := ListNode{2,&c}
	a := ListNode{1,&b}
	fmt.Println(removeNthFromEnd(&a,1))

}
