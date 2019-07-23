package main

import "fmt"

// 切片作为函数参数
func sliceParams(s []int) {
	fmt.Println(s)
	if len(s) > 0 {
		s[0] = 9999
		fmt.Println(s)
	}
}

// 切片作为函数返回值
func returnSlice(s [] int) [] int {
	// 如果函数中使用 append 进行数据添加，行参地址有可能发生改变，就会指向实参的地址
	s = append(s, 6, 7, 8, 9, 10)
	//fmt.Println(s)
	return s
}

func main() {
	s := [] int {1, 2, 3, 4, 5}
	fmt.Println(s)

	// 引用传递
	sliceParams(s)
	fmt.Println(s)

	// 切片返回值
	s1 := [] int {1, 2, 3, 4, 5}
	s1 = returnSlice(s1)
	fmt.Println(s1)
}
