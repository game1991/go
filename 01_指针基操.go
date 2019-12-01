package main

import "fmt"

func main() {
	var a int = 10
	//每个变量都有两层含义：变量的内存，变量的地址
	fmt.Printf("a= %d\n", a) //变量的内存
	fmt.Printf("&a= %v\n", &a)

	//保存某个变量的地址，需要指针类型  *int保存int的地址 **int保存*int的地址
	//定义一个变量p,类型为*int
	var p *int
	p = &a //指针变量指向谁，就把谁的地址赋值给指针变量
	fmt.Printf("p= %v\n,&a= %v\n", p, &a)

	*p = 666 //*p操作的不是p的内存，是p指向的内存（也就是a）
	fmt.Printf("*p= %v\n,&a= %v\n", *p, a)

}
