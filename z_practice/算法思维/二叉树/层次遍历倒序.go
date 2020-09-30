package main

import "fmt"

/*
	107. 二叉树的层次遍历 II
	给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

	例如：
	给定二叉树 [3,9,20,null,null,15,7],

		3
	   / \
	  9  20
		/  \
	   15   7
	返回其自底向上的层次遍历为：

	[
	  [15,7],
	  [9,20],
	  [3]
	]
 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {

	var level []*TreeNode
	var res [][]int

	level = append(level, root)

	for len(level)>0{
		list := make([]int, 0)
		l := len(level)
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
		res = append(res, list)
	}

	for i:=0;i<len(res)/2;i++{
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}


func main(){
	c1 := TreeNode{Val: 15, Left: nil, Right: nil}
	c2 := TreeNode{Val: 7, Left: nil, Right: nil}
	b1 := TreeNode{Val: 9, Left: nil, Right: nil}
	b2 := TreeNode{Val: 20, Left: &c1, Right: &c2}
	a := TreeNode{Val: 3, Left: &b1, Right: &b2}
	fmt.Println(levelOrderBottom(&a))
}