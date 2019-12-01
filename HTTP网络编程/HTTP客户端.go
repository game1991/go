package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/mike/go") //获取网址数据
	if err != nil {
		fmt.Println("http.Get err=", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status= ", resp.Status)
	fmt.Println("Statuscode= ", resp.StatusCode)
	fmt.Println("Header= ", resp.Header)
	//fmt.Println("Body= ", resp.Body)

	buf := make([]byte, 1024*4)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err=", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("tmp=", tmp)
}
