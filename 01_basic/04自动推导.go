package main

import "fmt"

func main() {
	// 自动推导
	var a = 10
	fmt.Println(a)
	fmt.Printf("%T\r\n", a)

	// 常用的自动推导
	// 自动推导 := 左边变量名
	b := 10
	b = 20
	b = 30
	fmt.Println(b)

	c := 3.14
	fmt.Println(c)
	fmt.Printf("%T\r\n", c)
}
