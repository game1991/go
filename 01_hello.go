package main

import "fmt"

func main() {
	const (
		a          = iota
		b1, b2, b3 = iota, iota, iota
		b          = iota
	)
	fmt.Printf("a= %d\nb1= %d,b2= %d,b3= %d\nb= %d\n", a, b1, b2, b3, b)

	const (
		a1 = iota
		a2
		a3
	)
	fmt.Printf("a1= %d,a2= %d,a3=%d\n", a1, a2, a3)
}
