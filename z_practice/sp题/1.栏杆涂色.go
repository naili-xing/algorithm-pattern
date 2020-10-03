package main

import "fmt"

func numWays(n int, k int) int {
	// n, 栏杆数目
	// k, 颜色个数

	if n == 0{
		return 0
	}
	if n == 1{
		return k
	}
	if n == 2{
		return k*k
	}
	slow:=k
	fast:=k*k
	for i:=3;i<=n;i++{
		//temp := fast
		fast ,slow = slow * (k-1) + fast * (k-1), fast
		//slow = temp
	}
	return fast
}

func main(){
	fmt.Println(numWays(3, 2))

}