package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}
	defer conn.Close()
	requestBuf := "GET /mike/go HTTP/1.1\r\nHost: 127.0.0.1:8000\r\nConnection: keep-alive\r\nCache-Control: max-age=0\r\nUpgrade-Insecure-Requests: 1\r\nUser-Agent: Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3\r\nAccept-Encoding: gzip, deflate, \r\nAccept-Language: zh-CN,zh;q=0.9,ja;q=0.8,en;q=0.7\r\n\r\n"

	//先发请求包，服务器才会回响应包

	conn.Write([]byte(requestBuf))

	//接收服务器回复的响应包

	buf := make([]byte, 4*1024)
	n, err1 := conn.Read(buf)
	if n == 0 {
		fmt.Println("conn.Read err =", err1)
		return
	}

	//打印响应报文格式
	fmt.Printf("##%v##", string(buf[:n]))

}
