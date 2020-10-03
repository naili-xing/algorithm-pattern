package main



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 通过递归，先找跟，然后找找左和右，
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0],nil,nil}

	rootIndx := 0

	for i:=0;i<len(inorder);i++{
		if inorder[i] == preorder[0]{
			rootIndx = i
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:rootIndx])+1], inorder[:rootIndx])
	root.Right = buildTree(preorder[len(inorder[:rootIndx])+1:], inorder[rootIndx+1:])
	return root
}


func main(){

}