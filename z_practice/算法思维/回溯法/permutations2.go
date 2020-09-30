package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
 */

func permuteUnique1(nums []int) [][]int {
	var res [][]int
	var list []int
	visited := make(map[int]bool)
	sort.Ints(nums)
	perhelper2(nums, visited, list, &res)
	return res
}

func perhelper2(nums []int, visited map[int]bool, list []int, res *[][]int){

	if len(nums)==len(list){
		ans := make([]int, len(list))
		copy(ans, list)
		*res = append(*res, ans)
	}

	for i:=0; i<len(nums); i++{
		// !visited[i-1] 的目的是确保不同枝中的重复路径不访问，而同一枝下的路径可以重复
		// 因为在第二次检查nums中的跟节点时， 第一个根节点的访问状态会推倒没访问过的状态
		if i>0 && nums[i] == nums[i-1] && !visited[i-1]{
			continue
		}

		if visited[i]{
			continue
		}
		visited[i] = true
		list = append(list, nums[i])
		perhelper2(nums, visited, list, res)
		list = list[0: len(list)-1]
		visited[i] = false
	}
}

func main(){
	fmt.Println(permuteUnique1([]int{2,2,1,1}))

}