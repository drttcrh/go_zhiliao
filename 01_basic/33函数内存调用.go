package _1_basic

import "fmt"

func ram(a, b int) int {
	return a + b
}

func main() {
	// 栈区内存
	// 内存相对于一个可执行程序来说，分为四个区
	// 代码区、数据区、堆区、栈区
	/**
	1、代码区
	2、数据区
	3、堆区
	4、栈区：用来存储程序执行过程中函数内部定义的信息和局部变量
	 */
	result := ram(1, 2)
	fmt.Println(result)
}
