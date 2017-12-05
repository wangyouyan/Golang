package main

import(
	"fmt"
)

const (
	a = iota
	b
	c
	d = 10 /*iota初始值为0， 每递增1行，变量自动加1*/
	e1 = iota /* 定义不生效*/
	e2 = iota /* 定义不生效*/
	e3 = iota /* 定义不生效*/
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)	

}