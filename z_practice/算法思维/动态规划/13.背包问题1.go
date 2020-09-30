package main

import "fmt"

/*
描述
在n个物品中挑选若干物品装入背包，最多能装多满？假设背包的大小为m，每个物品的大小为A[i]

你不可以将物品进行切割。

您在真实的面试中是否遇到过这个题？
样例
样例 1:
	输入:  [3,4,8,5], backpack size=10
	输出:  9

样例 2:
	输入:  [2,3,5,7], backpack size=12
	输出:  12

挑战
O(n x m) 的时间复杂度 and O(m) 空间复杂度
如果不知道如何优化空间O(n x m) 的空间复杂度也可以通过.

 */
/**
 * @param m: An integer m denotes the size of a backpack
 * @param A: Given n items with size A[i]
 * @return: The maximum size
 */
func backPack (m int, A []int) int {

	dp := make([]int,m+1)
	dp[0] = 0
	for i:=0;i<=m;i++{
		for j:=0;j<len(A);j++{
			if i-A[j] >=0{
				dp[i] = max5(dp[i-1], dp[i-A[j]]+A[j])
			}
		}
	}
	return dp[len(dp)-1]

}

func max5(x, y int)int{
	if x> y{
		return x
	}else{
		return y
	}
}

func main(){
	fmt.Println(backPack(10, []int{3,4,8,5}))
}
