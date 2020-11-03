package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock sync.Mutex
)

func t9(n int, mapper map[int]int64) {

	//fmt.Println("num of cpu is",runtime.NumCPU())

	res := n
	//加全局锁
	lock.Lock()
	mapper[n] = int64(res)
	lock.Unlock()
}

func t10(n int, ch chan int, wg *sync.WaitGroup) {
	defer (*wg).Done()
	res := n
	ch <- res
}

func main() {

	//通信1：全局加索方式
	var mapper = make(map[int]int64, 200)
	// 给map写东西不能并发的写，同一时间只能一个写，
	for i := 1; i < 200; i++ {
		go t9(i, mapper)
	}
	//lock.Lock()
	//for i,v := range mapper{
	//	fmt.Println(i, v)
	//}
	//lock.Unlock()
	//通信2：使用channel， + wg grup
	var map2 = make(map[int]int64, 200)
	var ch = make(chan int, 200)
	var wg sync.WaitGroup

	for i := 1; i < 200; i++ {
		wg.Add(1)
		go t10(i, ch, &wg)
	}
	wg.Wait()
	close(ch)

	//for v:= range ch{
	//	map2[v] = int64(v)
	//}
	for {
		if v, ok := <-ch; ok {
			map2[v] = int64(v)
		} else {
			break
		}
	}
	for k, v := range mapper {
		fmt.Println("key is ", k, "V is ", v)
	}

	//通信3：select
	intchan := make(chan int, 10)
	for i := 0; i < 10; i++ {

		go func(i int) {
			time.Sleep(time.Second * 20)
			intchan <- i
		}(i)
	}
	strchan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		strchan <- "hellow" + fmt.Sprintf("%d", i)
	}
	for {
		select {
		// 如果没关闭管道，不回一致阻塞而死锁，会自动向下匹配
		case v := <-intchan:
			fmt.Println("read from inthan", v)
		case v := <-strchan:
			fmt.Println("read from strchan", v)
		default:
			fmt.Println("nog found anything")
			return
		}
	}
}
