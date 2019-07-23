package _1_basic

import "fmt"

// 递归函数，调用函数自己本身
// 递归函数相同的结构
// 1、跳出条件
func recursive(a int) {
	if a == 1 {
		fmt.Println(a)
		return	// 函数结束
	}
	recursive(a-1)
	fmt.Println(a)
}

// 用递归的方式计算前一百个数相加的总和
func hundredSum(num int) int {
	if num == 1 {
		return num
	}
	return num + hundredSum(num - 1)
}

func main() {
	recursive(3)
	fmt.Println(hundredSum(100))	// 5050
}
