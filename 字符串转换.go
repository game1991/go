package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 5566, 10) //第二个数为追加的数，第三个为指定10进制方式追加
	slice = strconv.AppendQuote(slice, "hellogo")

	fmt.Println("slice= ", string(slice)) //转换成string再打印
	//其他类型转换成字符串
	var str string
	str = strconv.FormatBool(false)
	//'f'指的是打印格式，以小数方式，-1指的是小数点位数(紧缩模式),64以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	//整型转字符串，常用
	str = strconv.Itoa(666)
	fmt.Println("str= ", str)
	//字符串转其他类型
	var x bool
	var y error
	x, y = strconv.ParseBool("tuae")
	if y == nil {
		fmt.Println("x= ", x)
	} else {
		fmt.Println("y= ", y)
	}
	//把字符串转换成整型
	a, _ := strconv.Atoi("789")
	fmt.Println("a= ", a)

}
