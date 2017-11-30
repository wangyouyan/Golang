package test

import(
	"testing"
	"code.oldboy.com/day1/threading2"
	"fmt"
)

func TestOdd(t *testing.T) {
	var odd string
	var result string
	odd = "测试奇数"
	result = threading.PrintOdd(odd)
	
	fmt.Printf("2017/11/30\n %s测试成功", result)
	t.Fatalf("Execute sucessful......")
	
}

func TestEven(t *testing.T) {
	var even int
	var result int
	even = 123456
	result = threading.PrintEven(even)
	
	fmt.Printf("\n\n--------------程序分割线-------------------\n\n")
	fmt.Printf("测试偶数%d, 这次测试是成功的，请放心!", result)
	t.Fatalf("Execute sucessful......")
	
}