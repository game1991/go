package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for i := 0; i < 2; i++ {
			runtime.Gosched() //主协程结束先退出，导致子协程未能执行调用
			fmt.Println("go")
		}

	}()

	for i := 0; i < 5; i++ {
		//让出时间片，先让别的协程执行，他执行完，再回来执行此协程
		fmt.Println("hello")
	}

}
