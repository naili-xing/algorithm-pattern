package main

import (
	"fmt"
)

/*
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/data-structure-binary-tree/xebrb2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归
func postorderTraversal(root *TreeNode) []int {
	var res []int
	posHelper(root, &res)
	return res
}

func posHelper(root *TreeNode, res *[]int){
	if root == nil {return }

	posHelper(root.Left, res)
	posHelper(root.Right, res)
	*res = append(*res, root.Val)
}

//迭代
func postorderTraversal2(root *TreeNode) []int{

	 var res []int
	 var stack []*TreeNode
	 var viewed *TreeNode

	 for root != nil || len(stack) > 0 {

	 	for root != nil {
	 		// 进栈顺序不变
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right==nil ||  node.Right == viewed {
			stack = stack[:len(stack)-1] // pop
			res = append(res, node.Val)
			// 标记当前这个节点已经弹出过
			viewed = node
		}else{root = node.Right}
	 }
	 return res
}


func main(){

	c := TreeNode{Val: 3, Left: nil, Right: nil}
	b := TreeNode{Val: 2, Left: &c, Right: nil}
	a := TreeNode{Val: 1, Left: nil, Right: &b}
	fmt.Println(postorderTraversal2(&a))

}
