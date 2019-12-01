package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("usage: xxx srcFile dstFile")
		return
	}
	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件与目标文件名字不能相同")
		return
	}
	//只读方式打开源文件
	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1= ", err1)
		return
	}
	//创建目标文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2= ", err2)
		return
	}
	//操作完毕，关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心文件程序处理，把源文件拷贝到目标文件，长度为源文件长度
	b := make([]byte, 4*1024) //设置一个4k大小的缓存区
	for {
		n, err := sF.Read(b) //n代表文件读取的长度
		if err != nil {
			if err == io.EOF { //文件读取完毕
				break
			}
			fmt.Println("err= ", err)
		}
		//往目标文件写，读多少写多少
		dF.Write(b[:n])
	}

}
