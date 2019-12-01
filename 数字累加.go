package main

import "fmt"

//实现1+2+3+....100

// func test() (sum int) {
// 	for i := 1; i <= 100; i++ {
// 		sum += i //sum=sum+i
// 	}
// 	return
// }

// func main() {
// 	var sum int
// 	sum = test()
// 	fmt.Println("sum= ", sum)
// }

func test(i int) int {
	if i == 1 {
		return 1
	}
	return i + test(i-1)
}

func test01(i int) int {
	if i == 100 {
		return i
	}
	return i + test01(i+1)
}

func main() {
	sum := test(100)
	sum1 := test01(1)
	fmt.Println("sum1= ", sum1)
	fmt.Println("sum= ", sum)
}
