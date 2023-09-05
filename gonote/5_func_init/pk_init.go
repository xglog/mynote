/*
每个包都有一个init函数,类似java的构造函数,
init函数执行的顺序
main
↓
import pkg1 -->pkg1 const... -->pkg1 var...-->pkg1 init() ↓

↓
const...
↓
var...
↓
init()...
↓
main()

*/

package main

//有导入的包,则先初始化导入包的整体,const-->var-->init()
import (

	// 别名
	mylib1 "gonote/lessons_5/lib1"

	// 包名前面加 . 可以直接使用包内方法(**不建议使用**)
	// . "gonote/lessons_5/lib1"
	// go编译不允许导包后不使用
	// 当想用某个包的init()函数时,可以使用匿名导入,在包前面加入 _ .
	// 这样就可以调用这个包的 init()函数了
	_ "gonote/lessons_5/lib2"
)

func main() {

	mylib1.Test_lib1()

}
