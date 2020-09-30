package main

import "fmt"

/*
	给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
	aa说明：解集不能包含重复的子集。

	输入: nums = [1,2,3]
	输出:
	[
	  [3],
	  [1],
	  [2],
	  [1,2,3],
	  [1,3],
	  [2,3],
	  [1,2],
	  []
	]
 */

func subsets(nums []int) [][]int {
	res := [][]int{}
	list := []int{}

	helper(nums, 0, list, &res)
	return res
}

func helper(nums []int, pointer int, list []int, res *[][]int){
	temp := make([]int, len(list))
	copy(temp, list)
	*res = append(*res, temp)

	for i:=pointer;i<len(nums);i++{
		list = append(list, nums[i])
		fmt.Println("befor",i, list)
		helper(nums, i+1, list, res)
		list = list[0: len(list)-1]
		fmt.Println("after",i, list)
	}
}

func main(){
	subsets([]int{0,1,2})

}
