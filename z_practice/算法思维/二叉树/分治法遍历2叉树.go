package main

import "fmt"

/*
分治法应用 先分别处理局部，再合并结果

适用场景
	快速排序
	归并排序
	二叉树相关问题

分治法模板
    递归返回条件
	分段处理
	合并结果

func traversal(root *TreeNode) ResultType  {
    // nil or leaf
    if root == nil {
        // do something and return
    }

    // Divide
    ResultType left = traversal(root.Left)
    ResultType right = traversal(root.Right)

    // Conquer
    ResultType result = Merge from left and right
    return result
}
 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func preorderTraversal3(root *TreeNode) []int {
	res := divideAndConquer(root)
	return res
}


func divideAndConquer(root *TreeNode) []int{
	res := []int{}
	if root==nil{
		return res
	}

	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)
	return res
}

func main(){
	c1 := TreeNode{Val: 15, Left: nil, Right: nil}
	c2 := TreeNode{Val: 7, Left: nil, Right: nil}
	b1 := TreeNode{Val: 9, Left: nil, Right: nil}
	b2 := TreeNode{Val: 20, Left: &c1, Right: &c2}
	a := TreeNode{Val: 3, Left: &b1, Right: &b2}
	fmt.Println(preorderTraversal3(&a))
}
