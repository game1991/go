package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子，只需要一次
	rand.Seed(time.Now().UnixNano()) //以当前系统时间为种子
	for i := 0; i < 5; i++ {
		fmt.Println("rand= ", rand.Intn(100))
	}
}
