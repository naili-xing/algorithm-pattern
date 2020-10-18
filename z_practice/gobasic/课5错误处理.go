package main

import (
	"errors"
	"fmt"
)

func eh() {
	// recover 可以捕获异常
	if err := recover(); err != nil {
		fmt.Println("fuck u ,u just fucked up, error is  \"", err, "\"")
	}
}

// 自定义错误
func doSomething(i int) error {
	if i == 1 {
		return nil
	} else {
		return errors.New("fuck u ,i need 1")
	}
}

func t5() {
	// defer+recover
	defer eh()
	if err := doSomething(123); err != nil {
		panic(err)
	}

}

func main() {
	t5()
}
