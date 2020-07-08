package main

import "fmt"

func sum(a, b int) {
	sum := a + b
	fmt.Println(sum)
}

// 不定长参数
// 不定长参数必须为参数列表的最后一个
// 不定长参数可以不传参
func indefiniteLengthParamsFunc(args ... int)  {
	fmt.Println(args)
	sum := 0
	// 循环计算每个参数相加的总和
	for _, data := range args {
		sum += data
	}
	fmt.Println(sum)
}

func main() {
	sum(1, 2)
	indefiniteLengthParamsFunc(1, 2, 3)
}
