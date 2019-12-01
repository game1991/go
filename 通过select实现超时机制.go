package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num =", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
				break
			}

		}

	}()
	for i := 0; i < 6; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
	<-quit
	fmt.Println("程序结束")
}
