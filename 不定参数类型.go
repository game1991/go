package main

import "fmt"

func Myfunc(args ...int) {
	fmt.Println("len(args)= ", len(args))
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d]=%d\n", i, args[i])
	}
	fmt.Println("==========================")

	for i, data := range args {
		fmt.Printf("args[%d]=%d\n", i, data)
	}

}

func main() {
	Myfunc()
	fmt.Println("+++++++++++++++++++++")
	Myfunc(1)
	fmt.Println("+++++++++++++++++++++")
	Myfunc(1, 2, 3)

}
