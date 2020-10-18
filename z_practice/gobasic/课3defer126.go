package main

import "fmt"

// 执行到defer时候，压入独立栈中，执行完后，后进先出，依次出栈
// 压入栈时，相关的值会拷贝到栈

func test2(a, b int) int {
	defer fmt.Println("defer 1", a)
	defer fmt.Println("defer 2", b)
	a++
	b++
	fmt.Println(a, b)
	return 1
}

func main() {
	test2(1, 2)
}
