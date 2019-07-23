package _1_basic

import "fmt"

func main() {
	// break 	跳出当前整个 for 循环，break 在 for 和 switch 中都可使用
	// continue	跳出本次循环，for 循环会像后继续
	// goto		跳转

	i := 0
	for {	// for 后面不写任何条件判断，那就是一个死循环
		i++
		fmt.Println(i)

		//continue	//

		if i == 5 {
			break	// 跳出当前循环
		}
	}

	fmt.Println("a")
	goto End	// 跳转到 End 这一行
	fmt.Println("b")
	End:
	fmt.Println("c")
}
