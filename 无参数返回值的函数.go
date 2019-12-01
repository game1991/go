package main

import "fmt"

func Myfunc() {
	a := 666
	fmt.Println("a= ", a)
}
func main() {
	//无参数无返回值函数的调用：函数名（）
	Myfunc()
}
