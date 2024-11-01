```go
package main

var a = "G"

func main() {
   n()
   m()
   n()
}

func n() {
   print(a)
}

func m() {
   a = "O"
   print(a)
}
```go

输出结果为GOO
全局变量为a赋值为G
n()函数定义执行输出a的值
m()函数定义执行将全局变量a赋值为O，即改变了全局变量
最终在main函数中按顺序执行n()输出初始全局变量a，m()重新赋值全局变量为O，并输出新的全局变量a，n()输出新的全局变量a