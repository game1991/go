package main

import "fmt"

func main() {

	var p *int
	//指向一个合法的内存

	//p是*int类型，指向int类型
	p = new(int)
	*p = 666
	fmt.Println("*p= ", *p)
	q := new(int)
	*q = 777
	fmt.Println("*q= ", *q)
}
