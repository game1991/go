package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子，只需要一次
	rand.Seed(time.Now().UnixNano()) //以当前系统时间为种子

	var a [10]int
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100) //100以内的随机数组
		fmt.Printf("%d ,", a[i])
	}

	fmt.Printf("\n")
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}

	}
	fmt.Printf("\n排序后:\n")
	for i := 0; i < n; i++ {

		fmt.Printf("%d ,", a[i])
	}
	fmt.Printf("\n")
}
