package main

import (
	"fmt"
)

func main() {
	// switch 判断
	// 1.可以同时对多个值做判断
	// 2.不同的case之间不需要break
	// 3.fallthrough ，判断穿透，出现该关键字后就不执行默认的 break 动作

	// 方式一
	var score int = 90
	switch score {
	case 90:
		fmt.Println("A")
		fallthrough
	case 80:
		fmt.Println("B")
	case 60, 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	// 方式二
	switch score1 := 90; score1 {
	case 90:
		fmt.Println("A")
		fallthrough
	case 80:
		fmt.Println("B")
	case 60, 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	// 方式三
	var score2 int = 90
	switch {	// switch 没有条件
	case score2 >= 90:	// case 判断条件
		fmt.Println("A")
	case score2 >= 80:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	// 方式四
	switch socre3 := 90; {	// 只有初始化语句，没有条件
	case socre3 >= 90:	// case 判断条件
		fmt.Println("A")
	case socre3 >= 80:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	// 总结
	/**
	优点：
	1.if 适合进行区间判断，嵌套使用
	2.switch 适合做固定值的判断，执行效率高，可以将多个满足条件的值放在一起

	缺点：
	1.if 执行效率低
	2.switch 不建议使用嵌套
	 */
}
