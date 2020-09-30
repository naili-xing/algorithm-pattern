package main

import (
	"fmt"
)

//给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。说明：解集不能包含重复的子集。


func subsetsWithDup(nums []int) [][]int {

	var rs [][]int
	var list []int
	helper2(nums, 0, list, &rs)
	return rs

}

func helper2(nums []int, pos int, list []int, res *[][]int ){

	ans := make([]int, len(list))
	copy(ans, list)
	*res = append(*res, ans)

	for i := pos; i<len(nums);i++{
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		list = append(list, nums[i])
		helper2(nums, i+1, list, res)
		list = list[0:len(list)-1]
	}
}

func main(){
	fmt.Println(subsetsWithDup([]int{1,2,2}))

}
