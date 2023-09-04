package main

/*
	四种变量声明方式,:=只能用在方法内
*/

import (
	"fmt"
)

func My_var() {
	//方法一:声明一个变量,无初始化值,默认值是0
	var a int
	fmt.Println("a = ", a)

	// 方法二:带初始化值
	var a_a int = 100
	fmt.Println("a_a=", a_a)

	// 方法三:初始化的时候可以省略数据类型,自动推断
	var b = 4
	fmt.Printf("b = %v,type of b is :%T\n", b, b)

	// 方法四:用:=省略var关键字,只能在函数内使用,通过自动匹配变量的数据类型
	c, d := 3, 4
	fmt.Println("c=", c, "d=", d)
	fmt.Printf("type is c = %T,d = %T", c, d)
}

func main() {
	My_var()
}
