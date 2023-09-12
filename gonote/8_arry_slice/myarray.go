package main

import (
	"fmt"
)

func printArray(myArray [4]int) {

	for index, v := range myArray {
		myArray[index] = v + 1
	}
	//数组属于值传递,方法内数组是复制了传入的数组
	for _, v := range myArray {
		fmt.Print(v)
	}
}

func main() {
	//固定长度的数组
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{1, 2, 3, 4}
	// for i :=o;i<len(myArray1); i++
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i], "--")
	}
	for index, value := range myArray2 {
		fmt.Print(index, "-", value, "\n")
	}
	for _, v := range myArray3 {
		fmt.Print("myArray3-", v, "\n")
	}
	fmt.Printf("不同长度的数组,被看作不同的数据类型,myArray1=%T\n", myArray1)
	fmt.Printf("不同长度的数组,被看作不同的数据类型,myArray2=%T\n", myArray2)
	fmt.Printf("不同长度的数组,被看作不同的数据类型,myArray3=%T\n", myArray3)

	printArray(myArray3)

	//方法定义的参数是4长度的数组,传入10长度的数组会报错
	//printArray([4]int(myArray1))
	Myslice()
}
