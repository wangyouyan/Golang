package main

import (
	"fmt"
)

//定义一个名为add的函数
func add(a, b int) int {
	return a + b
}

func main() {
	//函数是一等公民，也可以和变量一样进行赋值
	c := add   //c是一个函数类型的变量
	//打印函数类型,%p为指针地址，%T为变量类型
	fmt.Printf("%p %T %p %T\n",c,add,c,add)
	sum := c(10,20)
	fmt.Println(sum)

	sum = add(10,20)
	fmt.Println(sum)
}