package main

import "fmt"

// func swap(a, b int) {
// 	a, b = b, a
// 	fmt.Printf("swap:a=%d,b=%d\n", a, b)
// }

// func main() {
// 	a, b := 10, 20
// 	swap(a, b)
// 	fmt.Printf("main:a=%d,b=%d\n", a, b)

// }    这是值传递

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}

func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Printf("main:a=%d,b=%d\n", a, b)

}
