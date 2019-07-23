package main

import "fmt"

type hero struct{
	name string
	age int
	power int
}

// 结构体 map 作为函数参数
func structMapParams(m map[int]hero) {
	for _, v := range m {
		fmt.Println(v.power)
	}

	//
}

func main() {
	// 将结构体作为 map 中的值
	m := make(map[int]hero, 3)
	m[1] = hero{"loedan", 18, 100}
	m[2] = hero{"harden", 11, 13}
	m[3] = hero{"macgrady", 32, 32}
	fmt.Println(m)		// map[1:{loedan 18 100} 2:{harden 11 13} 3:{macgrady 32 32}]

	// 删除 map 中的元素
	delete(m, 3)
	fmt.Println(m)		// map[1:{loedan 18 100} 2:{harden 11 13}]


	// 将结构体切片作为 map 的值
	m1 := make(map[int][]hero)
	m1[0] = []hero{
		{"loedan", 12, 1212},
		{"harden", 56, 123},
		{"macgrady", 23, 43},
	}
	m1[0] = append(m1[0], hero{"poul", 23, 234})
	fmt.Println(m1)		// map[0:[{loedan 12 1212} {harden 56 123} {macgrady 23 43} {poul 23 234}]]


	// 结构体 map 作为函数参数
	m2 := make(map[int]hero)
	m2[1] = hero{"loedan", 18, 100}
	m2[2] = hero{"harden", 11, 13}
	m2[3] = hero{"macgrady", 32, 32}
	structMapParams(m2)
}
