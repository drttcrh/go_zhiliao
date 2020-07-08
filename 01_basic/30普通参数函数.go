package main

import "fmt"

func fixedLengthParamsFunc(a, b int) {
	fmt.Printf("%d, %d\n", a, b)
}

func main() {
	fixedLengthParamsFunc(1, 2)
}
