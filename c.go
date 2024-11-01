package main

import (
	"fmt"
)

// F1 接受一个变长参数并把它传递给 F2 和 F3
func F1(s ...string) {
	F2(s...) // 将变长参数传递给 F2
	F3(s)    // 将变长参数作为切片传递给 F3
}

// F2 接受变长参数并打印每个元素
func F2(s ...string) {
	for _, str := range s {
		fmt.Println(str)
	}
}

// F3 接受一个字符串切片并打印每个元素
func F3(s []string) {
	for _, str := range s {
		fmt.Println(str)
	}
}

func main() {
	// 测试 F1 函数
	F1("Hello", "World", "Go", "Language")
}
