package _1_basic

import "fmt"

func main() {
	// 简写方式
	if score := 91; score > 90 {
		fmt.Println("a+")
	} else {
		fmt.Println("a-")
	}

	// 普通方式
	a := 8
	if a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("a < 10")
	}
}
