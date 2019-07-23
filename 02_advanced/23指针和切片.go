package main

import "fmt"

// 切片指针作为函数参数
func slicePointerParams(p *[]int) {
	*p = append(*p, 6, 7, 8, 9)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	// 指针和切片建立联系
	p := &s
	fmt.Printf("%p\n", s)		// 0xc00001c090
	fmt.Printf("%p\n", p)		// 0xc00000c080
	fmt.Printf("%T\n", p)		// *[]int	切片指针类型


	// 通过指针操作切片元素
	(*p)[1] = 222
	fmt.Println(s)


	// 切片指针作为函数参数
	s1 := []int{1, 2, 3, 4, 5}
	// 引用传递
	slicePointerParams(&s1)
	fmt.Println(s1)		// [1 2 3 4 5 6 7 8 9]
}
