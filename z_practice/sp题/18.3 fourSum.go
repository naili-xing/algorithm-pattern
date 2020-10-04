package main

import "sort"

/*
18. 四数之和
	给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
	注意：
	答案中不可以包含重复的四元组。
	示例：
	给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
	[
	  [-1,  0, 0, 1],
	  [-2, -1, 1, 2],
	  [-2,  0, 0, 2]
	]
 */
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int

	// 从第0个到倒数第四个
	for i := 0; i < n-3; i++ {
		// 如果重复，跳过
		if i > 0 && nums[i] == nums[i-1] {continue}

		// 从第1个到倒数第3个
		for j := i+1; j < n-2; j++ {

			// 如果重复，跳过
			if  j > i+1 && nums[j] == nums[j-1] {continue}

			// 双指针夹逼
			lo, hi := j + 1, n - 1

			for lo < hi {
				// 如果相等，添加
				if nums[i] + nums[j] + nums[lo] + nums[hi] == target {
					res = append(res, []int{nums[i], nums[j], nums[lo], nums[hi]})

					// 找到第一个不重复的
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}
					lo++
					hi--

					// 如果太大，右指针左移
				} else if nums[i] + nums[j] + nums[lo] + nums[hi] > target{
					hi--

					// 如果太小，左指针右移动
				} else {
					lo++
				}
			}
		}
	}
	return res
}


func fourSum2(nums []int, target int) [][]int {

	sort.Ints(nums)
	res := [][]int{}
	for i:=0;i<len(nums);i++{
		if i>=1 && nums[i] == nums[i-1]{continue}
		for j:=i+1;j<len(nums);j++{
			if j>i+1 && nums[j] == nums[j-1]{continue}
			l :=j+1
			r := len(nums)-1
			for l<r{
				if nums[i]+nums[j]+nums[l] + nums[r] == target{
					res = append(res, []int{nums[i],nums[j],nums[l] , nums[r]})
					for l<r && nums[l]==nums[l+1]{
						l++
					}
					for l<r && nums[l]==nums[r-1]{
						r--
					}
					l++
					r--
				}else if nums[i]+nums[j]+nums[l] + nums[r]>target{
					r--
				}else{
					l++
				}
			}
		}
	}
	return res
}

func main(){

}