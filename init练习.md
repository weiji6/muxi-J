init练习1
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

init练习2
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
```

输出结果为GOO
全局变量为a赋值为G
n()函数定义执行输出a的值
m()函数定义执行将全局变量a赋值为O，即改变了全局变量
最终在main函数中按顺序执行n()输出初始全局变量a，m()重新赋值全局变量为O，并输出新的全局变量a，n()输出新的全局变量a

init练习3
```go
package main

var a string

func main() {
   a = "G"
   print(a)
   f1()
}

func f1() {
   a := "O"
   print(a)
   f2()
}

func f2() {
   print(a)
}
```

输出结果GOG
全局变量为a,并在main函数中赋值为G
f1()函数定义执行定义局部变量a为O，输出局部变量a，再调用f2()函数
f2()函数定义执行输出全局变量a，即G，因为f1中的a是局部变量，在f2()中不可见
最终在main函数中按顺序执行