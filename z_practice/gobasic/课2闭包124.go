package main

import "fmt"

//闭包，返回一个匿名函数，这时候匿名函数和n构成一个整体， 构成闭包， n初始化一次而已
// 闭包的关键是分析出他使用的哪些变量，

func AddIt() func(int) int {
	n := 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	a := AddIt()

	fmt.Println(a(1))
	fmt.Println(a(1))
	fmt.Println(a(1))
	fmt.Println(a(1))
}
