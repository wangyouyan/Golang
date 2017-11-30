package main

import (
	"code.oldboy.com/day1/calc"
	"fmt"
)

func main() {
	var sum int 
	var sub int
	var sum1 int
	var sub1 int 
	sum = calc.Add(11, 22)
	sub = calc.Sub(20, 10)
	sum1, sub1 = calc.Calc(15, 10)

	fmt.Printf("调用加法函数运算结果为:%d \n", sum)
	fmt.Printf("调用减法函数结果为: %d \n", sub)
	fmt.Println("稍作休息, 咱们来调用一次加减法综合运算函数!!!")
	fmt.Println(sum1, sub1)
}