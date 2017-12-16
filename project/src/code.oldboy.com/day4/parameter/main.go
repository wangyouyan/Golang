package main

import (
	"fmt"
)

type add_func func(int,int) int

func add(a,b int) int {
	return a + b
}

func operator(op add_func, a int, b int) int {
	//使用传进来的函数，进行操作
	return op(a,b)
}

func main() {
	// c := add
	// fmt.Println(c)
	sum := operator(add, 100,200)
	fmt.Println(sum)
}