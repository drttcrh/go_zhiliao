package main

import (
	"fmt"
	"strconv"
)

// 将字符串转成字符切片
func stringToSlice() {
	str := "hello world"
	// 强制类型转换
	slice := []byte(str)
	fmt.Println(slice)
}

// 字符切片转成字符串
func sliceToString() {
	slice := []byte{'a', 'b', 'c', 'd', 98}
	fmt.Println(slice)

	// 强制转换
	fmt.Println(string(slice))
}

// 将其他类型转换成字符串
func otherToString() {
	str := strconv.FormatBool(false)
	fmt.Println(str)
}

func main() {

	// 将字符串转成字符切片
	stringToSlice()

	// 字符切片转成字符串
	sliceToString()

	// 将其他类型转换成字符串
	otherToString()
}
