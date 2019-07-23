package main

import "fmt"

// 指针作为函数参数
func pointerParams(a, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

func main() {
	// &变量：取地址操作，引用运算符
	// *指针变量：取值操作，解引用运算符
	a, b := 10, 20
	pointerParams(&a, &b)
	fmt.Println(a, b)

}
