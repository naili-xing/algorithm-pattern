package main


/*
450. Delete Node in a BST
Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.
Follow up: Can you solve it with time complexity O(height of tree)?

Example 1:

Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.


 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {

	//总共3种情况
	//1。 node只有左树，保留左
	//1。 node只有右树，保留右
	//1。 node有左树&右树，保留右和左
	if root ==nil{
		return nil
	}
	if root.Val < key{
		root.Right = deleteNode(root.Right,key)
	}else if root.Val > key{
		root.Left = deleteNode(root.Left,key)
	}else{
		if root.Right == nil{
			root = root.Left
		}else if root.Left == nil{
			root = root.Right
		}else{
			node := root.Right

			for node.Left!=nil{
				node = node.Left
			}
			node.Left = root.Left
			root = root.Right
		}
	}
	return root
}


func main(){

}