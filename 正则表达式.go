package main

import (
	"fmt"
	"regexp"
)

func main() {

	a := "abc azc a7c aac 888 a9c tac"
	//1、解释规则，它会解析正则表达式，如果成功返回解释器
	a1 := regexp.MustCompile(`a.c`)
	if a1 == nil { //说明解析失败，返回nil
		fmt.Println("regexp error")
		return
	}
	//2、根据规则提取关键信息
	result1 := a1.FindAllStringSubmatch(a, -1) //-1代表匹配所有
	fmt.Println("result= ", result1)

}
