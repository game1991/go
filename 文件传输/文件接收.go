package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func RecvFlie(FlieName string, conn net.Conn) {
	game := make([]byte, 1024*4)
	f, err := os.Create(FlieName)
	if err != nil {
		fmt.Println("os.Create err =", err)
		return
	}
	defer f.Close()

	//接收多少，写多少，一点不差
	for {
		n, err := conn.Read(game) //接收对方发送的文件
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn Read err =", err)
			}
			return
		}

		// if n == 0 {
		// 	fmt.Println("n==0 文件接收完毕")
		// 	break
		// }

		f.Write(game[:n]) //往文件写入内容
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	//阻塞等待用户连接
	conn, err1 := listener.Accept()
	if err != nil {
		fmt.Println("err1=", err1)
		return
	}
	defer conn.Close()

	game := make([]byte, 1024)
	var n int
	n, err = conn.Read(game) //读取对方发送的文件名
	if err != nil {
		fmt.Println("conn Read err=", err)
		return
	}
	FlieName := string(game[:n])

	//回复"ok"
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

	conn.Write([]byte(s))
	//接收文件内容
	RecvFlie(FlieName, conn)

}
