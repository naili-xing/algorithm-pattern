package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
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

func levelOrder(root *TreeNode) [][]int {

	if root==nil{return [][]int{}}

	var res [][]int
	var list []int
	var level []*TreeNode

	level = append(level, root)

	for len(level)>0{

		l := len(level)
		for i:=0; i<l;i++{
			node := level[0]
			level = level[0+1:]
			list = append(list, node.Val)

			if node.Left!=nil{
				level = append(level, node.Left)
			}
			if node.Right!=nil{
				level = append(level, node.Right)
			}
		}
		res = append(res, list)
		list = []int{}
	}
	return res
}

func main(){
	c1 := TreeNode{Val: 15, Left: nil, Right: nil}
	c2 := TreeNode{Val: 7, Left: nil, Right: nil}
	b1 := TreeNode{Val: 9, Left: nil, Right: nil}
	b2 := TreeNode{Val: 20, Left: &c1, Right: &c2}
	a := TreeNode{Val: 3, Left: &b1, Right: &b2}
	fmt.Println(levelOrder(&a))
}