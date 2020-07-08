package main

import "fmt"

func scope() {
	a = 5
	a += 1
}

// 全局变量，定义在函数外部
// 在整个文件可用
// 全局变量在整个项目目录中变量名唯一不可重复
var a int

func main() {
	// 局部变量，定义在函数内部的变量
	// 作用域：局部变量只在函数内部有效
	// 不同的函数可以定义相同的变量名，互不影响
	a = 10
	scope()
	fmt.Println(a)
}
