package main

import "fmt"

/*
33. 搜索旋转排序数组

	假设按照升序排序的数组在预先未知的某个点上进行了旋转。
	( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

	搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
	你可以假设数组中不存在重复的元素。
	你的算法时间复杂度必须是 O(log n) 级别。
	示例 1:
	输入: nums = [4,5,6,7,0,1,2], target = 0
	输出: 4
	示例 2:
	输入: nums = [4,5,6,7,0,1,2], target = 3
	输出: -1
 */

func search5(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums)-1

	for left+1 < right{
		mid := (left + right)/2
		if nums[mid] == target{return mid}
		// 说明mid是属于左边的那个升序队列的
		if nums[mid] > nums[left]{
			if nums[mid] >= target && nums[left] <= target{
				//target 在左边的队列
				right = mid
			}else{
				// target 在右边的队列
				left = mid
			}
			// 说明mid是属于右边的那个升序队列的
		}else if nums[mid] < nums[right]{
			if nums[mid] <= target && nums[right]>= target{
				//target 在左边的队列
				left = mid

			}else{
				// target 在右边的队列
				right = mid
			}
		}
	}
	if nums[left] == target{
		return left
	}
	if nums[right] == target{
		return right
	}
	return -1
}

func main(){
	fmt.Println(search5([]int{5,1,3}, 3))
}
