package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dialnet.Dial err=", err)
		return
	}
	//main调用完毕，关闭连接
	defer conn.Close()
	//从键盘输入内容，给服务器发送内容，新建协程
	go func() {
		//切片缓冲
		buf := make([]byte, 1024)

		for {
			n, err := os.Stdin.Read(buf) //从键盘读取内容，放在buf
			if err != nil {
				fmt.Println("os.Stdin.Read err=", err)
				return
			}
			//把输入的内容给服务器发送
			conn.Write(buf[:n])
		}

	}()
	//接受服务器回复的数据
	game := make([]byte, 1024)
	for {
		n, err := conn.Read(game) //接受服务器的请求
		if err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		//打印接收到的信息
		fmt.Println(string(game[:n]))

	}

}
