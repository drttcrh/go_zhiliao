package main

import "fmt"

func main() {

	var a float64
	// 进程阻塞，等待用不输入
	//fmt.Scanf("%f", &a)	// 方式一
	//fmt.Println(a)

	fmt.Println("请输入：")
	fmt.Scan(&a)	// 方式二，推荐使用
	fmt.Println(a)


}
