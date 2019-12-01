package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func SpiderPage(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	//读取网页Body内容
	buf := make([]byte, 4*1024)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 { //读取结束或者出问题
			fmt.Println("resp.Body.Read err=", err)
			break
		}
		result += string(buf[:n])
		//fmt.Println("result= ", result)
	}
	return

}

func DoWork(start, end int) {
	fmt.Printf("正在爬取%d到%d的页面\n", start, end)

	//1)明确目标（要知道你在哪个范围或者网站去搜索）
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=nba&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		fmt.Println("url= ", url)
		//2)爬取（把网站内容全部爬取出来）
		result, err := SpiderPage(url)
		if err != nil {
			fmt.Println("SpiderPage err=", err)
			continue
		}
		//把内容写入文件
		flieName := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(flieName)
		if err1 != nil {
			fmt.Println("os.Create err=", err1)
			continue
		}
		f.WriteString(result)

		f.Close()

	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页（>=1）：")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>=起始页）：")
	fmt.Scan(&end)

	DoWork(start, end)

}
