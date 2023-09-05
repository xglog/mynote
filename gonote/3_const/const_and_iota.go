package main

import "fmt"

//const 在go中来定义枚举类型(常量)

const (
	//关键字 iota(只能在const内使用), 每行的iota都会累加1,第一行所有默认值是0
	BEIJING  = iota //iota=0(第一行所有默认值)
	SHANGHAI        //iota=1
	SHENZHEN        //iota=2
)
const (
	//iota的公式默认跟随上一个
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, j // g,h iota公式变了,g,h下面的都跟着变
)

func Const_iota() {
	const length int = 10
	fmt.Println("lenth=", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("c,d =", c, d)
	fmt.Println("i,j =", i, j)
}
func main() {
	Const_iota()
}