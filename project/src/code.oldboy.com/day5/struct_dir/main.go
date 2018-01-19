package main

import (
	"fmt"
)

//定义一个类型
type Int int

//定义一个结构体类型
type Test struct {
	A int
	b int
}

//定义一个结构体类型
type Student struct {
	Age int
	Name string
	Sex string
	Grade string
	Score int
	t Int
	a Test
	c *int
}

func testStruct() {
	var s Student
	s.Name = "Rain"
	s.Age = 28
	s.Score = 90
	s.Sex = "Male"
	s.a.A = 111
	s.a.b = 222
	s.c = new(int)
	*(s.c) = 100

	fmt.Printf("name:%s age:%d score:%d sex:%s c=%v\n",s.Name,s.Age,s.Score,s.Sex,*(s.c))
	fmt.Printf("%+v\n",s)
	s1 := s
	s1.Name = "Alex"
	*(s1.c) = 200
	fmt.Printf("name:%s age:%d score:%d sex:%s c=%v\n",s.Name,s.Age,s.Score,s.Sex,*(s.c))
	fmt.Printf("%+v\n",s)

	var p1 *int = new(int)
	p2 := p1
	*p2 = 1234
	fmt.Printf("s1=%d\n",*p1)

	var p3 = new(Student)
	(*p3).Score = 99
	p4 := p3
	p4.Score = 1000
	fmt.Printf("p3=%+v\n",*p3)
}

func main() {
	testStruct()
}