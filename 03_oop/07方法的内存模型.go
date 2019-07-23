package main

import "fmt"

type student7 struct {
	name string
	age int
	score int
}

// 结构体方法，默认是值传递
func (s student7) valuePass() {
	s.score = -9
	fmt.Println(s)		// {loedan 19 -9}
}

// 结构体方法，引用传递
func (s *student7) referencePass() {
	s.score = -9
	fmt.Println(s)
}

func main() {
	s := student7{"loedan", 19, 199}
	fmt.Println(s)

	// 值传递
	s.valuePass()
	fmt.Println(s)		// {loedan 19 199}

	// 引用传递
	s.referencePass()
	fmt.Println(s)
}
