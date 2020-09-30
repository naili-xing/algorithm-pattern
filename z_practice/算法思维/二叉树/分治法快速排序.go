package main

import "fmt"

/*
	给你一个整数数组 nums，请你将该数组升序排列。

	示例 1：

	输入：nums = [5,2,3,1]
	输出：[1,2,3,5]
	示例 2：

	输入：nums = [5,1,1,2,0,0]
	输出：[0,0,1,1,2,5]
	 
	提示：
	1 <= nums.length <= 50000
	-50000 <= nums[i] <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


func sortArray(nums []int) []int {
	quicksort(nums,0, len(nums)-1)
	return nums
}

func quicksort(nums []int, begin int, end int){
	if begin < end{
		pro := partition(nums, begin ,end)
		quicksort(nums, begin, pro-1)
		quicksort(nums, pro+1, end)
	}
}

func partition(nums []int, begin int ,end int) int{
	rn := nums[begin]
	i := begin
	j := end

	for i < j {
		//必须先从右找后从左边找
		for i < j && rn<=nums[j]{
			j--
		}
		for i<j && rn>=nums[i]{
			i++
		}

		swap(nums, i, j)
	}
	swap(nums, begin, i)
	return i
}

func swap(nums []int, i int, j int){
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

func main(){
	fmt.Println(sortArray([]int{5,1,1,2,0,0}))
}