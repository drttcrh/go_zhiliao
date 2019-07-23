package main

import "fmt"

func test(x int) {
	v := 100 / x
	fmt.Println(v)
}

func main() {
	// defer 语句只能出现在函数内部
	//defer fmt.Println("hello")
	//fmt.Println("loedan")
	// loedan
	// hello

	// 一个函数有多个 defer 语句，他们的执行顺序是先进后出
	//defer fmt.Println("hello")
	//defer fmt.Println("world")
	//defer fmt.Println("i m")
	//defer fmt.Println("golang")
	// golang
	// i m
	// world
	// hello


	// 即使函数或者某个延迟调用发生错误，这些调用依旧被执行
	//defer fmt.Println("hello")
	//defer fmt.Println("world")
	//defer test(0)
	//defer fmt.Println("golang")
	// golang
	// world
	// hello
	// panic: runtime error: integer divide by zero

	a := 10
	b := 20
	defer func(a, b int) {
		fmt.Println("匿名函数a：", a)
		fmt.Println("匿名函数b：", b)
	}(a, b)

	a = 100
	b = 200
	fmt.Println("main函数a：", a)
	fmt.Println("main函数b：", b)
}
