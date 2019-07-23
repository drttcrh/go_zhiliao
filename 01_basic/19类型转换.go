package _1_basic

import (
	"fmt"
)

func main() {

	a, b, c := 50, 20, 30

	sum := a + b + c
	// 类型转换
	// 1.数据类型（变量） int(a)
	// 2.数据类型（表达式）int(a + b)
	fmt.Println(float64(sum) / 3)

	// 3.不兼容的数据类型之间不支持类型转换
	//var flag bool
	//fmt.Println(int(flag))
	//flag = bool(1)

	// 4.在类型转换时，推荐将低类型转为高类型
	// 如：int8 --> int16 --> int32 --> int64
	//	   float32 --> float64
	//     int64 --> float64	整形与浮点型转换时，一般整形转换成浮点型
	var d float32 = 3.1
	var e float64 = 3.4
	fmt.Println(float64(d) + e)

	// 5.高类型 转换 低类型，会丢失精度，数据溢出
	var g int = 1234
	fmt.Println(int8(g))
}
