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