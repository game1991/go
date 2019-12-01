package main

import "fmt"

// import "time"

func main() {
	// sum := 0
	// for i := 1; i <= 100; i++ {
	// 	sum = sum + i
	// }
	// fmt.Println("sum= ", sum)
	sum := [100]int{}

	for i, n := range sum {
		fmt.Println(i, n)
	}

}
