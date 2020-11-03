package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string `json:"n"`
	age  int    `json:"a"`
}

func (s Student) P() {
	fmt.Println("打印")
}

func t1(b interface{}) {
	rt := reflect.TypeOf(b)
	rv := reflect.ValueOf(b)
	kd := rv.Kind()

	if kd != reflect.Struct {
		fmt.Println("this is not the strucutre")
		return
	}
	nums := rv.NumField()
	for i := 0; i < nums; i++ {
		fmt.Println("value is ", rv.Field(i))
		tagVal := rt.Field(i).Tag.Get("json")
		fmt.Println("tag is ", tagVal)
	}
	numMethod := rv.NumMethod()
	fmt.Println(numMethod)
	rv.Method(0).Call(nil)

}

func main() {

	//stu := new(Student)
	//stu.name= "naili"
	//stu.age= 1000

	var stu Student = Student{"naili", 1000}

	t1(stu)
}
