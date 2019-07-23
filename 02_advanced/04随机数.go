package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 双色球案例
// 要求：红色区 6 个球，数字不可重复，红色球号码区间 1 - 33
// 		蓝色区 1 个球，数字可与红色区重复，蓝色球号码区间 1 - 16
func doubleColorBall() {
	// 创建随机数种子
	rand.Seed(time.Now().UnixNano())

	var red [6] int
	for i := 0; i < len(red); i++ {
		v := rand.Intn(33) + 1

		// 重复判断
		for j := 0; j < i; j++ {
			if v == red[j] {
				v = rand.Intn(33) + 1	// 如果出现重复，重新生成随机数
				j = -1	// 让内层循环重新开始
			}
		}
		red[i] = v
	}
	fmt.Println("红色球：", red, "蓝色球：", rand.Intn(16) + 1)
}

func main() {
	// 1、导入头文件	math/rand
	// 2、随机数种子
	// 3、创建随机数

	// 创建随机数种子
	rand.Seed(1)
	fmt.Println(rand.Int())	// 5577006791947779410

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())	// 4229865887998040371
	fmt.Println(rand.Intn(10))	// 生成 10 以内的随机数(0 -- 9)，该方法比较常用


	// 双色球案例
	doubleColorBall()
}
