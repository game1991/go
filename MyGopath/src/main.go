package main

import (
	"calc"
	"fmt"
)

func init() {
	fmt.Println("this is main init!!!")

}

func main() {
	a := calc.Add(10, 20)
	fmt.Println("a= ", a)
	fmt.Println("r= ", calc.Mins(20, 10))
}
