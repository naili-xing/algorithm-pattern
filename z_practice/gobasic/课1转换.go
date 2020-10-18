package main

import (
	"fmt"
	"strconv"
)

func init() {
	fmt.Println("this it the init")
}

func test() {
	a := 1
	b := 1.3
	c := []byte{'a', 'b', 'c', 'd'}

	//其他 to string
	//int to string
	var o1 string = strconv.Itoa(a)
	var o2 string = strconv.FormatInt(int64(a), 10)
	var o3 string = fmt.Sprintf("%d", a)

	//float to string
	var o4 string = fmt.Sprintf("%f", b)

	//bytes to string
	var o5 string = string(c)

	fmt.Println(o1, o2, o3, o4, o5)

	//string to 其他
	aa := "true"
	bb := "123"
	cc := "asdfasdfasdfasd中文应该不行"
	//str to bool
	res, _ := strconv.ParseBool(aa)

	//str to int
	res2, _ := strconv.ParseInt(bb, 10, 8)

	//str to byte
	res3 := []byte(bb)

	fmt.Print(res, res2, res3)

	//字符串遍历,最好转换成rune再遍历，他能处理中文问题
	r := []rune(cc)
	fmt.Println(r)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len([]byte(cc)); i++ {
		fmt.Printf("%c", []byte(cc)[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(cc); i++ {
		fmt.Printf("%c", cc[i])
	}
}

func main() {
	test()
}
