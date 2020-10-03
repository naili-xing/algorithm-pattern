package main

import "fmt"

/*
Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
Example:
Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
/**

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root==nil{return []int{}}
	var level []*TreeNode
	level = append(level, root)
	var res []int

	for len(level)>0{
		var list []int
		l := len(level)
		for i:=0;i<l;i++{
			root = level[i]
			list = append(list, root.Val)
			if root.Left!=nil{
				level = append(level, root.Left)

			}
			if root.Right!=nil{
				level = append(level, root.Right)
			}
		}
		level = level[l:]
		res = append(res, list[len(list)-1])
	}

	return res
}

func main(){
	c1 := TreeNode{Val: 5, Left: nil, Right: nil}
	c2 := TreeNode{Val: 4, Left: nil, Right: nil}
	b1 := TreeNode{Val: 3, Left: nil, Right: &c2}
	b2 := TreeNode{Val: 2, Left: &c1, Right: nil}
	a := TreeNode{Val: 1, Left: &b1, Right: &b2}
	fmt.Println(rightSideView(&a))

}