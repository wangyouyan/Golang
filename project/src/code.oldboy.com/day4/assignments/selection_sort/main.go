package main

import (
	"fmt"
)
/*选择排序的基本思想是对待排序的记录序列进行n-1遍的处理，
第i遍处理是将L[i..n]中最小者与L[i]交换位置。
这样，经过i遍处理之后，前i个记录的位置已经是正确的了
对给定的数组进行多次遍历，每次均找到最大的一个值的索引.
*/
func SelectionSort(nums []int) {
	//取数组的长度
	length := len(nums)
	//根据数组的长度遍历数组
	//当i为0时,j为1,如果数组nums[1]的值大于nums[0],此时就把j所在索引值赋给maxIndex
	for i:=0;i<length;i++ {
		//定义索引初始值为0
		maxIndex := 0
		//寻找最大的一个数,保存索引值
		for j := 1;j < length-i;j++ {
			if  nums[j] > nums[maxIndex] {
				maxIndex = j
		}
	}
	//交换位置.
	nums[length-i-1], nums[maxIndex] = nums[maxIndex], nums[length-i-1]
}
}

func main() {
	s := []int{9,0,6,5,8,2,1,7,4,3}
	fmt.Println(s)
	SelectionSort(s)
	fmt.Println(s)
}