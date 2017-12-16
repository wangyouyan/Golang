package main

import (
	"fmt"
)

func test1() {
	var a int
	a = 100
	fmt.Printf("%d\n",a)
	fmt.Printf("%p",&a)
}

func modify(a *int) {
	*a = 100
}

// func test1() {
// 	var a int = 1
// 	var p *int
// 	p = &b
// 	modify(p)
// 	fmt.Println(b)
// }

func test3() {
	///p 默认初始化nil
	var p *int
	var b int
	p = &b
	*p = 200 

	fmt.Printf("%p %p\n",p,&b,&p)
	p = new(int)
	*p = 1000
	fmt.Printf("%d\n", *p)
	fmt.Printf("%p %p %p\n", p, &b,&p)
}

func test5() {
	//定义一个切片类型
	var a []int
	a = make([]int,10)
	a[0] = 100
	fmt.Println(a)

	var p *[]int
	p = new([]int)
	(*p) = make([]int, 10)
	(*p)[0] = 100
	fmt.Println(p)
	p = &a
	(*p)[0] = 1000
	fmt.Println(a)
}

func main() {
	// test1()
	// test3()
	test5()
}