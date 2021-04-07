package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*ch1 := make(chan int)

	go func() {
		fmt.Println("子goroutine开始执行。。")
		//time.Sleep(3 * time.Second)
		data := <-ch1 //从ch1中读取数据
		fmt.Println("data：", data)
	}()

	time.Sleep(5 * time.Second)
	ch1 <- 10
	fmt.Println("main..over...")

	ch := make(chan int)
	ch <- 100 //阻塞*/
	Test()
}

func Test() {
	ch3 := make(chan bool)
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		defer wg.Done()
		<-ch3
		fmt.Println("goroutine run")

	}()
	time.Sleep(3 * time.Second)
	ch3 <- true
	wg.Wait()
	fmt.Println("main run")

	ch3 <- false
}
