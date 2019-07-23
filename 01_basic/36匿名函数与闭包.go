package _1_basic

import "fmt"

func main() {
	var num = 10
	// 无参数匿名函数
	
	// 方式一
	f := func() {
		num ++
		fmt.Println(num)	// 11 可以直接访问 main 函数内的变量
	}
	f()	// 函数调用
	fmt.Println(num)	// 11

	// 方式二
	// 定义并直接执行匿名函数
	func() {
		num --
		fmt.Println(num)	// 10
	}()
	fmt.Println(num)	// 10
	
	// 方式三
	type Anonymous func()
	var f1 Anonymous
	f1 = f
	f1()
	fmt.Println(num)	// 11


	// 带参匿名函数
	func(a, b int) {
		var sum int
		sum = a + b
		fmt.Println(sum)
	}(9, 8)


	// 带返回值的匿名函数
	min, max := func(a, b int) (min, max int) {
		if a > b {
			max = a
			min = b
		} else {
			max = b
			min = a
		}
		return
	}(30, 20)
	fmt.Println(min, max)


	// 闭包
	closure := closure()
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
}

// 闭包函数案例
func closure() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
