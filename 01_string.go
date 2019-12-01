package main

import "fmt"

func main() {
	var str1 string
	str1 = "adc"
	fmt.Println("str1=", str1)
	str2 := "adc"
	fmt.Printf("str2类型是%T\n", str2)
	fmt.Printf("len(str2)=%d", len(str2))

}
