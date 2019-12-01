package main

import "fmt"

// func funcc(c int) {
// 	fmt.Println("c= ", c)
// }
// func funcb(b int) {
// 	funcc(b - 1)
// 	fmt.Println("b= ", b)

// }
// func funca(a int) {
// 	funcb(a - 1)
// 	fmt.Println("a= ", a)
// }
// func main() {

// 	funca(3)
// 	fmt.Println("main")

// }
func test(a int) {
	if a == 1 { //函数终止调用的条件
		fmt.Println("a= ", a)
		return //终止函数调用
	}
	test(a - 1)
	fmt.Println("a= ", a)
}

func main() {
	test(3)
	fmt.Println("main")
}
