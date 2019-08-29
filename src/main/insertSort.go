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
func main() {
	arr := []int{7, 6, 10, 8, 3, 1}
	//insertSort(arr)
	//shellSort(arr)
	//quiklySort(arr,0,len(arr)-1)

}

/**
  quik
*/
func quiklySort(arr []int, st, ed int) []int {
	//选取基准
	if st < ed {
		//以第一个数为基准，第一遍过滤
		pos := part(arr, st, ed)
		//递归 左部分 和 右部分 直到st>= ed,因为POS是从1开始的，所以要-1 才是左边部分的实际下标
		quiklySort(arr, st, pos-1)
		quiklySort(arr, pos+1, ed)
	}
	fmt.Print(arr)
	return nil
}

func part(arr []int, st, ed int) int {
	//设置第一个为基准
	pv := st
	//大数指针
	bigPtr := pv + 1
	for i := bigPtr; i <= ed; i++ {
		if arr[i] < arr[pv] {
			//I 是前面的比基准小的数，
			sort.Ch(arr, i, bigPtr)
			// 发生交换后，大数指针往前走，指向后一个大数,如果没发生交换，大数就是当前i 不变
			bigPtr += 1
		}
	}
	//最后 把基准指针放在大数指针前一的位置， 因为这时候最大的数就是大数指针，所以要换成大数指针前一位，前面就是比基准小的部分
	sort.Ch(arr, pv, bigPtr-1)
	//返回 大数指针的位置前的部分
	return bigPtr - 1
}

/**
希尔
*/
func shellSort(arr []int) {
	cnt := 0

	//定义高度H 的数组
	h := len(arr) / 3
	if h < len(arr) {
		h = len(arr)*3 + 1
	}
	//h=1 高度为1是最后一次全排
	for h >= 1 {
		for i := h; i < len(arr); i++ {
			//从第二个分组开始，逐个跟之前H个位置比较
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				sort.Ch(arr, j-h, j)
				cnt++
				fmt.Print(arr)
			}

		}
		h /= 3
	}
	fmt.Print(arr, "次数为", cnt)

}

/**
插入排序
*/
func insertSort(arr []int) {
	//arr:=[]int{343,233,3444,121,33,432,433342342,4334,5677,67676,878,003}
	cnt := 0
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			sort.Ch(arr, j-1, j)
			cnt++
		}

	}
	fmt.Println(arr, "比较次数", cnt)

}
