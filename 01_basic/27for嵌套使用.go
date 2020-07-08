package main

import "fmt"

func main() {
	// for 循环嵌套
	count := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			//fmt.Println(i, j)
			count++
		}
	}
	fmt.Println(count)	// 25

	// 练习
	// 百钱白鸡
	// 中国古代数学家张丘建在他的《算经》中提出了一个著名的"百钱百鸡"问题：
	// 一只公鸡值五钱，一只母鸡值三钱，三只小鸡值一钱
	// 现在要用百钱买百鸡，请问公鸡，母鸡，小鸡各多少只？
	// cock 🐓，hen 母鸡， chicken 🐤
	// 答案一
	count1 := 0
	for cock := 0; cock < 20; cock++ {
		for hen := 0; hen < 33; hen++ {
			for chicken := 0; chicken < 100; chicken += 3 {
				count1++
				if cock + hen + chicken == 100 && cock * 5 + hen *3 + chicken / 3 == 100 {
					fmt.Printf("🐓 -- %d, 母鸡 -- %d, 🐤 -- %d\n", cock, hen, chicken)
				}
			}
		}
	}
	fmt.Println(count1)	// 循环了 22440 次
	// 答案二，推荐使用，效率更高
	count2 := 0
	for cock := 0; cock < 20; cock++ {
		for hen := 0; hen < 33; hen++ {
			count2++
			chicken := 100 - cock - hen
			if chicken % 3 == 0 && cock * 5 + hen * 3 + chicken / 3 == 100 {
				fmt.Printf("🐓 -- %d, 母鸡 -- %d, 🐤 -- %d\n", cock, hen, chicken)
			}
		}
	}
	fmt.Println(count2)	// 循环了 660 次

}
