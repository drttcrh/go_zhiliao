package main

import "fmt"

// 切片截取
func intercept() {
	// 方式一
	s := [] int {1, 2, 3, 4, 5}
	// 截取 s[起始下标:结束下标:max]
	slice := s[0:3:5]
	fmt.Println(slice)
	// len = 结束下标 - 起始下标
	fmt.Println(len(slice))
	// cap = max - 起始下标
	fmt.Println(cap(slice))


	// 方式二
	s1 := [] int {1, 2, 3, 4, 5}
	ss1 := s1[:]	// 等同于 ss1 := s1
	fmt.Println(ss1)
	fmt.Println(len(ss1))
	fmt.Println(cap(ss1))

	// 方式三
	s2 := [] int {1, 2, 3, 4, 5}
	ss2 := s2[0:4]	// 等同于 ss2 := s2[:4]
	fmt.Println(ss2)
	fmt.Println(len(ss2))
	fmt.Println(cap(ss2))

	// 方式四
	s3 := [] int {1, 2, 3, 4, 5}
	ss3 := s3[3:5]	// 等同于 ss3 := s3[3:5]
	fmt.Println(ss3)
	fmt.Println(len(ss3))
	fmt.Println(cap(ss3))
}

// 修改切片数据
func sliceChange() {
	s := [] int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5]
	fmt.Println(s1)		// [2 3 4]

	// 修改数据
	// 截取后的切片还是原来切片中的一块内容，修改截取后的切片，会影响原始切片
	s1[2] = 999
	fmt.Println(s1)		// [2 3 999]
	fmt.Println(s)		// [0 1 2 3 999 5 6 7 8 9]
}

func main() {
	// 切片截取
	intercept()

	// 修改切片数据
	sliceChange()
}
