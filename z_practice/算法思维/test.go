package main

import (
	"fmt"
	"time"
)

func main() {

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("expired")
	})

	fmt.Println("reached")
	fmt.Println("reached")
	fmt.Println("reached")
	fmt.Println("reached")
	time.Sleep(time.Second * 10)
}
