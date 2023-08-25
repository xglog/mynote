/* vscode git 配置问题

# Linux终端
sudo apt-get install update
sudo apt-get install git
git --version

git config --global user.name "your name"
git config --global user.email "you email"

# 找到项目文件夹进入，进行初始化git
git init

# 添加远程仓库（这里使用ssh）
git remote add origin git@github.com:xglog/mynote.git
# 查看远程仓库名称
git remote
# origin

# 创建一个公钥,取名字位git.hub,默认生成在~/.ssh目录下面
ssh-keygen -C "your email" -t rsa

# 添加公钥到github上

# 测试是否添加成功
ssh -T git@github.com
# Hi xglog! You've successfully authenticated, but GitHub does not provide shell access.


#将本地仓库推送到远程仓库
git push -u origin main
*/

/*
包
每个 Go 程序都是由包构成的。
程序从 main 包开始运行。
本程序通过导入路径 "fmt" 和 "math/rand" 来使用这两个包。
按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。
*注意：此程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字
*/

package lessons_1

import ( //此代码用圆括号组合了导入，这是“分组”形式的导入语句。
	"fmt"
	"math/rand"
)

// func add(x int, y int) int { 	//*注意类型在变量名 之后。
func Add(x, y int) int { //当连续两个或多个函数的已命名形参类型相同时，可以简写

	sum := x + y
	fmt.Println("the sum is :", sum)
	return sum
}

func MyFavoriteNumber() {

	fmt.Println("My favorite number is", rand.Intn(10))
}
