package main

import "fmt"

func main() {

	a := 10
	fmt.Println("a= ", a)
	b, c := 20.5, 30
	fmt.Println("b= ", b, "c= ", c)
	i, j := 10, 20
	i, j = j, i
	fmt.Printf("i= %d,j= %d\n", i, j)
}
