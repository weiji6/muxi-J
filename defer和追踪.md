关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
```go
package main
import "fmt"

func main() {
    Function1()
}

func Function1() {
    fmt.Printf("In Function1 at the top\n")
    defer Function2()
    fmt.Printf("In Function1 at the bottom!\n")
}

func Function2() {
    fmt.Printf("Function2: Deferred until the end of the calling function!")
}
```
输出：
```go
In Function1 at the top
In Function1 at the bottom!
Function2: Deferred until the end of the calling function!
```
使用 defer 的语句同样可以接受参数
```go
package main

import "fmt"

func main() {
	a()
}
func a() {
	i := 0
	defer fmt.Println(i)
    i++
	return
}
```
```go
package main

import "fmt"

func main() {
	a()
}
func a() {
	i := 0
    i++
    defer fmt.Println(i)
	return
}
```
第一种defer捕获定义时的值，第二种defer捕获i++后的值
输出结果分别为：1，0

当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：
```go
func f() {
    for i := 0; i < 5; i++ {
        defer fmt.Printf(“%d “, i)
    }
}
```
关键字 defer 允许我们进行一些函数执行完成后的收尾工作，例如：
1.关闭文件流：
// open a file defer file.Close() 
2.解锁一个加锁的资源:
mu.Lock() defer mu.Unlock() （详见第 9.3 节）
3.打印最终报告
printHeader() defer printFooter()
4.关闭数据库链接
// open a database connection defer disconnectFromDB()