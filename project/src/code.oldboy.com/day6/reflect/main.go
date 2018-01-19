package main

import (
	"fmt"
	"reflect"
)

/* 反射 */

// 定义一个student结构体类型
type Student struct {
	Name string
	Age int
	Sex int64
}

// 定义一个函数，给对象赋值
func (s *Student) Set(name string, Age int, Sex int64) {
	s.Name = name
	s.Age = Age
	s.Sex = Sex
}

// 获取对象属性
func (s *Student) GetName(name string) {
	s.Name = name
}

func testStruct () {
	var stu *Student = &Student{}
		stu.Set("Jim", 18, 1)

	valueInfo := reflect.ValueOf(stu)
	fieldNum := valueInfo.Elem().NumField()
	fmt.Println("field name:", fieldNum)
	sexValueInfo := valueInfo.Elem().FieldByName("Jim")
	fmt.Println("sex=", sexValueInfo.Int())

	sexValueInfo.SetInt(100)
	fmt.Println(stu)

	setMethod := valueInfo.MethodByName("Set")
	fmt.Println(setMethod)

	var params []reflect.Value
	name := "Tom"
	age := 1000
	sex := 3883

	params = append(params, reflect.ValueOf(name))
	params = append(params, reflect.ValueOf(age))
	params = append(params, reflect.ValueOf(sex))

	setMethod.Call(params)
	fmt.Println(stu)
}

func main() {
	testStruct()
}

