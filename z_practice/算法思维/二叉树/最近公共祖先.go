package main

import "fmt"


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var visited = make(map[int]bool)
	var parents  = make(map[int]*TreeNode)
	ComHelper(root, parents)
	for p!=nil{
		visited[p.Val]=true
		p=parents[p.Val]
	}
	for q!= nil{
		if visited[q.Val]{
			return q
		}
		q=parents[q.Val]
		}
		return nil
}

func ComHelper(root *TreeNode, parents map[int]*TreeNode){

	if root !=nil{
		if root.Left!=nil{
			parents[root.Left.Val] = root
			ComHelper(root.Left, parents)
		}
		if root.Right!=nil{
			ComHelper(root.Right, parents)
			parents[root.Right.Val] = root
		}
	}
}

func main(){
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Left.Right.Right).Val)

}