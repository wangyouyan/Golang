package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Talk()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d *Dog) Eat() {
	fmt.Println(d.Name,"is eating")
}

func (d *Dog) Talk() {
	fmt.Println(d.Name, "is wawa!")
}

func (c *Cat) Eat() {
	fmt.Println(c.Name,"is eating")
}

func (c *Cat) Talk() {
	fmt.Println(c.Name,"is miaomiao!")
}

func Test() {
	var a Animal
	var d Dog
	d.Eat()

	a = &d
	a.Eat()
}

func TestOperator() {
	var animalList []Animal
	d := &Dog{
		Name: "小黄",
	}
	animalList = append(animalList, d)

	d1 := &Dog{
		Name: "旺财",
	}
	animalList = append(animalList, d1)

	c := &Cat{
		Name: "咪咪",
	}
	animalList = append(animalList, c)

	for _, v := range animalList {
		v.Eat()
		v.Talk()
	}
}

func main() {
	// Test()
	TestOperator()
}