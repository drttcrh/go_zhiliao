package main

import "fmt"

func main() {
	// byte字符类型，也是 unit8 类型的别名
	// 1.在 go 语言中，单引号表示单个字符，双引号表示字符串
	var a byte = 'a'
	fmt.Println(a)	// 打印 97
	// %c 是一个占位符，表示打印输出一个字符
	fmt.Printf("%c\n", a)	// 打印字符，而不是 ascii 码
	fmt.Printf("%c\n", 97)	// 用等值的 ascii 码也可以打印到对应的字符
	fmt.Printf("%T\n", a)	// uint

	var c byte = '0'
	fmt.Println(c)	// 打印字符 0 的 ascii 码为 48

	var b string = "abc"
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
