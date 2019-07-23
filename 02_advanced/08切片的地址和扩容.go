package main

import "fmt"

// 内存地址
func sliceAddress()  {
	// 空切片，指向内存地址中编号为 0 的空间，切片名本身就是一个地址
	var s [] int
	fmt.Printf("%p\n", s)		// 0x0

	// 当使用 append 进行追加数据时，切片地址可能会发生改变
	s = append(s, 1, 2, 3)
	fmt.Printf("%p\n", s)		// 0xc000094000

	s = append(s, 4, 5, 6)
	fmt.Printf("%p\n", s)		// 0xc000090000
}

// 切片扩容
func expansion() {
	s := make([] int, 5, 8)
	fmt.Println(s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))	// len = 5, cap = 8

	s = append(s, 1, 2, 3)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))	// len = 8, cap = 8

	s = append(s, 1, 2, 3)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))	// len = 11, cap = 16


	// 在使用 append 追加数据时，长度超过容量，容量自动扩容
	// 扩大的容量大小规律，一般时 容量 * 2，如果超过 1024 个字节，每次扩容上一次的 1/4
	// 容量扩容每次偶数
	s1 := make([] int, 0, 1)
	oldCap := cap(s1)
	for i := 0; i < 200000; i++ {
		s = append(s, i)
		newCap := cap(s)
		if oldCap < newCap {
			fmt.Printf("oldcap = %d, newcap = %d\n", oldCap, newCap)
			oldCap = newCap
		}
	}
}

func main() {
	// 内存地址
	sliceAddress()

	// 扩容
	expansion()
}
