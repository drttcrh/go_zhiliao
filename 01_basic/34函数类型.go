package _1_basic

import "fmt"

func funcType(a, b int) {
	fmt.Println(a + b)
}

// 定义一个函数类型的别名
type FUNCDEMO func(int, int)	// 定义一个函数类型的类型别名

func main() {
	// 函数名表示一个地址，函数在内存中代码区的地址
	funcType(1, 2)
	fmt.Println(funcType)	// 0x1092ee0，打印的是函数在内存中的地址

	a := 1
	fmt.Println(&a)			// 0xc00008a000 打印变量在内存中的地址

	f := funcType
	f(3, 4)	// 等同于 funcType(3, 4)

	// 获取 f 的类型
	fmt.Printf("%T\n", f)	// func(int, int)

	// 定义一个函数类型的变量
	var fu func(int, int)
	fmt.Println(fu)

	/**
	重点，对函数类型的使用
	 */
	var fu1 FUNCDEMO	// 定义变量为函数类型的一个变量
	fmt.Println(fu1)
	fu1 = funcType		// 让变量实现具体的函数
	fu1(1, 3)			// 执行函数

	//
	fu2 := funcType
	fu2(9, 9)
}
