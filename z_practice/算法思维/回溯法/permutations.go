package main

import "fmt"

/*
	给定一个 没有重复 数字的序列，返回其所有可能的全排列。
	示例:
	输入: [1,2,3]
	输出:
	[
	  [1,2,3],
	  [1,3,2],
	  [2,1,3],
	  [2,3,1],
	  [3,1,2],
	  [3,2,1]
	]
 */

func permute(nums []int) [][]int {

	var res [][]int
	var list []int
	var visited []int

	perhelper(nums, visited, list, &res)
	return res

}

func isIn(list []int, ele int)bool{
	for _, v := range list{
		if v==ele{
			return true
		}
	}
	return false
}

func perhelper(nums []int, visited []int, list []int, res *[][]int){

	if len(list)==len(nums){
		ans := make([]int, len(list))
		copy(ans, list)
		*res = append(*res, ans)
	}
	for i:=0; i<len(nums); i++{
		if isIn(visited, i){
			continue
		}
		list = append(list, nums[i])
		visited = append(visited, i)
		perhelper(nums, visited, list, res)
		list = list[0:len(list)-1]
		visited = visited[0:len(visited)-1]
	}
}

func main(){
	fmt.Println(permute([]int{1,2,3}))
}