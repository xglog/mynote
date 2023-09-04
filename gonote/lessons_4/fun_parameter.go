package main

import (
	"fmt"
)

// 返回一个匿名参数
func test1(a int, b int) int {
	return a + b
}

// 返回多个匿名参数
func test2(a int, b int) (int, int) {
	return a + 1, b + 1
}

// 返回多个有形参数,名称对上直接return
func test3(a int, b int) (a_ int, b_ int) {

	a_ = a + 1
	b_ = b + 1
	return
}

// 返回同类型参数简写
func test4(a int, b int) (a_, b_ int) {
	a_ = a + 1
	b_ = b + 1
	return
}

func main() {
	c := test1(1, 3)
	fmt.Println(c)

	c, d := test2(1, 2)
	fmt.Println(c, d)

	e, f := test3(1, 2)
	fmt.Println(e, f)

	a_, b_ := test4(1, 2)
	fmt.Println(a_, b_)

}
