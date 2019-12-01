package main

import (
	"fmt"

	"time"
)

//定义一个打印机，参数为字符串，按每个字符打印

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println("\n")
}

func Person1() {
	Printer("hello")
}

func Person2() {
	Printer("world")
}

func main() {
	//新建2个协程，代表2个人同时使用打印机
	go Person1()
	go Person2()
	//特定不让主协程结束，设置一个死循环
	for {

	}

}
