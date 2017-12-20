package main

import (
	"fmt"
)
/* 插入排序的基本思想是，经过i-1遍处理后,L[1..i-1]己排好序。
第i遍处理仅将L[i]插入L[1..i-1]的适当位置，使得L[1..i]又是排好序的序列。
要达到这个目的，我们可以用顺序比较的方法。首先比较L[i]和L[i-1]，
如果L[i-1]≤ L[i]，则L[1..i]已排好序，第i遍处理就结束了；
否则交换L[i]与L[i-1]的位置，继续比较L[i-1]和L[i-2]，
直到找到某一个位置j(1≤j≤i-1)，使得L[j] ≤L[j+1]时为止。
*/

//该函数主要是核心逻辑处理
func InsertionSort(slice_data []int) {
	//取切片的长度
	n := len(slice_data)
	//当n小于2的时候，没法排序了
	if n < 2 {
		return
	}
	//遍历切片的长度,从而进行切片元素两两比较
	for i := 1;i < n;i++ {
		for j := i;j>0 && slice_data[j] < slice_data[j - 1];j-- {
			swap(slice_data, j, j - 1)
		}
	}
}

//定义一个交换元素的swap函数
func swap(_slice []int, i int, j int) {
	_slice[i], _slice[j] = _slice[j], _slice[i]
}

func main() {
	//切片的测试数据，元素是0~9的随机顺序
	test_data := []int{8,2,4,1,7,5,3,9,0,6}
	//排序前的数据
	fmt.Println(test_data)
	InsertionSort(test_data)
	//排序后的数据
	fmt.Println(test_data)
}