package main

import (
	"fmt"

	"time"
)

//全局变量，创建一个channel

var ch = make(chan int)

//定义一个打印机，参数为字符串，按每个字符打印

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(5 * time.Second)
	}
	fmt.Println("\n")
}

//Person1执行完才能Person2执行
func Person1() {
	Printer("hello")
	ch <- 666 //给管道写数据，发送
}

func Person2() {

	<-ch //从管道取数据，如果没有数据管道它就会阻塞
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
