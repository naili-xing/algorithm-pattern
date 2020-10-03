package main

import "fmt"

/*
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
例如输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {

	res := root
	var level []*TreeNode
	level = append(level, root)

	for len(level) >0{
		l := len(level)
		for i:=0;i<l;i++{
			node := level[0]
			level = level[1:]
			//交换左右树
			tmp := node.Left
			node.Left = node.Right
			node.Right = tmp
			if node.Left!=nil{
				level= append(level, node.Left)
			}
			if node.Right!=nil{
				level= append(level, node.Right)
			}
		}
	}
	return res
}


func main(){
	c3 := &TreeNode{Val: 9, Left: nil, Right: nil}
	c2 := &TreeNode{Val: 6, Left: nil, Right: nil}
	b1 := &TreeNode{Val: 3, Left: nil, Right: nil}
	a1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	c := &TreeNode{Val: 7, Left: c2, Right: c3}
	b := &TreeNode{Val: 2, Left: a1, Right: b1}
	a := &TreeNode{Val: 4, Left: b, Right: c}
	fmt.Println(mirrorTree(a))
}
