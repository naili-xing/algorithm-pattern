package main

import "fmt"

type usb interface {
	start()
	stop()
}

type phone struct {
}

func (p phone) start() {
	fmt.Println("Phone start")
}

func (p phone) stop() {
	fmt.Println("phone stop")
}

func (p phone) call() {
	fmt.Println("im a phone")
}

type camera struct {
}

func (p camera) start() {
	fmt.Println("computer start")
}

func (p camera) stop() {
	fmt.Println("computer stop")
}

func (p camera) sssss() {
	fmt.Println("computer stop")
}

type computer struct {
}

//实现usb接口就是指实现了usb接口声明的所有方法。
func (c computer) working(usb usb) {
	usb.start()
	//类型断言 pyhone想调用call
	if phone, ok := usb.(phone); ok {
		phone.call()
	}
	usb.stop()
}

func typeJudge(item ...interface{}) {
	for index, v := range item {

		switch v.(type) {
		case bool:
			fmt.Println("this is bool", index, v)
		case float32:
			fmt.Println("this is float32", index, v)
		}
	}
}

func main() {
	//测试
	var usbarr [3]usb
	usbarr[0] = phone{}
	usbarr[1] = camera{}
	usbarr[2] = phone{}

	var computer = computer{}
	for _, v := range usbarr {
		computer.working(v)
	}

	typeJudge(10, true, 1.1)

}
