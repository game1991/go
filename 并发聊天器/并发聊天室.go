package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string //用于发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

//用于保存在线用户cliAddr====>Client
var online map[string]Client

var message = make(chan string)

func Manager() {
	//给map分配空间
	online = make(map[string]Client)

	for {
		msg := <-message //没有消息前，这里会阻塞
		//遍历map,给map每个成员都发送此消息
		for _, cli := range online {
			cli.C <- msg
		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { //给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (game string) {
	game = "[" + cli.Addr + "]" + cli.Name + ":" + msg
	return

}

func HandleConn(conn net.Conn) { //处理用户连接
	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()
	defer conn.Close()

	//创建一个新结构体,默认，用户名与网络地址一样
	cli := Client{make(chan string), cliAddr, cliAddr}
	//把结构体添加到map
	online[cliAddr] = cli //map的赋值

	//新开一个协程，专门给当前客户端发送信息
	go WriteMsgToClient(cli, conn)
	//广播某个在线
	//message<-"["+cli.Addr+"]"+cli.Name+":login"
	message <- MakeMsg(cli, "login")
	//提示，我是谁
	cli.C <- MakeMsg(cli, "I am here")

	Quit := make(chan bool) //对方是否主动退出
	Data := make(chan bool) //对方是否有数据

	//新建一个协程，接收用户发送过来的数据
	go func() {
		game := make([]byte, 2048)
		for {
			n, err := conn.Read(game)
			if n == 0 { //代表对方断开或者出问题
				Quit <- true
				fmt.Println("conn.Read err=", err)
				return
			}
			msg := string(game[:n-1])          //通过Windows nc测试，多一个换行'\n'字符
			if len(msg) == 3 && msg == "who" { //遍历map,给当前用户所有成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range online {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
				//重命名rename|allen
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				rename := strings.Split(msg, "|")[1] //新的名字
				cli.Name = rename
				online[cliAddr] = cli
				conn.Write([]byte("rename compelte!\n"))

			} else { //转发此内容
				message <- MakeMsg(cli, msg)
			}
			Data <- true
		}

	}()

	for {
		//通过select检测channel的流动
		select {
		case <-Quit:
			delete(online, cliAddr)           //当前用户从map移除
			message <- MakeMsg(cli, "logout") //广播谁下线了
			return
		case <-Data:

		case <-time.After(60 * time.Second): //超时操作
			delete(online, cliAddr)                       //当前用户从map移除
			message <- MakeMsg(cli, "time out leave out") //广播谁下线了
			return
		}
	}
}

func main() {

	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net Listen err =", err)
		return
	}
	defer listener.Close()

	//新开一个协程，转发消息，只要有消息来了，变量map，给map每个成员都发送次消息
	go Manager()

	//主协程，循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener Accept err =", err)
			continue
		}
		//新建协程,处理用户连接
		go HandleConn(conn)
	}

}
