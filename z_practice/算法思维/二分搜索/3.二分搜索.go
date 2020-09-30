package main

import "fmt"

/*
	35. Search Insert Position
	Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
	You may assume no duplicates in the array.
	Example 1:
	Input: [1,3,5,6], 5
	Output: 2
	Example 2:

	Input: [1,3,5,6], 2
	Output: 1
	Example 3:

	Input: [1,3,5,6], 7
	Output: 4
	Example 4:

	Input: [1,3,5,6], 0
	Output: 0

 */

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)-1

	for left+1 < right{
		mid := (left + right)/2
		if nums[mid] > target{
			right = mid
		}else if nums[mid] < target{
			left = mid
		}else{
			return mid
		}
	}
	if nums[left] >= target{
		return left
	}else if nums[right] >= target{
		return right
	}else if nums[right] < target{
		return right  +1
	}
	return 0
}

func main(){

	fmt.Println(searchInsert([]int{1,3,5,6}, 2))

}