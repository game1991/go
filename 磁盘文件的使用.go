package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	//新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//使用完毕，关闭文件
	defer f.Close()
	var s string
	for i := 0; i < 10; i++ {
		//i=1\n，这个字符串存储到s中
		s = fmt.Sprintf("i=%d\n", i)
		//fmt.Println("s= ",s)
		n, err := f.WriteString(s)
		if err != nil {
			fmt.Println("err= ", err)
		}
		fmt.Println("n= ", n)

	}

}

func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//关闭文件
	defer f.Close()

	b := make([]byte, 1024*2) //2k大小
	//n代表从文件读取内容的长度
	n, err1 := f.Read(b)
	if err1 != nil && err1 != io.EOF { //文件出错，同时还没到结尾
		fmt.Println("err1= ", err1)
		return
	}
	fmt.Println(string(b[:n]))

}

//每次读取一行

func ReadFileline(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//关闭文件
	defer f.Close()
	//新建一个缓冲区，把内容先放缓冲区
	r := bufio.NewReader(f)
	for { //遇到'\n'就结束读取,但是会把'\n'也写进来
		n, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} //文件已经结束
			fmt.Println("err=", err)
		}
		fmt.Printf("n= #%s#\n", string(n))
	}

}

func main() {
	path := "./demo.txt"
	//WriteFile(path)
	//ReadFile(path)
	ReadFileline(path)
}
