package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err=", err)
		return
	}
	defer f.Close()

	game := make([]byte, 1024*4) //4K缓冲区
	//读文件内容，读多少写多少，一点不差
	for {
		n, err := f.Read(game)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err=", err)
			}
			return
		}
		//发送内容
		conn.Write(game[:n]) //给服务器发送内容
	}
}

func main() {
	//提示输入文件名
	fmt.Println("请输入需要传输的文件：")
	var path string
	fmt.Scan(&path)
	//获取文件名info.Name
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err=", err)
		return
	}

	//主动连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("net.Dial err=", err1)
		return
	}
	defer conn.Close()
	//给接收方先发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	//接收对方回复，如果回复“ok”,说明对方准备好了，可以发文件
	game := make([]byte, 1024)
	n, err := conn.Read(game)
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	if "ok\r\n" == string(game[:n]) {
		//发送文件内容
		SendFile(path, conn)
	}

}
