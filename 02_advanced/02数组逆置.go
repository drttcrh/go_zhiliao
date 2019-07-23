package main

import "fmt"

// 数组逆置
func reverse() {
	// 数组逆置
	arr1 := [10] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 方式一
	for i := range arr1 {
		if i < len(arr1) - (i + 1) {
			arr1[i], arr1[len(arr1) - (i + 1)] = arr1[len(arr1) - (i + 1)], arr1[i]
		}
	}
	fmt.Println(arr1)
	// 方式二
	for i := 0; i < len(arr1); i++ {
		if i < len(arr1) - (i + 1) {
			arr1[i], arr1[len(arr1) - (i + 1)] = arr1[len(arr1) - (i + 1)], arr1[i]
		}
	}
	fmt.Println(arr1)
	// 方式三
	i := 0 // 最小下标
	j := len(arr1) - 1	// 最大下标
	for i < j {
		arr1[i], arr1[j] = arr1[j], arr1[i]
		i++
		j--
	}
	// 方式四
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

func main() {
	// 练习：从以下一个整数数组中取出最大的整数，最小的整数，总和，平均值
	// arr := [6] int {0, 1, 2, 3, 4, 5}

	arr := [6] int {0, 1, 2, 3, 4, 5}
	max := arr[0]
	min := arr[0]
	var sum int
	var avg float64

	// 最大值，最小值，总和
	for _, v := range arr{
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
		sum += v
	}
	fmt.Println("最大值：", max)
	fmt.Println("最小值：", min)
	fmt.Println("总和为：", sum)

	// 平均值
	avg = float64(sum) / float64(len(arr))
	fmt.Println("平均值为：", avg)


	//
}
