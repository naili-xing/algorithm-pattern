package main

import "fmt"

/*
	给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

	示例 :

	输入: [1,2,1,3,2,5]
	输出: [3,5]

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/single-number-iii
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */

func singleNumber3(nums []int) []int {

	res := 0
	for i:=0;i<len(nums);i++{
		res ^= nums[i]
	}
	rest := []int{res, res}
	//res = (res&(res-1))^res
	res &= -res
	for i:=0;i<len(nums);i++{
		if res & nums[i] == 0{
			rest[0] ^= nums[i]
		}else{
			rest[1] ^= nums[i]
		}
	}
	return rest
}

func main(){
	fmt.Println(singleNumber3([]int{1,2,1,3,2,5}))

}