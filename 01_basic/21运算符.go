package main

import "fmt"

func main() {
	// 位运算符
	a := 60
	b := 13
	var c int
	c = a&b
	fmt.Println(c)	// 12，按机器语言01来判断每一位上的值是否相等
}
