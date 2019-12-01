package main

import (
	"fmt"

	"time"
)

func main() {
	//创建管道channel
	ch := make(chan string)
	defer fmt.Println("主协程也调用结束")

	go func() {
		defer fmt.Println("子协程调用完毕")
		for i := 0; i < 5; i++ {

			fmt.Println("子协程i= ", i)
			time.Sleep(time.Second)

		}
		ch <- "我是子协程，工作完毕"
	}()
	str := <-ch //没有数据前阻塞
	fmt.Println("str= ", str)

}
