package main

import (
	"fmt"
	"net/http"
)

func HandConn(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello world")) //w,给客户端回复数据
	//req,读取客户端发送的数据
	fmt.Println("r.Method=", req.Method)
	fmt.Println("r.URL=", req.URL)
	fmt.Println("r.Header=", req.Header)
	fmt.Println("r.Body=", req.Body)

}

func main() {
	//注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/mike/go", HandConn)

	//监听绑定
	http.ListenAndServe(":8000", nil)

}
