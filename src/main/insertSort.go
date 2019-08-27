package main

import (
	"fmt"
	"main/sort"
)

/**
和选择排序不同的是，插入排序所需的时间取决于输入中元素的初始顺序。例如，对一个很大
且其中的元素已经有序（或接近有序）的数组进行排序将会比对随机顺序的数组或是逆序数组进行
排序要快得多。
 */
func main()  {
	arr:=[]int{343,233,3444,121,33,432,433342342,4334,5677,67676,878,3}
	//insertSort(arr)
	//shellSort(arr)

}

func quiklySort()  {

}

/**
希尔
 */
func shellSort(arr []int )  {
	cnt:=0

	//定义高度H 的数组
	h:= len(arr) / 3
	if(h<len(arr)){
		h=len(arr)*3+1
	}
	//h=1 高度为1是最后一次全排
	for h>=1 {
		for i:=h;i<len(arr);i++ {
			for j:=i;j>=h && arr[j]<arr[j-h];j-=h{
				sort.Ch(arr,j-h,j)
				cnt++
				fmt.Print(arr)
			}

		}
		h/=3
	}
	fmt.Print(arr,"次数为",cnt)




}
/**
插入排序
 */
func insertSort(arr []int) {
	//arr:=[]int{343,233,3444,121,33,432,433342342,4334,5677,67676,878,003}
	cnt:=0;
		for i:=1;i<len(arr);i++ {
			for j:=i;j>0 && arr[j]<arr[j-1];j--{
					 sort.Ch(arr,j-1,j)
					cnt++
			}

		}
		fmt.Println(arr,"比较次数",cnt)

}
