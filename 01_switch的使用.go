package main

import "fmt"

func main() {
	var num int
	fmt.Printf("请按下楼层：")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("按下的是1楼")
	case 2:
		fmt.Println("按下的是2楼")
	case 3:
		fmt.Println("按下的是3楼")
	case 4:
		fmt.Println("按下的是4楼")
	default:
		fmt.Println("按下的是其他楼层")

	}

	score := 85
	switch {
	case score > 90:
		fmt.Println("成绩优秀")
	case score > 80:
		fmt.Println("成绩良好")
	case score > 70:
		fmt.Println("成绩一般")
	default:
		fmt.Println("其他")
	}

}
