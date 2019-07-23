package main

import "fmt"

func testA() {
	fmt.Println("testA")
}

func testB(x int) {
	// 设置 recover()
	// recover() 必须在 defer 调用的函数中使用
	defer func() {
		// 防止程序崩溃
		//fmt.Println(recover())
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var a [3]int
	a[x] = 999
}

func testC() {
	fmt.Println("testC")
}

func main() {
	testA()
	testB(3)
	testC()
}
