package main

import "fmt"

func main() {

	//对比数组， 数组只存一类数据
	var a [3]int
	//这个地址是数组的首个地址
	fmt.Printf("%v,%p\n", a, &a)
	fmt.Print(&a[0], &a[1], &a[2], a[:], "\n")

	//切片：切片是数组的引用，所以切片是引用类型， 切片有3中创建方式

	// 切片是一个数据结构，类似于
	/*
		struct slice{
			prt *[2]int，存的是他所指向的数组的首个元素的地址
			len int
			cap int
		}
	*/

	//切片创建方式1：从一个array或者string创建
	slice := a[0:2]
	slice[1] = -1
	slice[0] = -3
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
	fmt.Println(&slice[6])
	fmt.Println(slice, a)

	//切片创建方式2：make， 用make创建的， 会生成一个数组，但是不可见，只能通过slice操作
	var sl2 = make([]int, 5, 30)
	sl2[0] = 1
	sl2[1] = 10
	fmt.Println(sl2)

	//切片创建方式3：只声明,后续增加，或者带默认值的
	var s3 []int
	var s4 = []int{1, 2, 3, 4}

	s3 = append(s3, 1)
	s3 = append(s3, 2)
	s3 = append(s3, 3)
	fmt.Println(s3, s4)

}
