package main

import "fmt"

/*
	200. 岛屿数量
	给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

	岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

	此外，你可以假设该网格的四条边均被水包围。

	示例 1:

	输入:
	[
	['1','1','1','1','0'],
	['1','1','0','1','0'],
	['1','1','0','0','0'],
	['0','0','0','0','0']
	]
	输出: 1
	示例 2:

	输入:
	[
	['1','1','0','0','0'],
	['1','1','0','0','0'],
	['0','0','1','0','0'],
	['0','0','0','1','1']
	]
输出: 3
 */

func numIslands(grid [][]byte) int {

	res := 0
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if grid[i][j]=='1'{
				Search(&grid, i, j )
				fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++", i, j)
				res +=1
			}

		}
	}
	return res
}

func Search(grid *[][]byte, i, j int) {
	fmt.Println("========================================================================", i, j)
	for _, row := range *grid {
		for _, val := range row {
			// 3位对齐，打印结果
			fmt.Printf(string(val))
		}
		fmt.Println()
	}
	//fmt.Println("searching ", i, j )
	if i>=len(*grid) || i<0 || j>=len((*grid)[0]) || j<0{
		return
	}
	if (*grid)[i][j] == '1'{
		(*grid)[i][j] = '0'
		Search(grid, i, j-1)
		Search(grid, i, j+1)
		Search(grid, i+1, j)
		Search(grid, i-1, j)
	}
}


func  main(){
	a := [][]byte{{'1','0','1','1','1'}, {'1','0','1','0','1'}, {'1','1','1','0','1'}}
	fmt.Println(numIslands(a))
}
