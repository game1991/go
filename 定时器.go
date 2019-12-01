package main

import (
	"fmt"

	"time"
)

func main02() {

	time.Sleep(2 * time.Second)
	fmt.Println("时间到")

}

func main() {

	<-time.After(2 * time.Second) //定时2秒，阻塞2秒，2秒后产生一个事件，往channel写内容
	fmt.Println("时间到")

}
func main01() {

	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")

}
