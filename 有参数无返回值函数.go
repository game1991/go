package main

import "fmt"

func Myfunc(a int, b int) {

	fmt.Printf("a=%d,b=%d\n", a, b)
}

func main() {
	Myfunc(666, 777)
}
