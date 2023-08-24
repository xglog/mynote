/*
包
每个 Go 程序都是由包构成的。
程序从 main 包开始运行。
本程序通过导入路径 "fmt" 和 "math/rand" 来使用这两个包。
按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。
*注意：此程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字
*/

package main

import ( //此代码用圆括号组合了导入，这是“分组”形式的导入语句。
	"fmt"
	"math/rand"
)

// func add(x int, y int) int { 	//*注意类型在变量名 之后。
func add(x, y int) int { //当连续两个或多个函数的已命名形参类型相同时，可以简写
	return x + y
}

func main1() {
	fmt.Println("My favorite number is", rand.Intn(10)) //在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。(类似java的private public)
	fmt.Println("执行函数add", add(3, 5))
}
