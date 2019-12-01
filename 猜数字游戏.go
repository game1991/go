package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatNum(p *int) {

	//设置一个种子
	rand.Seed(time.Now().UnixNano())

	//确保是5位数
	var num int
	for {
		num = rand.Intn(100000)
		if num >= 10000 {
			break
		}
	}
	//fmt.Println("num= ", num)

	*p = num

}

func GetNum(s []int, num int) {

	//取万位
	s[0] = num / 10000

	//取千位
	s[1] = num % 10000 / 1000
	//取百位
	s[2] = num % 1000 / 100
	s[3] = num % 100 / 10
	s[4] = num % 10
}

func Game(randSlice []int) {
	var num int
	keySlice := make([]int, 5)

	for {
		for {
			fmt.Println("请输入一个五位数: ")
			fmt.Scan(&num)
			//9999 <num < 1000000
			if 9999 < num && num < 100000 {
				break
			}
			fmt.Println("输入的不是五位数，请重新输入")
		}

		//fmt.Println("num = ", num)

		GetNum(keySlice, num)
		//fmt.Println("keySlice = ", keySlice)
		n := 0
		for i := 0; i < 5; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了一些\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了一些\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了!!!\n", i+1)
				n++
			}

		}
		if n == 5 { //5位数字全部猜对
			fmt.Println("全部猜对!!!游戏结束")
			break //跳出循环
		}
	}
}

func main() {
	var randNum int

	//产生一个5位数
	CreatNum(&randNum)
	//fmt.Println("randNum= ", randNum)

	//保存这5个数字以切片方式提取出来
	randSlice := make([]int, 5)

	GetNum(randSlice, randNum)

	//fmt.Println("randSlice= ", randSlice)

	Game(randSlice) //进行游戏
}
