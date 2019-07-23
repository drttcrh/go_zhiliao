package _1_basic

import "fmt"

func main() {
	// 一、语法
	//for 初始化条件; 判断条件; 条件变化 {
	//	// coding
	//}

	// 二、 计算 1 + 2 + 3 + ... + 10 的和
	sum := 0
	for i := 1; i < 11; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
	
	// 三、range
	// 通过 for 循环打印每一个字符
	str := "loedan"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])	// 获取单个字符的文本需要格式化输出，否则输出的是 ascii 码
	}
	// 通过 range 迭代打印每一个字符
	// range 默认返回两个值，值一为元素位置（下标），值二为内容
	for i, data := range str {
		fmt.Println(i, data)
	}
	// 丢弃内容
	for i, _ := range str {
		fmt.Println(i)
	}
	// 丢弃下标
	for _, data := range str {
		fmt.Println(data)
	}

	// 只有一个判断条件
	arr1 := [10] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 0 // 最小下标
	j := len(arr1) - 1	// 最大下标
	for i < j {
		arr1[i], arr1[j] = arr1[j], arr1[i]
		i++
		j--
	}

	// 没有任何表达式判断
	for {
		if i >=  j {
			break
		}
		arr1[i], arr1[j] = arr1[j], arr1[i]
		i++
		j--
	}
	fmt.Println(arr1)
}
