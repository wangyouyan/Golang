package main

import (
	"fmt"
)

type Int int 

func Add(a, b int) int{
	return a + b
}

func (i *Int)Add(a,b int) {
	*i = Int(a + b)
	return
}

func (i Int)Sub(a,b int) {
	i = Int(a - b)
	return
}

func testInt() {
	var a Int
	a.Add(100,200)
	fmt.Println(a)
	a.Sub(100,200)
	fmt.Println(a)
}

type Student struct {
	Name string
	Age int
}

func (s *Student) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

func main() {
	testInt()
	var s Student
	s.Set("abcd",100)
	fmt.Println(s)
}