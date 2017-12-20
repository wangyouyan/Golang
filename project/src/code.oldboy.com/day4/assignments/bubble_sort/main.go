package main

import (
	"fmt"
)
/*附上python实现冒泡算法的代码
li=[13,22,33,99,11]
print("Before sorting:%s",li)
for n in range(1,len(li)):
    for m in range(len(li)-n):
        #获取m个值
        num1 = li[m]
        #获取m+1个值
        num2 = li[m+1]
        if num1 > num2:
        #将较大的值放在右侧
            temp = li[m]
            li[m] = num2
            li[m+1] = temp
print("After sorting:%s",li)
*/

func main() {
	//构造初始切片数据
	slice_data := [8]int{18,20,11,17,8,9,5,22}
	fmt.Println("The slice data before sorting is:\n", slice_data)
	len_num := len(slice_data)
	for i := 0;i < len_num;i++ {
		//从大到小，道理是一样的
		for j := i + 1;j < len_num;j++ {
			if slice_data[i] > slice_data[j] {
				slice_data[i], slice_data[j] = slice_data[j],slice_data[i]
			}
		}
	}
	//打印按照从小到大的切片元素
	fmt.Println("The slice data after sorting is:")
	fmt.Println(slice_data)
}