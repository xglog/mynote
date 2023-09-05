package main

import (
	"fmt"
)

/*
多个defer的执行顺序是栈的结构,先入先出
defer实在return之后执行
*/

func return_after() int {
	fmt.Println("我是1里面的return 执行了")
	return 0
}


func defer_call() int {
	fmt.Println("1")
	defer fmt.Println("我是1里面的defer,看看我和1里面的return谁先执行")
	return return_after()
}
func defer_call_2()  {
	fmt.Println("2")
}
func defer_call_3()  {
	fmt.Println("3")
}

// 先执行 defer defer_call方法,所以他被压在栈底,采用后进先出,所以执行结果321
/*
|				 | <-- 栈	 call defer_call3()
|                |
| defer_call2()	 |
| defer_call()	 |	
|________________|_	
*/
func main() {
	fmt.Println("我是一个方法")
	defer defer_call()
	defer defer_call_2()
	defer defer_call_3()
}

