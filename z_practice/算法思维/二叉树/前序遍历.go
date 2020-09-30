package main
import "fmt"

/*
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// 递归法
func preorderTraversal(root *TreeNode) []int {

	var res []int
	preHelper(root, &res)
   	return res
}

func preHelper(root *TreeNode, res *[]int){
	if root == nil{return}
	*res = append(*res, root.Val)
	preHelper(root.Left, res)
	preHelper(root.Right, res)
}


// 非递归法
func preorderTraversal2(root *TreeNode) []int{

	var res []int
	var stack []*TreeNode

	// 检查右边的是否为空，如果为空，是否队列还有数据
	for root != nil || len(stack)!=0 {

		// 一直往左找， 直到找到空为止
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		root = node.Right
	}
	return res
}

func main(){

	c := TreeNode{Val: 3, Left: nil, Right: nil}
	b := TreeNode{Val: 2, Left: &c, Right: nil}
	a := TreeNode{Val: 1, Left: nil, Right: &b}
	fmt.Println(preorderTraversal2(&a))
}
