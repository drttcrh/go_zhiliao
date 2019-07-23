package main

import "fmt"

// 切片定义
func sliceDefinition()  {
	// len(slice)	计算切片的长度
	// cap(slice)	计算切片的容量

	// 切片定义一
	var slice [] int	// 定义一个空切片
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// 像切片中添加元素
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))


	// 切片定义二，定义并初始化
	var s1 [] int = [] int {1, 2, 3}
	fmt.Println(s1)

	// 切片定义三，自动推导
	s2 := [] int {1, 2, 3, 4, 5}
	fmt.Println(s2)

	// 切片定义四，make 创建
	s3 := make([] int, 3, 5)	// 创建一个长度为 3，容量为 5 的切片
	fmt.Println(s3)		// [0 0 0]
	s3 = append(s3, 1, 2, 3)
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
}

// 赋值
func assignSlice() {
	// 下标赋值
	s := make([] int, 5)
	s[0] = 1
	s[1] = 2

	// 循环赋值
	s1 := make([] int, 5, 10)
	for i := range s1 {
		s1[i] = i
	}
	fmt.Println(s1)
}

func main() {
	// 切片定义
	sliceDefinition()

	// 赋值
	assignSlice()
}
