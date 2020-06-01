package test

import "fmt"

// 声明的一种方式
var aa, bb int

// 因式分解关键字 的写法 一般用于全局变量声明
var (
	aaa int
	bbb bool
)

func test() {

	var a = 10
	var str = "chen hong"

	fmt.Println(a)
	fmt.Println(str)
}

func test_1() {

	// 主动声明 变量 b 的类型
	var b bool = true
	fmt.Println(b)

	// 没有java中的字符串的 直接拼接，如果 打印多个参数， 必须使用逗号进行分开
	// 有点类似C++的打印方式
	fmt.Println("布尔：", b)
}

func test_2() {

	// 这两种有什么区别，待区分
	// 占用32 或 64 位
	var c1 uint8 = 32
	var c2 uint = 32

	fmt.Println(c1, " : ", c2)
}

func test_3() {

	// 如果默认不初始化，则系统默认初始化为0
	var b int
	fmt.Println(b)

	//  字符串默认不初始化，则系统默认初始化为""
	var str string
	fmt.Println(str)

	//  新的声明方式
	//  :=  符号左侧一定要有新的变量（即没有被声明过的），否则就会编译出错
	//  :=  是一个声明语句
	s := "chen hong"
	fmt.Println(s)
}

func test_4() {

	//  指针，默认初始值 为nil，nil 相当于 nullpointer
	var a *int = nil

	fmt.Println(a)
}
