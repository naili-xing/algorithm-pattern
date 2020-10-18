package main

import "fmt"

func main() {

	var ch chan int = make(chan int, 3)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 3
		ch <- 3
		ch <- 3
		close(ch)
	}()

	go func() {
		<-ch
		<-ch
		<-ch
	}()

	for {
		if v, ok := <-ch; ok {
			fmt.Println(v)
		} else {
			break
		}
	}

	fmt.Println("asdf")

}
