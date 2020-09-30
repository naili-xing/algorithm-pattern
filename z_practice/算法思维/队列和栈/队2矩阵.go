package main

/*
	542. 01 矩阵
	给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。

	两个相邻元素间的距离为 1 。

	示例 1:
	输入:

	0 0 0
	0 1 0
	0 0 0
	输出:

	0 0 0
	0 1 0
	0 0 0
	示例 2:
	输入:

	0 0 0
	0 1 0
	1 1 1
	输出:

	0 0 0
	0 1 0
	1 2 1
	注意:

	给定矩阵的元素个数不超过 10000。
	给定矩阵中至少有一个元素是 0。
	矩阵中的元素只在四个方向上相邻: 上、下、左、右。
 */

func updateMatrix(matrix [][]int) [][]int {

	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[i]);j++{
			if matrix[i][j] == 1{
				num :=0
				for m:=i+1;m<len(matrix);m++{
					if matrix[m][j] == 0{
						break
					}else{
						num += 1
					}
				}
				num2 := 0
				for n:=i+1;n<len(matrix[i]);n++{
					if matrix[i][n] == 0{
						break
					}else{
						num2 += 1
					}
				}
				if num2>num &&num>1 &&num2>1{
					matrix[i][j] = num
				}
				if num2<num &&num>1 &&num2>1{
					matrix[i][j] = num2
				}

			}
		}
	}
	return matrix
}



func main(){

}