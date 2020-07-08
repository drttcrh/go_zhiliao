package main

import "fmt"

func main() {
	score := 70
	// if 语法结构
	//if 条件表达式 {	// 做括号必须与 if 在同一行
	//	// coding...
	//}
	if score > 70 {
		fmt.Println("b+")
	} else {
		fmt.Println("c")
	}

	// 简写方式
	// if 支持一个初始化语句，初始化语句和判断条件用 ; 分割
	if score := 91; score > 90 {
		fmt.Println("a+")
	}
}
