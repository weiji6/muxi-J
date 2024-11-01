```go
package main

var a = "G"

func main() {
   n()
   m()
   n()
}

func n() { print(a) }

func m() {
   a := "O"
   print(a)
}
```

输出结果GOG
全局变量为a赋值为G
n()函数定义执行输出a的值
m()函数定义执行重新定义一个与全局变量a不同的变量a赋值为O，并且输出赋值后的a
最终在main函数中按顺序执行n(),m(),n()