package main

import "fmt"

/*
	120. Triangle
	Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.
	For example, given the following triangle
	[
		 [2],
		[3,4],
	   [6,5,7],
	  [4,1,8,3]
	]
	The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
	Note:
	Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
 */

func minimumTotal(triangle [][]int) int {

	//自底向上动态规划
	var it = make([][]int, len(triangle))

	// 先分配动态规划辅助矩阵
	for i:=0;i<len(triangle);i++{
		it[i] = make([]int, len(triangle[i]))
		for j:=0;j<len(triangle[i]);j++{
			it[i][j] = triangle[i][j]
		}
	}
	// 更新矩阵, 自底往上,从倒数第二行开始
	for i := len(triangle)-2;i>=0;i--{
		for j:=0;j<len(triangle[i]);j++{
			it[i][j] = min(it[i+1][j], it[i+1][j+1]) + it[i][j]
		}
	}
	return it[0][0]
}
func min(x,y int)int{
	if x>y{
		return y
	}else{
		return x
	}
}

func main(){
	fmt.Println(minimumTotal([][]int{{2},{3,4},{6,5,7}, {4,1,8,3}}))
}