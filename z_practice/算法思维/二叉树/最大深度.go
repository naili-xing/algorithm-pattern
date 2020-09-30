package main

import "fmt"

/*
最深的深度
 */



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root==nil{return 0}
	leng := 0
	var list []*TreeNode
	list = append(list, root)

	for len(list)>0{
		l := len(list)
		for i:=0; i<l; i++{
			node := list[0]
			list = list[0+1:]
			if node.Left!=nil{
				list = append(list, node.Left)
			}
			if node.Right!=nil{
				list = append(list, node.Right)
			}
		}
		leng+=1
	}
	return leng
}

func main(){
	c1 := TreeNode{Val: 15, Left: nil, Right: nil}
	c2 := TreeNode{Val: 7, Left: nil, Right: nil}
	b1 := TreeNode{Val: 9, Left: nil, Right: nil}
	b2 := TreeNode{Val: 20, Left: &c1, Right: &c2}
	a := TreeNode{Val: 3, Left: &b1, Right: &b2}
	fmt.Println(maxDepth(&a))

}

