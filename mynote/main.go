/*
包
每个 Go 程序都是由包构成的。
程序从 main 包开始运行。
*/
package main

import (
	"mynote/lessons_1"
)

func main() {
	lessons_1.Add(3, 5)
	lessons_1.MyFavoriteNumber()
}
