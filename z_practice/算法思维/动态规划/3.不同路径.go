package main

import "fmt"

/*
	62. Unique Paths
	A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

	The robot can only move either down or right at any point in time. The robot is trying to
	reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

	How many possible unique paths are there?
 */


func uniquePaths(m int, n int) int {

	it := make([][]int, m)

	for i:=0;i<m;i++{
		it[i] = make([]int, n)
		for j:=0;j<n;j++{
			if i==0 || j==0{
				it[i][j] =1
			}else{
				it[i][j]= it[i-1][j] + it[i][j-1]
			}
		}
	}
	return it[len(it)-1][len(it[0])-1]
}

