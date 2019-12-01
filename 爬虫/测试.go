package main

import (
	"fmt"
	//"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetHttp(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	//读取网页内容
	game := make([]byte, 1024*4)
	for {
		n, _ := resp.Body.Read(game)
		if n == 0 { //读取结束或者出问题
			//fmt.Println("resp.Body.Read err=", err)
			break
		}
		result += string(game[:n]) //累加读取的内容
	}
	return

}

func SpiderOnePage(url string) (title, content string, err error) {
	//读取网页内容
	result, err1 := GetHttp(url)
	if err1 != nil {
		err = err1
		return
	}
	//正则表达式，筛选标题
	re1 := regexp.MustCompile(`<div class="artical-importantPic">(?s:(.*?))">`)
	if re1 == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	tmpTitle := re1.FindAllStringSubmatch(result, -1)
	var mtitle string
	for _, data := range tmpTitle {
		mtitle += string(data[1]) + "a>"
		break
	}

	//筛选标题真实内容
	reg1 := regexp.MustCompile(`alt="(?s:(.*?))a>`)
	if reg1 == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	stitle := reg1.FindAllStringSubmatch(mtitle, -1)
	for _, data := range stitle {
		title = strings.Replace(data[0], "alt=\"", "", -1)
		title = strings.Replace(title, "a>", "", -1)
		break
	}
	//筛选文章内容

	re2 := regexp.MustCompile(`<div class="artical-main-content">(?s:(.*?))<span id="editor_baidu">`)
	if re2 == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	tmpContent := re2.FindAllStringSubmatch(result, -1)
	var mcontent string
	for _, data := range tmpContent {
		mcontent += string(data[1])
		content = strings.Replace(mcontent, " ", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "<p>", "", -1)
		content = strings.Replace(content, "</p>", "", -1)
		content = strings.Replace(content, "&rdquo;", "\"", -1)
		content = strings.Replace(content, "&ldquo;", "\n", -1)
		content = strings.Replace(content, "&nbsp", " ", -1)

		break
	}
	return
}

func Spider(i int, page chan int) {
	//明确爬取的url
	url := "https://voice.hupu.com/nba/" + strconv.Itoa(i)
	fmt.Printf("正在爬取第%d个网页:%s\n", i, url)

	//开始爬取网页内容
	result, err := GetHttp(url)
	if err != nil {
		fmt.Println("GetHttp err=", err)
		return
	}
	//fmt.Println("r=", result)
	//取网页内容,<h4></h4>
	r := regexp.MustCompile(`<h4>(?s:(.*?))</h4>`)
	if r == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	mainWord := r.FindAllStringSubmatch(result, -1)
	//fmt.Println("mainWord=", mainWord)
	var tmpWord string
	for _, data := range mainWord {
		tmpWord += string(data[1])
	}
	//筛选出新的链接
	r1 := regexp.MustCompile(`<a href="(?s:(.*?))"`)
	if r1 == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	newUrl := r1.FindAllStringSubmatch(tmpWord, -1)

	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)

	for _, data := range newUrl {
		title, content, err := SpiderOnePage(data[1])
		if err != nil {
			fmt.Println("SpiderOnePage err=", err)
			continue
		}
		// fmt.Println("title=\n", title)
		// fmt.Println("content=\n", content)
		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)
	}
	//把内容写入文件
	StoreFile(i, fileTitle, fileContent)
	page <- i //写第几页爬完

}

func StoreFile(i int, fileTitle, fileContent []string) {

	//新建文件
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("os.Create err=", err)
		return
	}
	defer f.Close()
	//写内容
	n := len(fileTitle)
	for i := 0; i < n; i++ {
		//写标题
		f.WriteString(fileTitle[i] + "\n")
		//写内容
		f.WriteString(fileContent[i] + "\n")

		f.WriteString("\n==========================\n")
	}

}

func DoWork(start, end int) {
	fmt.Printf("准备第%d页到%d页的网址\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		//定义一个函数，爬主页面
		go Spider(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完毕\n", <-page)
	}

}

func main() {
	var start, end int
	fmt.Println("请输入起始页>=1:")
	fmt.Scan(&start)
	fmt.Println("请输入终止页（>=起始页）:")
	fmt.Scan(&end)

	DoWork(start, end)

}
