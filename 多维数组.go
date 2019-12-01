package main

import "fmt"

func main() {
	//有多少个[]就是多少维
	//有多少个[]就有多少个循环
	a := [3][4]int{}
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d]=%d,", i, j, k)
		}
		fmt.Println("\n")
	}
	fmt.Println("a= ", a)

}
