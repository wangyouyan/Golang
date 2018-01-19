package main

import (
	"fmt"
)
/*快速排序基本思想是：通过一趟排序将要排序的数据分割成独立的两部分，
其中一部分的所有数据都比另外一部分的所有数据都要小，
然后再按此方法对这两部分数据分别进行快速排序，
整个排序过程可以递归进行，以此达到整个数据变成有序序列。
1）设置两个变量i、j，排序开始的时候：i=0，j=N-1；
2）以第一个数组元素作为关键数据，赋值给key，即 key=A[0]；
3）从j开始向前搜索，即由后开始向前搜索（j -- ），找到第一个小于key的值A[j]，A[i]与A[j]交换；
4）从i开始向后搜索，即由前开始向后搜索（i ++ ），找到第一个大于key的A[i]，A[i]与A[j]交换；
5）重复第3、4、5步，直到 I=J； (3,4步是在程序中没找到时候j=j-1，i=i+1，直至找到为止。
	找到并交换的时候i， j指针位置不变。另外当i=j这过程一定正好是i+或j-完成的最后令循环结束。）
*/

//定义函数，传参为int类型的数组,两个int类型的参数分别为left,right
func QuickSort(arr []int, left, right int) {
	//
	if left < right {
		i, j := left, right
		//取数组长度的1/2赋值给key
		key := arr[(left+right)/2]
		for i <= j {
			//如果数组左边的值比key小，则一直递加
			for arr[i] < key {
				i++
			}
			//如果数组右边的值比key大，则一直递减
			for arr[j] > key {
				j--
			}
			//当左边的元素所在的索引长度和右边的索引长度相等时
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		//如果左边的元素小于右边的元素
		if left < j {
			QuickSort(arr, left, j)
		}
		//如果右边的元素小于左边的元素
		if right > i {
			QuickSort(arr, i, right)
		}
	}
}

func main() {
	test_data := []int{8,2,4,1,7,5,3,9,0,6}
	fmt.Println("Before sorting:\n",test_data)
	QuickSort(test_data, 0, len(test_data)-1)
	fmt.Println("After sorting:\n",test_data)
}