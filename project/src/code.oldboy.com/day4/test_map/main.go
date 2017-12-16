package main

import (
	"fmt"
)

func testMap() {
	var a map[string]int
	//使用make分配内存
	a = make(map[string]int, 100)
	//为map的a插入key为age,value为28
	a["age"] = 28
	a["weight"] = 180
	fmt.Println(a)
	//遍历map的key,value
	for k, v := range a {
		fmt.Printf("a[%s] = %d\n",k,v)
	}
	//查找key是否存在
	//方法一:
	val, exist := a["abc"]
	fmt.Printf("val=%d ok=%t\n",val,exist)
	if exist {
		fmt.Printf("val = %d\n",val)
	}else {
		fmt.Println("not found")
	}
	//为a增加key为aaaaa的map,并且赋值给val
	//方法二
	val = a["aaaaa"]
	fmt.Println(val)
}

func modify(a map[string]int) {
	a["one"] = 1234
}

// func testMap2() {
// 	var b map[string]int
// 	b = make(map[string]int,8)

// }

//map切片,the slice for map
func testMapSlice() {
	s := make([]map[string]int,10)
	for i := 0;i < len(s); i++{
		s[i] = make(map[string]int,100)
	}
	s[0]["abc"] = 100
	s[5]["abc"] =100
	fmt.Println(s)
	fmt.Println("\n")
}

func main() {
	// testMap()
	testMapSlice()
}