package _1_basic

import "fmt"

func fixedLengthParamsFunc(a, b int) {
	fmt.Printf("%d, %d\n", a, b)
}

func main() {
	fixedLengthParamsFunc(1, 2)
}
