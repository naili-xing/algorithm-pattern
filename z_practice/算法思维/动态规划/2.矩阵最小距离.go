package main

import "fmt"

/*
64. 最小路径和
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

 */


func minPathSum(grid [][]int) int {

	it := make([][]int, len(grid))

	for i:=0;i<len(grid);i++{
		it[i]= make([]int, len(grid[i]))
		for j:=0;j<len(grid[i]);j++{
			it[i][j] = grid[i][j]
		}
	}
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if i==0 && j==0 {
				it[i][j] = it[i][j]
			}else if i ==0 && j!=0{
				it[i][j] = it[i][j-1]+it[i][j]
			}else if j==0 && i!=0{
				it[i][j] = it[i-1][j]+it[i][j]
			}else{
				it[i][j] = min(it[i-1][j], it[i][j-1]) + it[i][j]
			}
		}
	}
	return it[len(it)-1][len(it[0])-1]
}


func min(x,y int)int{
	if x>y{
		return y
	}else{
		return x
	}
}

func main(){
	fmt.Println(minPathSum([][]int{{1,3,1},{1,5,1},{4,2,1}}))
}
