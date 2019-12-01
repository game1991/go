package main

import "fmt"

func testa() {

	fmt.Println("aaaaaaaaaa")
}

func testb(x int) {
	//设置recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Println(error)
		}

	}() //()意味着调用此匿名函数
	var a [10]int
	a[x] = 111
}
func testc() {

	fmt.Println("ccccccccccc")
}

func main() {
	testa()
	testb(20)
	testc()
}
