package _1_basic

import "fmt"

func main() {
	var t complex64
	t = 2.4 + 3.14i
	fmt.Println(t)

	// 自动推导类型，默认为 complex128 类型
	t1 := 3.3 + 4.4i
	fmt.Printf("%T\n", t1)

	// 获取一个复数的内容
	// 获取实部 real(), 获取虚部 imag()
	cm := 3.1 + 1.4i
	fmt.Println(real(cm), imag(cm))
}
