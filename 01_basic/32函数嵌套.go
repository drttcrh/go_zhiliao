package _1_basic

import "fmt"

// 第二层
func second(c, d int) {
	fmt.Println(c + d)
}

// 第一层
func first(a, b int) {
	second(a, b)
}

func first2(args ...int)  {

}

func first1(args ... int) {
	// 在不定参函数中调用另一个不定参函数的传参方式
	// [0:4] 表示 args 下标大于等于 0 小于 4 的值
	first2(args[0], args[1], args[2], args[3])	// 传递部分参数
	first2(args[0:4]...)	// 传递部分参数
	first2(args[:]...)	// 传递全部参数
	first2(args...)		// 传递全部参数
}

func main() {
	// 1、什么是函数嵌套
	// 在一个函数中调用另一个函数
	
	// 2、函数嵌套的执行过程
	first(1, 2)

	// 3、函数嵌套调用的应用
	// ./同名练习文件

	// 4、不定参函数的调用
	first1(1, 2, 3, 4)
}
