package main

import "fmt"

/*
103. Binary Tree Zigzag Level Order Traversal
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func zigzagLevelOrder(root *TreeNode) [][]int {

	var res [][]int
	var level []*TreeNode
	level= append(level, root)

	num :=0
	for len(level)>0{
		l := len(level)

		list := make([]int, 0)
		for i:=0;i<l;i++{
			node := level[0]
			level = level[1:]
			list = append(list, node.Val)
			if node.Left!=nil{
				level = append(level, node.Left)
			}
			if node.Right!=nil{
				level = append(level, node.Right)
			}
		}
		// 翻转
		if num%2==1{
			for i:=0;i<len(list)/2;i++{
				list[i], list[len(list)-i-1] = list[len(list)-i-1], list[i]
			}
		}
		res = append(res, list)
		num+=1
	}

	return res
}


func main(){

	c1 := TreeNode{Val: 15, Left: nil, Right: nil}
	c2 := TreeNode{Val: 7, Left: nil, Right: nil}
	b1 := TreeNode{Val: 9, Left: nil, Right: nil}
	b2 := TreeNode{Val: 20, Left: &c1, Right: &c2}
	a := TreeNode{Val: 3, Left: &b1, Right: &b2}

	fmt.Println(zigzagLevelOrder(&a))
}
