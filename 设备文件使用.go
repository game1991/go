package main

import (
	"fmt"
	"os"
)

func main() {

	os.Stdout.WriteString("are u ok?") //往标准输出设备(屏幕)写内容
	//os.Stdout()默认打开，用户可以直接使用
	var a int
	fmt.Println("请输入a: ")
	fmt.Scan(&a) //从标准输入设备（键盘）中读取内容，放在a
	fmt.Println("a = ", a)
	defer os.Stdin.Close() //关闭后，无法输出
}
