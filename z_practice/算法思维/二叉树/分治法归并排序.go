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


func MergeSort(nums []int) []int {

	res := merge(nums)
	return res

}
func merge(nums []int)[]int{
	var res []int
	if len(nums) ==1{
		res = append(res, nums[0])
		return res}
	if len(nums) == 2{
		if nums[0] > nums[1]{
			res = append(res, nums[1])
			res = append(res, nums[0])
		} else{
			res = append(res, nums[0])
			res = append(res, nums[1])
		}
		return res
	}
	i := len(nums)/2
	left := merge(nums[:i])
	right := merge(nums[i:])
	return combine(left, right)
}

func combine(nums1 []int, nums2 []int) []int {
	var res []int

	i:=0
	j:=0
	for i<len(nums1) && j <len(nums2){
		if i<len(nums1) && j<len(nums2){
			if nums1[i]>nums2[j]{
				res = append(res, nums2[j])
				j+=1
			}else{
				res = append(res, nums1[i])
				i+=1
			}
		}
	}
	res = append(res, nums1[i:]...)
	res = append(res, nums2[j:]...)
	return res

}

func main(){
	fmt.Println(MergeSort([]int{5,1,1,2,0,0}))
}