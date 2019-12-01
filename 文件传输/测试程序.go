package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// for {
	// 	var in string
	//
	// 	fmt.Scan(&in)
	// 	if in == "ok" {
	// 		break
	// 	} else {
	// 		fmt.Println("输入错误，确认接收请输入：ok")
	// 	}

	// }
	// data := make([]byte, 1024)
	// data = []byte(in)
	var s string
	for {
		fmt.Println("请输入接收文件确认信息：")
		r := bufio.NewReader(os.Stdin) //从标准输入设备读取
		in, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("err=", err)
		}
		if in == "ok\r\n" {
			s = in
			break
		} else {
			fmt.Println("输入错误，确认接收请输入：ok")
		}

	}
	fmt.Printf("s=%c", []byte(s))

}
