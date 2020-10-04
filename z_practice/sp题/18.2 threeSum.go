package main

import (
	"fmt"
	"sort"
)

/*
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

 */

func threeSum(nums []int) [][]int {
	//核心是循环固定前2个数，然后移动3。知道相加为0，
	var res [][]int
	sort.Ints(nums)
	for i:=0;i<len(nums);i++{
		if i>=1&&nums[i] == nums[i-1]{continue}
		target := 0-nums[i]
		k := len(nums)-1
		for j:=i+1;j<len(nums);j++{
			if j>i+1&&nums[j] == nums[j-1]{continue}

			for j<k && nums[j] + nums[k] > target{
				k--
			}

			if k==j{break}

			if nums[j] + nums[k] == target{
				res = append(res, []int{nums[i],nums[j],nums[k]})
			}
		}
	}
	return res
}


func main(){
	fmt.Println(threeSum([]int{3,0,-2,-1,1,2}))
}