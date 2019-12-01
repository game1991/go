package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(4) //四核运行程序
	fmt.Println("n= ", n)

	for {
		go fmt.Print(1)
		fmt.Print(0)

	}

}
