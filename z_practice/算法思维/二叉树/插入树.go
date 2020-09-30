package main

import "fmt"

/*
701. Insert into a Binary Search Tree
Given the root node of a binary search tree (BST) and a value to be inserted into the tree,
insert the value into the BST. Return the root node of the BST after the insertion.
It is guaranteed that the new value does not exist in the original BST.

Note that there may exist multiple valid ways for the insertion,
as long as the tree remains a BST after insertion. You can return any of them.

For example,

Given the tree:
        4
       / \
      2   7
     / \
    1   3
And the value to insert: 5
You can return this binary search tree:

         4
       /   \
      2     7
     / \   /
    1   3 5
This tree is also valid:

         5
       /   \
      2     7
     / \
    1   3
         \
          4
 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := root
	if root==nil{return &TreeNode{val, nil,nil}}
	for {
		if val > node.Val{
			if node.Right==nil{
				node.Right=&TreeNode{val, nil, nil}
				return root
			}else{
				node = node.Right
			}
		}
		if val < node.Val{
			if node.Left==nil{
				node.Left=&TreeNode{val, nil, nil}
				return root
			}else{
				node= node.Left
			}
		}
	}
}

func main(){
	d1 := TreeNode{Val: 50, Left: nil, Right: nil}
	d2 := TreeNode{Val: 70, Left: nil, Right: nil}
	c1 := TreeNode{Val: 10, Left: nil, Right: nil}
	c2 := TreeNode{Val: 30, Left: nil, Right: nil}
	b2 := TreeNode{Val: 60, Left: &d1, Right: &d2}
	b1 := TreeNode{Val: 20, Left: &c1, Right: &c2}
	a := TreeNode{Val: 40, Left: &b1, Right: &b2}
	fmt.Println(insertIntoBST(&a, 25))
}