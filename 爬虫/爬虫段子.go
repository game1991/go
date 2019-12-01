package main

import (
	"fmt"
	//"io"
	"net/http"
	//"os"
	"regexp"
	"strconv"
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
	//开始爬取网页内容
	result, err1 := GetHttp(url)
	if err1 != nil {
		//fmt.Println("GetHttp err= ",err1)
		err = err1
		return
	}

	//取关键信息，取标题，正则表达式方式
	re1 := regexp.MustCompile(`<div class="artical-importantPic">(?s:(.*?))">`)
	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}
	tmptitle := re1.FindAllStringSubmatch(result, -1)
	for _, data := range tmptitle {
		title = data[1]
		//strings.Replace(title, "\r\n", "", -1)
		break
	}

	re2 := regexp.MustCompile(`<div class="artical-main-content">(?s:(.*?))<span id="editor_baidu">`)
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}
	tmpcontent := re2.FindAllStringSubmatch(result, -1)
	for _, data := range tmpcontent {
		content = data[1]
		//strings.Replace(title, "\r\n", "", -1)
		break
	}
	return

}

// func StoreFile(i int, fileTitle, fileContent []string) {
// 	//新建文件
// 	f, err := os.Create(strconv.Itoa(i) + ".txt")
// 	if err != nil {
// 		fmt.Println("os.Create err=", err)
// 		return
// 	}
// 	defer f.Close()
// 	//写内容
// 	n := len(fileTitle)
// 	for i := 0; i < n; i++ {
// 		//写标题
// 		f.WriteString(fileTitle[i] + "\n")
// 		//写内容
// 		f.WriteString(fileContent[i] + "\n")

// 		f.WriteString("\n==========================\n")
// 	}

// }

func Spider(i int, page chan int) {
	//明确爬取的url
	url := "https://voice.hupu.com/nba/" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d个网页:%s\n", i, url)

	//开始爬取网页内容
	result, err := GetHttp(url)
	if err != nil {
		fmt.Println("GetHttp err=", err)
		return
	}
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
	//fmt.Printf("tmpWord=#%v#\n", tmpWord)
	//筛选出新的链接
	r1 := regexp.MustCompile(`<a href="(?s:(.*?))"`)
	if r1 == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	newUrl := r1.FindAllStringSubmatch(tmpWord, -1)
	//fmt.Printf("newUrl=#%v#\n", newUrl)

	// flieUrl := make([]string, 0)
	// fileTitle := make([]string, 0)
	// fileContent := make([]string, 0)

	for _, data := range newUrl {
		//开始爬取每个链接的网页，定义一个函数

		title, content, err := SpiderOnePage(data[1])
		if err != nil {
			fmt.Println("SpiderOnePage err", err)
			continue
		}
		fmt.Printf("content=#%v#", content)
		fmt.Printf("title=#%v#", title)
		// fileTitle = append(fileTitle, title)
		// fileContent = append(fileContent, content)

	}

	// fmt.Println("fileTitle=", fileTitle)
	// fmt.Println("fileContent=", fileContent)

	//把内容写入文件
	// StoreFile(i, fileTitle, fileContent)

}

func DoWork(start, end int) {
	fmt.Printf("准备第%d页到%d页的网址\n", start, end)
	page := make(chan int)

	for i := start; i <= end; i++ {
		//定义一个函数，爬主页面
		go Spider(i, page)
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
