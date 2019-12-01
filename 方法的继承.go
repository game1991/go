package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//Person类型实现了一个方法
func (one *Person) PrintInfo() {

	fmt.Printf("name=%s, sex=%c, age=%d\n", one.name, one.sex, one.age)
}

//有个学生继承Person，成员和方法都继承了
type Student struct {
	Person
	id   int
	addr string
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "hb"}
	s.PrintInfo()
}
