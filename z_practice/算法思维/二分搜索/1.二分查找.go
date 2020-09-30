package main
import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	for left+1 < right{
		mid := (left + right)/2
		if nums[mid] == target{
			return mid
		}
		if nums[mid] < target{
			left = mid
		}
		if nums[mid] > target{
			right = mid
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

func  main(){
	fmt.Println(search([]int{-1,0,3,5,9,12}, 2))
}

