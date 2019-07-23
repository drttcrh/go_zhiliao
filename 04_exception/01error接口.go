package main

import "fmt"

// 除法函数
func divide(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		fmt.Println("err=", err)
	} else {
		result = a / b
	}
	return
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
