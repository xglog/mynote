package main

import "fmt"

func printslice(myslice []int) {
	for index, v := range myslice {
		myslice[index] = v + 1
	}

}

func Myslice() {
	myslice := []int{1, 2, 3} //动态数组,切片 slice
	//slice是指针传递,会修改原slice的值
	printslice(myslice)

	fmt.Printf("%T\n", myslice)
	for _, v := range myslice {
		fmt.Print(v,"\n")
	}
}
