package main

import "fmt"

/*
124. 二叉树中的最大路径和
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1：

输入：[1,2,3]

       1
      / \
     2   3

输出：6
示例 2：

输入：[-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出：42

 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func maxPathSum(root *TreeNode) int {




}

func main(){
	c1 := TreeNode{Val: 15, Left: nil, Right: nil}
	c2 := TreeNode{Val: 7, Left: nil, Right: nil}
	b1 := TreeNode{Val: 9, Left: nil, Right: nil}
	b2 := TreeNode{Val: 20, Left: &c1, Right: &c2}
	a := TreeNode{Val: 3, Left: &b1, Right: &b2}
	fmt.Println(maxPathSum(&a))
}