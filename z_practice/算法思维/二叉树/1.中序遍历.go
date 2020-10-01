package main

import "fmt"

/*
	给定一个二叉树，返回它的中序 遍历。

	示例:

	输入: [1,null,2,3]
	   1
		\
		 2
		/
	   3

	输出: [1,3,2]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xecaj6/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// 递归法
func inorderTraversal(root *TreeNode) []int {

	var res  []int
	inorderHelper(root, &res)
	return res
}


func inorderHelper(root *TreeNode, res *[]int){

	if root==nil{
		return
	}
	inorderHelper(root.Left, res)
	*res = append(*res, root.Val)
	inorderHelper(root.Right, res)
}


// 迭代法
func inorderTraversal2(root *TreeNode) []int {

	var stack []*TreeNode
	var res []int

	for root != nil || len(stack)>0{

		// 左子树遍历
		for root!=nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		res = append(res, node.Val)

		root  = node.Right
	}
	return res
}

//注意不能用root代替node变量， 为啥？？？？
func inorderTraversal3(root *TreeNode) []int {

	var res  []int
	var stack []*TreeNode

	for root!=nil || len(stack)>0{
		// 往左走，放入栈
		for root!=nil{
			stack = append(stack, root)
			root = root.Left
		}

		root =  stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func main(){
	c := &TreeNode{Val: 3, Left: nil, Right: nil}
	b := &TreeNode{Val: 2, Left: c, Right: nil}
	a := &TreeNode{Val: 1, Left: nil, Right: b}

	fmt.Println(inorderTraversal3(a))
}