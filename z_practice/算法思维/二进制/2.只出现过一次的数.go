package main

import "fmt"

/*
	137. 只出现一次的数字 II
		给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
		说明：
		你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
		示例 1:
			输入: [2,2,3,2]
			输出: 3
		示例 2:
			输入: [0,1,0,1,0,1,99]
			输出: 99
*/

func singleNumber2(nums []int) int {
	res := 0
	for i:=0;i<=64;i++{
		sum := 0
		for j:=0;j<len(nums);j++{
			//移位后统计个数
			sum += (nums[j] >> i) &1
		}
		// 归位，注册到res中
		res ^= (sum%3)<<i
	}
	return res
}

func main(){
	fmt.Println(singleNumber2([]int{0,1,0,1,0,1,99}))

}