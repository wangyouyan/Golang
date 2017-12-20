package main

import(
	"fmt"
	"math/rand"
)

//数组学习
func array_test1() {
	var a [10]int //默认是64位操作系统,占用8个字节
	// var a [10]int32 占用4个字节 
	// var a [10]int8  占用1个字节
	for i := 0; i < 10; i++ {
		fmt.Printf("%p\n", &a[i]) //&符号可以取内存地址
	}
	// 使用for....range遍历数组
	// _代表可以忽略的值
	for index,_ := range a {
		fmt.Printf("a[%d]=%d\n", index, a[index])
	}
}

func array_test2() {
	var a [5]int =[5]int {1,2,3,4,5}
    var b [5]int
	b = a
	fmt.Printf("b=%v", b)
	b[0] = 200
	fmt.Printf("b=%v", b)
	fmt.Printf("a=%v", a)	
}

func test3() {
	//右边需要申明数组长度及类型
	var a [5] int = [5]int{1,2,3,4,5}
	fmt.Printf("%v\n", a)
	// 无需在左边定义长度及类型的时候，需要用....
	var b = [...]int{1,2,3,4,5,6,7,8} 
	fmt.Printf("%v\n", b)
	// 当未在左边定义长度及类型
	var c = [5]int{1,2,3}
	fmt.Printf("%v\n", c)
	// 字符串类型格式化,下边长度5表示下标长度,最多可以定义5个字符串，下标分别为0,1,2,3,4
	var d [5]string = [5]string{1:"rain",2:"alex",3:"tom",4:"danierjajajajja"}
	fmt.Printf("%+v\n", d)
	
}
//二维数组
func test4() {
	var a [4][2]int
	for i := 0; i < 4;i++ {
		for j := 0; j < 2;j++ {
			a[i][j] = (i+1)*(j+1)
		} 
		for i := 0; i < 4;i++ {
			for j := 0; j < 2;j++ {
				fmt.Printf("%d\n",a[i][j])
			} 
	}
}
}

// 给数组赋值100个随机数
func generateRandoms() {
	// 定义一个数组
	var a [100]int
	// for i := 0;i<100;i++  当没有定义i的时候,需要使用:= 
	for i:=0;i<len(a);i++{ //len(a)可以动态获取数组a的长度
		//赋值
		a[i] = rand.Int()
	}
	for i:=0;i<len(a);i++ {
		//取下标为i的元素的值
		fmt.Printf("%d\n", a[i])
	}
}

// 给数组赋值100个随机字符串
func GetRandomString() {
	var a [100]string
	// 以下字符串默认为utf-8
	var b string = "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(a); i++ {
		var str string
		for j := 0; j < 32;j++ {
			index := rand.Intn(len(b))
			str = fmt.Sprintf("%s%c", str, b[index])
		}
		a[i] = str
		fmt.Printf("a[%d]=%s\n", i, a[i])
	}	
 }

 func arr(s []int) {
	 length := len(s)
	//  for i := 0;i <= length;i ++ {
	// 	 return
	//  }
	 fmt.Print(s[length+1])
 }

func main() {
	// array_test1()
	// array_test2()
	// test3()
	// test4()
	// generateRandoms()
	// GetRandomString()
	test_data := []int{8,2,4,1,7,5,3,9,0,6}
	arr(test_data)
}