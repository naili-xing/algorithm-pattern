package main

import "fmt"

/*
300. Longest Increasing Subsequence
Given an unsorted array of integers, find the length of longest increasing subsequence.

Example:

Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Note:

There may be more than one LIS combination, it is only necessary for you to return the length.
Your algorithm should run in O(n2) complexity.
Follow up: Could you improve it to O(n log n) time complexity?
 */

func lengthOfLIS(nums []int) int {

	it := make([]int, len(nums))

	for i:=0; i<len(nums);i++{
		it[i] = 1
		for j:=0; j<i;j++{
			if nums[j] < nums[i]{
				it[i] = max(it[i], it[j]+1)
			}
		}
	}
	res := 0
	for i:=0;i<len(it);i++{
		if it[i] > res{
			res = it[i]
		}
	}
	return res
}

func max(x,y int)int{
	if x>=y{
		return x
	}else{
		return y
	}
}

func main(){
	fmt.Println(lengthOfLIS([]int{1,3,6,7,9,4,10,5,6}))

}