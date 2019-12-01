package main

import (
	"errors"
	"fmt"
)

func Mydiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}
	return
}

func main() {

	/*err1 := fmt.Errorf("%s", "this a  normal err")
	fmt.Println("err= ", err1)
	err2 := errors.New("this an other err")
	fmt.Println("err2= ", err2)
	*/ //这是error使用的方法
	result, err := Mydiv(10, 0)
	if err != nil {
		fmt.Println("err= ", err)
	} else {
		fmt.Println("result= ", result)
	}
}
