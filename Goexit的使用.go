package main

import (
	"fmt"
	//"runtime"
)

func test() {
	defer fmt.Println("hello world")
	return //终止此函数
	//runtime.Goexit() //终止所在的协程
	fmt.Println("world")
}

func main() {

	go func() {
		fmt.Println("go")
		//调用一个别的函数
		test()
		fmt.Println("hello")

	}()
	//特地设置一个死循环，不让主协程结束
	for {

	}

}
