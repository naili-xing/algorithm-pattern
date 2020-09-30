package main

import "fmt"

func searchRange(nums []int, target int) []int {
	res := make([]int, 2)
	left := 0
	right := len(nums)-1

	for left+1 < right{
		mid := (left +right)/2
		if nums[mid] > target{
			right = mid
		}else if nums[mid] < target{
			left = mid
		}else{
			right = mid
		}
	}
	//搜左边


	if nums[left] == target{
		res[0] = left
	}else if nums[right] == target{
		res[0] = right
	}else{
		res[0] = -1
		res[1] = -1
		return res
	}
	left = 0
	right = len(nums)-1
	for left+1 < right{
		mid := (left +right)/2
		if nums[mid] > target{
			right = mid
		}else if nums[mid] < target{
			left = mid
		}else{
			left = mid
		}
	}
	//搜左边

	if nums[right] == target{
		res[1] = right
	}else if nums[left] == target{
		res[1] = left
	}else{
		res[0] = -1
		res[1] = -1
		return res
	}
	return res
}


func  main(){
	fmt.Println(searchRange([]int{2,2}, 2))
}