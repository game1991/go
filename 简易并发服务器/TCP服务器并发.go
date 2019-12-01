package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	//函数调用完毕，自动关闭
	defer conn.Close()

	//获取客户端网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connet sucessful!")
	game := make([]byte, 2048)
	for { //读取用户数据
		n, err := conn.Read(game)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		fmt.Printf("[%s]:%s\n", addr, string(game[:n]))
		// fmt.Println("len=", len(string(game[:n])))
		// if "exit" == string(game[:n-1]) {//nc测试的
		if "exit" == string(game[:n-2]) { //自己写的客户端测试的，发送时多了两个字符，\r\n
			fmt.Println("connet exit")
			return
		}

		//把数据转换成大写，再给用户发送
		conn.Write([]byte(strings.ToUpper(string(game[:n]))))

	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	defer listener.Close()
	//接受多个用户请求,建立协程
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		//处理用户请求数据
		go HandleConn(conn)
	}

}
