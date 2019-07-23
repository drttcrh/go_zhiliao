package main

import "fmt"

type student struct {
	id int
	name string
	score int
}

// 用数组放置多个结构体
func structArray() {

	var arr [3]student = [3]student{
		student{1, "loedan", 100},
		student{2, "jordan", 100},
		student{3, "harden", 100},
	}
	fmt.Println(arr)	// [{1 loedan 100} {2 jordan 100} {3 harden 100}]

	// 循环获取每一个结构体信息
	for _, v := range arr {
		fmt.Println(v)
	}

	// 修改数组内结构体成员的信息
	arr[1].score = 97
	fmt.Println(arr[1])		// {2 jordan 97}
}

// 结构体数组作为函数参数
func structArrayParams(arr [3]student) {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j].score > arr[j + 1].score {
				arr[j + 1], arr[j] = arr[j], arr[j + 1]
			}
		}
	}
	fmt.Println(arr)	// [{2 jordan 32} {3 harden 74} {1 loedan 98}]
}

// 结构体切片作为函数参数
func struceSliceParams(s []student) {
	for i := 0; i < len(s) - 1; i++ {
		for j := 0; j < len(s) - i - 1; j++ {
			if s[j].score > s[j + 1].score {
				s[j + 1], s[j] = s[j], s[j + 1]
			}
		}
	}
	fmt.Println(s)	// [{4 kobe 31} {2 jordan 32} {3 harden 74} {1 loedan 98}]
}

func main() {
	// 用数组放置多个结构体
	structArray()

	// 结构体数组作为函数参数
	var arr [3]student = [3]student{
		student{1, "loedan", 98},
		student{2, "jordan", 32},
		student{3, "harden", 74},
	}
	structArrayParams(arr)
	fmt.Println(arr)	// [{1 loedan 98} {2 jordan 32} {3 harden 74}]	值传递


	// 结构体切片作为函数参数
	var s []student = []student{
		student{1, "loedan", 98},
		student{2, "jordan", 32},
		student{3, "harden", 74},
	}
	s = append(s, student{4, "kobe", 31})
	struceSliceParams(s)
	fmt.Println(s)	// [{4 kobe 31} {2 jordan 32} {3 harden 74} {1 loedan 98}]	引用传递
}
