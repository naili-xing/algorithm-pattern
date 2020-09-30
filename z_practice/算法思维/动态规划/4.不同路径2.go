package main

import "fmt"

/*
	63. 不同路径 II一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）
	机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
	现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
	网格中的障碍物和空位置分别用 1 和 0 来表示。
	说明：m 和 n 的值均不超过 100。
	示例 1:
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	f := make([][]int, m)

	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
		for j := 0; j < n; j++ {
			f[i][j] = 1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i==0 && j !=0 {
				if f[i][j-1] == 0 || obstacleGrid[i][j] == 1{
					f[i][j] = 0
				}else{
					f[i][j] = 1
				}
			}else if i!=0 && j ==0{
				if f[i-1][j] == 0 || obstacleGrid[i][j] == 1{
					f[i][j] = 0
				}else{
					f[i][j] = 1
				}
			}else if i==0 && j==0{
				f[i][j] = 1
			}else{
				if obstacleGrid[i][j] == 1{
					f[i][j] = 0
				}else{
					f[i][j] = f[i-1][j] + f[i][j-1]
				}
			}
		}
	}
	return f[len(f)-1][len(f[0])-1]
}

func main(){
	fmt.Println(uniquePathsWithObstacles([][]int{{0,0,0},{0,1,0},{0,0,0}}))
}