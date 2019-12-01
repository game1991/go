package main

import "fmt"

func main() {
	//var a [10]int=[5]int{1,2,3,4,5}
	//初始化数组，两种写法都可以
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("a= ", a)
	b := [5]int{1, 2, 5}
	fmt.Println("b= ", b)
	c := [5]int{2: 5, 4: 10}
	fmt.Println("c= ", c)

}
