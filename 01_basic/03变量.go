package _1_basic

import "fmt"

func main() {

	// 一、变量的声明
	// 1.声明 var 变量名 类型	// 变量声明之后必须使用
	// 2.只是声明变量，默认值为0
	// 3。在同一个 {} 里，声明的变量是唯一的

	var a int
	fmt.Println(a)

	var b, c int
	fmt.Println(b, c)

	// 二、变量的初始化，声明变量时，同时赋值

	var d float64 = 3.14	// 声明时直接初始化变量的值
	fmt.Println(d)

	var value float64 = 2
	var sum float64 = value + value
	fmt.Println(sum)
}
