package main

import (
	"fmt"
	"strings"
)

func main() {
	//"hellogo"中是否包含"hello"
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "one"))
	//Joins组合
	s := []string{"abc", "hello", "go"}
	a := strings.Join(s, "xxx")
	fmt.Println("a= ", a)
	//Index,查找子串的位置
	fmt.Println(strings.Index("abcdhello", "hello"))
	fmt.Println(strings.Index("abcdhello", "go"))
	//Repeat
	a = strings.Repeat("go", 3)
	fmt.Println("a= ", a)
	//Split 以指定的分隔符拆分
	a = "abc***hello*go*"
	s2 := strings.Split(a, "*")
	fmt.Println("s2= ", s2)
	//Trim
	a = strings.Trim("   are u ok    ", " ") //去掉两边空格
	fmt.Printf("a= #%s#\n", a)
	a = strings.Trim("***are  u**ok***", "*") //去掉两边空格
	fmt.Printf("a= #%s#\n", a)

	//Fields
	s3 := strings.Fields("   are     u   ok  ")
	fmt.Println("s3= ", s3)
	/*for i, data := range s3 {
		fmt.Println(i, ",", data)
	}*/
}
