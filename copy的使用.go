package main

import "fmt"

func main() {
	srcSlice := []int{1, 2, 3, 4, 5}
	dstSlice := []int{666}

	copy(dstSlice, srcSlice)

	fmt.Println("s= ", dstSlice)

}
