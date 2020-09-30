package main

import (
	"fmt"
)

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/balanced-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func isBalanced(root *TreeNode) bool {
	if root ==nil{
		return true
	}
	return abs(Hight(root.Left)-Hight(root.Right))<=1 && isBalanced(root.Right) && isBalanced(root.Left)
}

func Hight(root *TreeNode) int{
	if root==nil{return 0}
	return max(Hight(root.Left) ,Hight(root.Right))+1
}

func max(a int, b int)int{
	if a>b{return a} else{return b}
}

func abs(v int)int{
	if v>0{return v}else{return -v}
}


func main(){
	c1 := TreeNode{Val: 15, Left: nil, Right: nil}
	c2 := TreeNode{Val: 7, Left: nil, Right: nil}
	b1 := TreeNode{Val: 9, Left: nil, Right: nil}
	b2 := TreeNode{Val: 20, Left: &c1, Right: &c2}
	a := TreeNode{Val: 3, Left: &b1, Right: &b2}
	fmt.Println(isBalanced(&a))

}

