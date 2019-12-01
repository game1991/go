package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//带有接收者的函数叫做方法
func (one Person) PrintInfo() {
	fmt.Println("one= ", one)

}

//通过一个函数，给成员赋值
func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

func main() {

	p := Person{"game", 'm', 18}
	p.PrintInfo()

	//定义一个结构体变量
	var a Person
	(&a).SetInfo("yoyo", 'f', 18)
	a.PrintInfo()
}
