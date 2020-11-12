package main

import "fmt"

func main() {
	a := 10
	fmt.Print(a)	// 只打印，不换行
	fmt.Println("a = ", a)	// 打印并换行
	fmt.Printf("a = %T\r\n", a)	// 格式化打印
}
