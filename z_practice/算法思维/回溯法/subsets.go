package main

import "fmt"

//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

func subsets(nums []int) [][]int {

	var res [][]int
	var list []int

	helper(nums, 0 , list, &res)
	return res

}

func helper(nums []int, pointer int, list []int, res *[][]int){

	temp :=make([]int, len(list))
	copy(temp, list)
	*res = append(*res, temp)

	for i:=pointer; i<len(nums); i++{
		list = append(list, nums[i])
		helper(nums, i+1, list, res)
		list = list[0: len(list)-1]
	}

}

func main(){
	fmt.Println(subsets([]int{0,1,2}))

}
