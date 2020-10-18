package main

import "fmt"

//主要是new, make, len 3个
func t3() {

	// new分配内存，返回的是新分配的(该类型的)0值的指针
	a := new(int)
	fmt.Printf("%T,%v, %v, %v\n", a, a, &a, *a)
	*a = 100
	fmt.Printf("%T,%v, %v, %v", a, a, &a, *a)

	// make分配内存，返回的是新分配的(该类型的)0值的指针

}

func main() {
	t3()
}
